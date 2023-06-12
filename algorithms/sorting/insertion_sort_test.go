package sorting

import (
	"reflect"
	"testing"
)

func TestInsertionSortInt(t *testing.T) {
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "Insertion sort 1",
			args: args{
				arr: []int64{77, 99, 44, 55, 22, 88, 11, 0, 66, 33},
			},
			want: []int64{0, 11, 22, 33, 44, 55, 66, 77, 88, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
