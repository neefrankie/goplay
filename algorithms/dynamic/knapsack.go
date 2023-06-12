package dynamic

import "log"

type sackItem struct {
	size  int
	value int
}

var knapsackItems = []sackItem{
	{
		size:  3,
		value: 4,
	},
	{
		size:  4,
		value: 5,
	},
	{
		size:  7,
		value: 10,
	},
	{
		size:  8,
		value: 11,
	},
	{
		size:  9,
		value: 13,
	},
}

func NaiveKnapsack(capacity int, items []sackItem) int {
	var maxVal = 0

	log.Printf("Capcacity %d\n", capacity)

	for _, item := range items {
		space := capacity - item.size
		if space >= 0 {

			log.Printf("Divide %d = %d + %d\n", capacity, space, item.size)
			valSum := NaiveKnapsack(space, items) + item.value
			log.Printf("Capacity %d could carry value %d\n", capacity, valSum)

			if valSum > maxVal {
				maxVal = valSum
				log.Printf("Capacity %d could carry max value %d\n", capacity, maxVal)
			}
		}
	}

	return maxVal
}
