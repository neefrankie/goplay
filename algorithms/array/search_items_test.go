package array

import (
	"reflect"
	"testing"
)

func Test_checkIfExist(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[-20,8,-6,-14,0,-19,14,4]",
			args: args{
				arr: []int{-20, 8, -6, -14, 0, -19, 14, 4},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfExist(tt.args.arr); got != tt.want {
				t.Errorf("checkIfExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slope(t *testing.T) {
	var dir slope
	t.Logf("%b", slopeUp)
	t.Logf("%b", slopeDown)
	dir = dir | slopeUp
	t.Logf("Up %b", dir)
	dir = dir | slopeDown
	t.Logf("Down %b", dir)
	t.Logf("Contains up %t", dir&slopeUp > 0)
	t.Logf("Contains down %t", dir&slopeDown > 0)
}

func Test_validMountainArray1(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[0, 2, 3, 4, 5, 2, 1, 0]",
			args: args{
				arr: []int{0, 2, 3, 4, 5, 2, 1, 0},
			},
			want: true,
		},
		{
			name: "[0, 2, 3, 3, 5, 2, 1, 0]",
			args: args{
				arr: []int{0, 2, 3, 3, 5, 2, 1, 0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray1(tt.args.arr); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validMountainArray2(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[0, 2, 3, 4, 5, 2, 1, 0]",
			args: args{
				arr: []int{0, 2, 3, 4, 5, 2, 1, 0},
			},
			want: true,
		},
		{
			name: "[0, 2, 3, 3, 5, 2, 1, 0]",
			args: args{
				arr: []int{0, 2, 3, 3, 5, 2, 1, 0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray2(tt.args.arr); got != tt.want {
				t.Errorf("validMountainArray2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_thirdMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3,2,1",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: 1,
		},
		{
			name: "1,2",
			args: args{
				nums: []int{1, 2},
			},
			want: 2,
		},
		{
			name: "2,2,3,1",
			args: args{
				nums: []int{2, 2, 3, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}

			t.Logf("%v\n", tt.args.nums)
		})
	}
}

func Test_findDisappearedNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "4,3,2,7,8,2,3,1",
			args: args{
				nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			},
			want: []int{5, 6},
		},
		{
			name: "1,1",
			args: args{
				nums: []int{1, 1},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDisappearedNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
