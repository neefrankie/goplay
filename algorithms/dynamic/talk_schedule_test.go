package dynamic

import (
	"reflect"
	"testing"
)

func TestScheduleTalks(t *testing.T) {
	type args struct {
		talks []TalkSchedule
	}
	tests := []struct {
		name string
		args args
		want []TalkSchedule
	}{
		{
			name: "Should have last 2",
			args: args{
				talks: []TalkSchedule{
					{
						StartTime: 800,
						EndTime:   1200,
					},
					{
						StartTime: 900,
						EndTime:   1000,
					},
					{
						StartTime: 1100,
						EndTime:   1200,
					},
				},
			},
			want: []TalkSchedule{
				{
					StartTime: 900,
					EndTime:   1000,
				},
				{
					StartTime: 1100,
					EndTime:   1200,
				},
			},
		},
		{
			name: "Should have first and last",
			args: args{
				talks: []TalkSchedule{
					{
						StartTime: 800,
						EndTime:   915,
					},
					{
						StartTime: 900,
						EndTime:   1000,
					},
					{
						StartTime: 945,
						EndTime:   1100,
					},
				},
			},
			want: []TalkSchedule{
				{
					StartTime: 800,
					EndTime:   915,
				},
				{
					StartTime: 945,
					EndTime:   1100,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ScheduleTalks(tt.args.talks)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScheduleTalks() = %v, want %v", got, tt.want)
				return
			}

			t.Logf("Talk schedules: %v", got)
		})
	}
}
