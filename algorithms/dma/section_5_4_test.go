package dma

import "testing"

func TestMultiply(t *testing.T) {
	type args struct {
		n int64
		x int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "6 * 8",
			args: args{
				n: 6,
				x: 8,
			},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecurMultiply(tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("RecurMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecurSumN(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "First 10 int",
			args: args{
				n: 10,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecurSumN(tt.args.n); got != tt.want {
				t.Errorf("RecurSumN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecurSumOddN(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "First 1 odd",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "First 2 odd",
			args: args{
				n: 2,
			},
			want: 4,
		},
		{
			name: "First 3 odd",
			args: args{
				n: 3,
			},
			want: 9,
		},
		{
			name: "First 4 odd",
			args: args{
				n: 4,
			},
			want: 16,
		},
		{
			name: "First 5 odd",
			args: args{
				n: 5,
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecurSumOddN(tt.args.n); got != tt.want {
				t.Errorf("RecurSumOddN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecurMax(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "{1, 8, 9, 3, 5, 7, 2}",
			args: args{
				arr: []int{1, 8, 9, 3, 5, 7, 2},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecurMax(tt.args.arr); got != tt.want {
				t.Errorf("RecurMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
