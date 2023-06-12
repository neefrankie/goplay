package heap

import (
	"sort"
	"testing"
)

func TestTrickleUp(t *testing.T) {
	type args struct {
		data sort.Interface
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Trickle Up",
			args: args{
				data: sort.IntSlice{2, 3, 4},
				k:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TrickleUp(tt.args.data, tt.args.k)
		})

		t.Logf("%v", tt.args.data)
	}
}

func TestTrickleDown(t *testing.T) {
	type args struct {
		data sort.Interface
		k    int
		n    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Trickle down",
			args: args{
				data: sort.IntSlice{2, 3, 4},
				k:    0,
				n:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TrickleDown(tt.args.data, tt.args.k, tt.args.n)
		})

		t.Logf("%v", tt.args.data)
	}
}
