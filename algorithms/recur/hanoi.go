package recur

import "fmt"

func Hanoi(disks int, from, interim, to string) {
	if disks == 1 {
		fmt.Printf("Disk 1 from %s to %s\n", from, to)
	} else {
		Hanoi(disks-1, from, to, interim)

		fmt.Printf("Disk %d from %s to %s\n", disks, from, to)
		Hanoi(disks-1, interim, from, to)
	}
}
