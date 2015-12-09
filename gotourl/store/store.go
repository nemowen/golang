package store

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

const SaveQueueLength = 1000

type URLStore struct {
	urls map[string]string
	lock sync.RWMutex
	ch   chan record
}

type record struct {
	Key, Url string
}

func NewURLStore(filename string) *URLStore {
	s := &URLStore{urls: make(map[string]string, 100)}
	if filename != "" {
		s.ch = make(chan record, SaveQueueLength)
		if err := s.load(filename); err != nil {
			log.Fatal("Load Error:", err)
		}

		go s.saveLoop(filename)
	}
	return s
}

func (u *URLStore) load(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	de := json.NewDecoder(file)
	for err != io.EOF {
		var r record
		if err = de.Decode(&r); err == nil {
			u.set(r.Key, r.Url)
		}
	}
	return nil
}

func (u *URLStore) Get(key, url *string) error {
	u.lock.RLock()
	defer u.lock.RUnlock()
	if u, ok := u.urls[*key]; ok {
		*url = u
		return nil
	}
	return errors.New("key not found")
}

func (u *URLStore) set(key, url string) error {
	u.lock.Lock()
	defer u.lock.Unlock()
	if _, present := u.urls[key]; present {
		return errors.New("key already exists")
	}
	u.urls[key] = url
	return nil
}

func (u *URLStore) Count() int {
	u.lock.RLock()
	defer u.lock.RUnlock()
	return len(u.urls)
}

func (u *URLStore) Put(url, key *string) error {
	for {
		*key = genKey(u.Count())
		if err := u.set(*key, *url); err != nil {
			return err
		} else {
			if u.ch != nil {
				u.ch <- record{*key, *url}
			}
			return nil
		}
	}

}

func (u *URLStore) saveLoop(filename string) {
	var file *os.File
	var err error
	file, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("URLStore:", err)
	}
	defer file.Close()

	encode := json.NewEncoder(file)
	for {
		r := <-u.ch
		if err = encode.Encode(r); err != nil {
			log.Println("URLStore:", err)
		}
	}
}

func genKey(v int) string {
	return strconv.Itoa(v) + "a"
}
