package store

import (
	"log"
	"net/rpc"
)

type ProxyStore struct {
	urls   *URLStore
	client *rpc.Client
}

func NewProxyStore(addr string) *ProxyStore {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal("Error constructing ProxyStore:", err)
	}
	return &ProxyStore{urls: NewURLStore(""), client: client}
}

func (p *ProxyStore) Get(key, url *string) error {
	// url found in local map
	if p.urls.Get(key, url) == nil {
		return nil
	}

	// url not found in local map, make rpc-call
	if err := p.client.Call("Store.Get", key, url); err != nil {
		return err
	}

	// add key,url to local map
	p.urls.Put(key, url)
	return nil
}

func (p *ProxyStore) Put(url, key *string) error {
	err := p.client.Call("Store.Put", url, key)
	if err != nil {
		return err
	}
	p.urls.set(url, key)
	return nil
}
