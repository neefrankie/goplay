package array

import (
	"testing"
)

func TestCountNumbers1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "555",
			args: args{
				n: 555,
			},
			want: 3,
		},
		{
			name: "12",
			args: args{
				n: 12,
			},
			want: 2,
		},
		{
			name: "7896",
			args: args{
				n: 7896,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountNumbers1(tt.args.n); got != tt.want {
				t.Errorf("CountNumbers1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountNumbers2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "555",
			args: args{
				n: 555,
			},
			want: 3,
		},
		{
			name: "12",
			args: args{
				n: 12,
			},
			want: 2,
		},
		{
			name: "7896",
			args: args{
				n: 7896,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountNumbers2(tt.args.n); got != tt.want {
				t.Errorf("CountNumbers2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountNumber3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "555",
			args: args{
				n: 555,
			},
			want: 3,
		},
		{
			name: "12",
			args: args{
				n: 12,
			},
			want: 2,
		},
		{
			name: "7896",
			args: args{
				n: 7896,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountDigits(tt.args.n); got != tt.want {
				t.Errorf("CountNumber3() = %v, want %v", got, tt.want)
			}
		})
	}
}
