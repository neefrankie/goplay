package chapter3

import "testing"

func TestHelloChannels(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Hello channels",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HelloChannels()
		})
	}
}

func TestReadClosedChannel(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Read closed channel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadClosedChannel()
		})
	}
}

func TestRangeOverChannel(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Range over channel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RangeOverChannel()
		})
	}
}

func TestUnblockMultiGoroutines(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Unblock multi goroutines",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UnblockMultiGoroutines()
		})
	}
}

func TestChanOwner(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Channel ownership",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ChanOwner()
		})
	}
}
