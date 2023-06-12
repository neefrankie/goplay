package dma

import (
	"reflect"
	"testing"
)

func TestSumList(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sum int",
			args: args{
				arr: []int{8, 4, 2, 10, 19},
			},
			want: 43,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumList(tt.args.arr); got != tt.want {
				t.Errorf("SumList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLargestDifference(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Largest difference",
			args: args{
				arr: []int{1, 4, 9, 3, 10},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestDifference(tt.args.arr); got != tt.want {
				t.Errorf("LargestDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindDuplicates(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Find duplicates",
			args: args{
				arr: []int{2, 2, 3, 5, 7, 7, 7, 9, 10, 11, 12, 12, 13},
			},
			want: []int{2, 7, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicates(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountNegativeInts(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Count Negative",
			args: args{
				ints: []int{10, -8, 4, 2, -13, -15, -6},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountNegativeInts(tt.args.ints); got != tt.want {
				t.Errorf("CountNegativeInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastEvenInt(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Last event position",
			args: args{
				arr: []int{2, 3, 4, 5, 6, 7},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastEvenInt(tt.args.arr); got != tt.want {
				t.Errorf("LastEvenInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Is palindrome",
			args: args{
				s: "A man, a plan, a canal, Panama!",
			},
			want: true,
		},
		{
			name: "Not palindrome",
			args: args{
				s: "Hello, world",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXPower(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Positive power",
			args: args{
				x: 2,
				n: 3,
			},
			want: 8,
		},
		{
			name: "Negative power",
			args: args{
				x: 2,
				n: -3,
			},
			want: 0.125,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XPower(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("XPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestWord(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Find longest word",
			args: args{
				sentence: "Find the longest word in an English sentence",
			},
			want: "sentence",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestWord(tt.args.sentence); got != tt.want {
				t.Errorf("LongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTernarySearch(t *testing.T) {
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
			name: "Ternary search",
			args: args{
				x:   7,
				arr: []int{1, 2, 3, 5, 6, 7, 8},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TernarySearch(tt.args.x, tt.args.arr); got != tt.want {
				t.Errorf("TernarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClosestPair(t *testing.T) {
	type args struct {
		arr []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{
			name: "Closest pair",
			args: args{
				arr: []float64{1, 5, 9, 3, 10},
			},
			want:  9,
			want1: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ClosestPair(tt.args.arr)
			if got != tt.want {
				t.Errorf("ClosestPair() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ClosestPair() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGreaterThanPreviousTermsSum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "All greater items",
			args: args{
				arr: []int{1, 4, 9, 7, 14, 5, 41},
			},
			want: []int{4, 9, 41},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GreaterThanPreviousTermsSum(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThanPreviousTermsSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
