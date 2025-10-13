package reverb

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func HandleReverb1(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}

	c.Close()
}

func HandleReverb2(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		// go 后跟的函数的参数会在 go 语句自身执行时求值
		// 因此 input.Text() 会在 main goroutine 中被求值
		go echo(c, input.Text(), 1*time.Second)
	}

	c.Close()
}
