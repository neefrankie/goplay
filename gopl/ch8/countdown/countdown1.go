package countdown

import (
	"fmt"
	"os"
	"time"
)

func Countdown1() {
	fmt.Println("Commencing countdown. Press return to abort.")
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	select {
	// 返回抑遏channel，创建一个新的goroutine，在经过特定事件后想改chnnel发送一个独立的值。
	case <-time.After(10 * time.Second):

	case <-abort:
		fmt.Println("Lanuch Aborted!")
	}

	launch()
}

func launch() {
	fmt.Println("Lanuching!")
}
