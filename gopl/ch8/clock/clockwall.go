package clock

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

// ClockWall connects to a set of servers and displays the time
// TZ=US/Eastern ./build/clock2 -port 8010
// TZ=Asia/Tokyo ./build/clock2 -port 8020
// Usage: ./build/clockwall NewYork=localhost:8010 Tokyo=localhost:8020
func ClockWall(addresses []string) {
	var wg sync.WaitGroup
	var messages = make(chan string)

	for _, arg := range addresses {
		parts := strings.Split(arg, "=")
		if len(parts) != 2 {
			log.Printf("invalid argument: %s", arg)
			continue
		}

		city := parts[0]
		addr := parts[1]

		wg.Add(1)
		go func(city string, addr string) {
			defer wg.Done()
			connect(messages, city, addr)
		}(city, addr)
	}

	go func() {
		wg.Wait()
		close(messages)
	}()

	for m := range messages {
		fmt.Print(m)
	}
}

func connect(ch chan<- string, city string, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		ch <- fmt.Sprintf("%s: cannot connect: %v\n", city, err)
		log.Fatal(err)
	}

	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ch <- fmt.Sprintf("%-10s: %s\n", city, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		ch <- fmt.Sprintf("%s: %v\n", city, err)
	} else {
		ch <- fmt.Sprintf("%s: connection closed\n", city)

	}
}
