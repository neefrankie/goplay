package config

import (
	"testing"
)

func TestConfig_Get(t *testing.T) {
	c := MustLoad()

	type args struct {
		key string
	}
	tests := []struct {
		name string
		c    Config
		args args
		want string
	}{
		{
			name: "Get Host",
			c:    c,
			args: args{
				key: "MYSQL_HOST",
			},
			want: "127.0.0.1",
		},
		{
			name: "Get Port",
			c:    c,
			args: args{
				key: "MYSQL_PORT",
			},
			want: "3306",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Get(tt.args.key); got != tt.want {
				t.Errorf("Config.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_GetInt(t *testing.T) {
	c := MustLoad()

	type args struct {
		key string
	}
	tests := []struct {
		name string
		c    Config
		args args
		want int
	}{
		{
			name: "Get Port",
			c:    c,
			args: args{
				key: "MYSQL_PORT",
			},
			want: 3306,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetInt(tt.args.key); got != tt.want {
				t.Errorf("Config.GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
