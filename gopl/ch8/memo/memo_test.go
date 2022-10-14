package memo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.google.cn",
			"https://pkg.go.dev/",
			"http://gopl.io",
			"https://golang.google.cn",
			"https://pkg.go.dev/",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()

	return ch
}

func TestSequential(t *testing.T) {
	m := New(httpGetBody)

	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

// Run this with go test -run=TestConcurrent -race -v
func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				t.Error(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))

			n.Done()
		}(url)
	}

	n.Wait()
}
