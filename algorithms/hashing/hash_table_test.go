package hashing

import (
	"fmt"
	"testing"
)

func TestNewHashTable(t *testing.T) {
	table := NewHashTable(12)

	table.Insert(IntHashable(108))
	table.Insert(IntHashable(13))
	table.Insert(IntHashable(0))
	table.Insert(IntHashable(113))
	table.Insert(IntHashable(5))
	table.Insert(IntHashable(66))
	table.Insert(IntHashable(117))
	table.Insert(IntHashable(47))

	t.Logf("%v", table.arr)

	item, err := table.Find(66)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", item)

	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
