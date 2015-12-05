package store

import (
	"encoding/gob"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

type URLStore struct {
	urls map[string]string
	lock sync.RWMutex
	file *os.File
}

type record struct {
	Key, Url string
}

func NewURLStore(filename string) *URLStore {
	s := &URLStore{urls: make(map[string]string, 100)}
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	s.file = f
	if s.load() != nil {
		log.Fatal("Load Error:", err)
	}
	return s
}

func (u *URLStore) load() error {
	if _, err := u.file.Seek(0, 0); err != nil {
		return err
	}
	de := gob.NewDecoder(u.file)
	var err error
	for err != io.EOF {
		var r record
		if err = de.Decode(&r); err == nil {
			u.set(r.Key, r.Url)
		}
	}
	return nil
}

func (u *URLStore) Get(key string) string {
	u.lock.RLock()
	defer u.lock.RUnlock()
	return u.urls[key]
}

func (u *URLStore) set(key, url string) bool {
	u.lock.Lock()
	defer u.lock.Unlock()
	if _, present := u.urls[key]; present {
		return false
	}
	u.urls[key] = url
	return true
}

func (u *URLStore) Count() int {
	u.lock.RLock()
	defer u.lock.RUnlock()
	return len(u.urls)
}

func (u *URLStore) Put(url string) string {
	for {
		key := genKey(u.Count())
		if u.set(key, url) {
			if err := u.save(key, url); err != nil {
				log.Fatal("save url error:", err)
			}
			return key
		}
		return ""
	}
}

func (u *URLStore) save(key, url string) error {
	e := gob.NewEncoder(u.file)
	return e.Encode(record{key, url})
}

func genKey(v int) string {
	return strconv.Itoa(v) + "a"
}
