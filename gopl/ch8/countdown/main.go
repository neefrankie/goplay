package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.")

	ticker := time.NewTicker(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)

		select {
		case <-ticker.C:

		case <-abort:
			fmt.Println("Launch aborted")
			ticker.Stop()
			return
		}
	}

	ticker.Stop()
	fmt.Println("Launch")
}
