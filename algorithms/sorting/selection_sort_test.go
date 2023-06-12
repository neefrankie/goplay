package sorting

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "Selection sort",
			args: args{
				arr: []int64{3, 5, 4, 1, 2},
			},
			want: []int64{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args.arr)

			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}
