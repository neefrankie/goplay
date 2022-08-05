package patterns

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_generator(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		t.Log(v)
	}
}

func TestRepeat(t *testing.T) {
	done := make(chan interface{})

	repeatStream := Repeat(done, 1, 2, 3, 4, 5, 6)

	go func() {
		defer close(done)
		time.Sleep(1 * time.Minute)
	}()

	for v := range repeatStream {
		t.Log(v)
	}
}

func TestTake(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	for num := range Take(done, Repeat(done, 1), 10) {
		fmt.Printf("%v", num)
	}
}

func TestRepeatFn(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	rand := func() interface{} {
		return rand.Int()
	}

	for num := range Take(done, RepeatFn(done, rand), 10) {
		fmt.Println(num)
	}
}

func Test_toString(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	var message string
	for token := range toString(done, Take(done, Repeat(done, "I", "am."), 5)) {
		message += token
	}

	fmt.Printf("message: %s...", message)
}
