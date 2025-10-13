package netcat

import (
	"io"
	"log"
	"net"
	"os"
)

func Netcat(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// Reads the standard input and sends it to the server.
	MustCopy(os.Stdout, conn)
}

func Netcat2(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	// 把服务器的输出复制到标准输出
	go MustCopy(os.Stdout, conn)
	// 把标准输入复制到服务端
	MustCopy(conn, os.Stdin)
}

func Netcat3(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		// 读方向的连接将导致 io.Copy 函数返回错误，
		// 因此没有处理错误
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	// 当用户关闭了标准输入，MustCopy 函数调用将返回，
	// 然后调用 conn.Close() 关闭读和写方向的网络连接。
	// press Ctrl-D on Mac/Linux
	// Ctrl-Z + Return on Windows
	MustCopy(conn, os.Stdin)
	// 关闭写方向的连接将导致 server 程序收到一个文件（end-of-file）结束的信号。
	// 关闭读方向的连接将导致前面的后台goroutine的io.Copy 函数返回
	// 一个 read from closed connection 类似的错误
	conn.Close()
	// 退出前先等待从 done 对应的channel接收一个值。
	<-done
}
