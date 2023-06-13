package array

import (
	"reflect"
	"testing"
)

func Test_replaceElements(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[17,18,5,4,6,1]",
			args: args{
				arr: []int{17, 18, 5, 4, 6, 1},
			},
			want: []int{18, 6, 6, 6, 1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceElements(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,1,2,3,4,4,4]",
			args: args{
				nums: []int{1, 1, 2, 3, 4, 4, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[0,1,0,3,12]",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "[0,0,0,9,3,0,0,12,0,0]",
			args: args{
				nums: []int{0, 0, 0, 9, 3, 0, 0, 12, 0, 0},
			},
			want: []int{9, 3, 12, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)

			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("replaceElements() = %v, want %v", tt.args.nums, tt.want)
			}

			t.Logf("Got %v", tt.args.nums)
		})
	}
}

func Test_sortArrayByParity(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[3,1,2,4]",
			args: args{
				nums: []int{3, 1, 2, 4},
			},
			want: []int{2, 4, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParity(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[3,2,2,3]",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}

			t.Logf("%v", tt.args.nums)
		})
	}
}

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "-4,-1,0,3,10",
			args: args{
				nums: []int{-4, -1, 0, 3, 10},
			},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "-7,-3,2,3,11",
			args: args{
				nums: []int{-7, -3, 2, 3, 11},
			},
			want: []int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
