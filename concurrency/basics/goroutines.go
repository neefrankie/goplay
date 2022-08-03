package basics

import (
	"fmt"
	"sync"
)

func SayHello() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}

	wg.Add(1)
	go sayHello()

	wg.Wait()
}

func Salutation() {
	var wg sync.WaitGroup

	salutation := "hello"

	wg.Add(1)

	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()

	wg.Wait()

	fmt.Println(salutation)
}

func Salutations() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}

	wg.Wait()
}
