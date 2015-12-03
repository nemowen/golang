package gotourl

import (
	"strconv"
	"sync"
)

type URLStore struct {
	urls map[string]string
	lock sync.RWMutex
}

func NewURLStore() *URLStore {
	return &URLStore{urls: make(map[string]string, 100)}
}

func (u *URLStore) get(key string) string {
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
			return key
		}
		return ""
	}
}

func genKey(v int) string {
	return strconv.Itoa(v) + "a"
}
