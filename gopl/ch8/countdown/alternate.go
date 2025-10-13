package countdown

import "fmt"

func Alternate() {
	ch := make(chan int, 1)

	for i := range 10 {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:

		}
	}
}
