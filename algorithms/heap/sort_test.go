package heap

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		data sort.Interface
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Heap Sort",
			args: args{
				data: sort.IntSlice{
					81,
					6,
					23,
					38,
					95,
					71,
					72,
					39,
					34,
					52,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.data)
		})

		t.Logf("%v", tt.args.data)
	}
}
