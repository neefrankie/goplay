package main

import (
	"fmt"
	"log"

	"example.com/fetch"
	"github.com/fsnotify/fsnotify"
	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.ToUpper("Hello"))

	fetch.New()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(".tmp")
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}