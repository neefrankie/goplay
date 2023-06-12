package search

import "testing"

func TestNaiveStringMatcher(t *testing.T) {
	type args struct {
		search  string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "String matcher",
			args: args{
				search:  "eceyeye",
				pattern: "eye",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NaiveStringMatcher(tt.args.search, tt.args.pattern); got != tt.want {
				t.Errorf("NaiveStringMatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
