package dynamic

import "testing"

func TestRecurKnapsack(t *testing.T) {
	type args struct {
		capacity int
		items    []sackItem
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Naive knapsack",
			args: args{
				capacity: 17,
				items:    knapsackItems,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NaiveKnapsack(tt.args.capacity, tt.args.items)

			//if got != tt.want {
			//	t.Errorf("NaiveKnapsack() = %v, want %v", got, tt.want)
			//}

			t.Logf("%d\n", got)
		})
	}
}
