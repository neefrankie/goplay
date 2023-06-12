package sorting

import "testing"

func Test_partition(t *testing.T) {
	ints := []int64{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11}
	type args struct {
		arr []int64
		p   int
		r   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Partition int",
			args: args{
				arr: ints,
				p:   0,
				r:   len(ints) - 1,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.args.arr, tt.args.p, tt.args.r)
			if got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
				return
			}

			t.Logf("%v", ints)
		})
	}
}

func TestQuickSort(t *testing.T) {
	ints := []int64{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11}

	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Sort int",
			args: args{
				arr: ints,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr)

			t.Logf("%v", ints)
		})
	}
}
