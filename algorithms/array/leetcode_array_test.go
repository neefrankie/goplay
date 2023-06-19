package array

import (
	"reflect"
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1,7,3,6,5,6",
			args: args{
				nums: []int{1, 7, 3, 6, 5, 6},
			},
			want: 3,
		},
		{
			name: "1,2,3",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: -1,
		},
		{
			name: "2,1,-1",
			args: args{
				nums: []int{2, 1, -1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dominantIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3,6,1,0",
			args: args{
				nums: []int{3, 6, 1, 0},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dominantIndex(tt.args.nums); got != tt.want {
				t.Errorf("dominantIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "9,9,9",
			args: args{
				digits: []int{9, 9, 9},
			},
			want: []int{1, 0, 0, 0},
		},
		{
			name: "9,9,8",
			args: args{
				digits: []int{9, 9, 8},
			},
			want: []int{9, 9, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "11 + 1",
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRune(t *testing.T) {
	for _, v := range "1001" {
		t.Logf("%v", v)
	}

	t.Logf("%v", '0')
	t.Logf("%v", int('1'-'0'))

	t.Logf("%b", byte(1))
}

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sad in sadbutsad",
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			name: "mississippi issip",
			args: args{
				haystack: "mississippi",
				needle:   "issip",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "flower,flow,flight",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "dog,racecar,car",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "hello",
			args: args{
				s: []byte("hello"),
			},
			want: []byte("olleh"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arrayPairSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1,4,3,2",
			args: args{
				nums: []int{1, 4, 3, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayPairSum(tt.args.nums); got != tt.want {
				t.Errorf("arrayPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "2,7,11,15",
			args: args{
				numbers: []int{2, 7, 11, 15},
				target:  9,
			},
			want: []int{1, 2},
		},
		{
			name: "2,3,4",
			args: args{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum2(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "2,7,11,15",
			args: args{
				numbers: []int{2, 7, 11, 15},
				target:  9,
			},
			want: []int{1, 2},
		},
		{
			name: "2,3,4",
			args: args{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum2(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1,2,3,4,5,6,7",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)

			t.Logf("%v", tt.args.nums)
		})
	}
}
