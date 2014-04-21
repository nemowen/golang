package main

import (
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"strings"
	"time"
)

func NewWatcher2(filepath string, reply chan string) {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case evn := <-watcher.Event:
				log.Println(evn.Name)
				if evn.IsModify() {
					f, e := os.Open(filepath)
					if e != nil {
						log.Println("打开文件失败：", filepath)
						continue
					}
					defer f.Close()
					buf := make([]byte, 2)
					f.Read(buf)
					reply <- string(buf)
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch(filepath)
	if err != nil {
		log.Fatal(err)
	}

	<-done

	/* ... do stuff ... */
	watcher.Close()
}
