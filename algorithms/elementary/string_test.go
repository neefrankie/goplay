package elementary

import (
	"testing"
)

func TestStringSearch(t *testing.T) {
	type args struct {
		src    string
		target string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "All substrings",
			args: args{
				src:    "Library functions, all too often, cannot guarantee to provider the best performance for all applications",
				target: "to",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringSearch(tt.args.src, tt.args.target)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("StringSearch() = %v, want %v", got, tt.want)
			//}

			t.Logf("%v", got)
		})
	}
}
