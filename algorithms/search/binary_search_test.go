package search

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		x   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Binary search",
			args: args{
				x:   19,
				arr: []int{1, 2, 3, 5, 6, 7, 8, 10, 12, 13, 15, 16, 18, 19, 20, 22},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.x, tt.args.arr); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
