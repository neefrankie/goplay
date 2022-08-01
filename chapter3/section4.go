package chapter3

import (
	"fmt"
	"time"
)

func SelectExample() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")

	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}

func SelectDistribution() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1count, c2count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1count++

		case <-c2:
			c2count++
		}
	}

	fmt.Printf("c1count: %d\nc2count: %d\n", c1count, c2count)
}

func SelectTimeout() {
	var c <-chan int
	select {
	case <-c:
		// Never unblocked when reading from a nil channel

	case <-time.After(1 * time.Second):
		fmt.Println("Time out")
	}
}

func SelectSignal() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop

		default:

		}

		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
