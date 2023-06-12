package elementary

import "fmt"

func quickFind(pairs [][]int) {
	var id [1000]int
	for i := 0; i < 1000; i++ {
		id[i] = i
	}

	for _, pair := range pairs {
		p := pair[0]
		q := pair[1]

		if id[p] == id[q] {
			continue
		}

		t := id[p]
		for i := 0; i < 1000; i++ {
			if id[i] == t {
				id[i] = id[q]
			}
		}

		fmt.Printf(" %d %d\n", p, q)
	}

	fmt.Printf("%v", id)
}
