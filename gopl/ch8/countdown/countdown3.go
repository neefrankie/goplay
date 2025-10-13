package countdown

import (
	"fmt"
	"time"
)

func Countdown3() {
	fmt.Println("Commencing countdown. Press return to abort.")

	tick := time.NewTicker(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick.C:
		case <-time.After(5 * time.Second):
			fmt.Println("Countdown aborted!")
			return
		}
	}

	tick.Stop()
}
