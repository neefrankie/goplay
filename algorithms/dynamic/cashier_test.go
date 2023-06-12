package dynamic

import (
	"reflect"
	"testing"
)

func TestCashierChange(t *testing.T) {
	type args struct {
		coins []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				coins: []int{25, 10, 5, 1},
				n:     67,
			},
			want: []int{2, 1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CashierChange(tt.args.coins, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CashierChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
