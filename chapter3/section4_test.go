package chapter3

import "testing"

func TestSelectExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Select block",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectExample()
		})
	}
}

func TestSelectDistribution(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Select distribution",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectDistribution()
		})
	}
}

func TestSelectTimeout(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Select timeout",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectTimeout()
		})
	}
}

func TestSelectSignal(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Signal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectSignal()
		})
	}
}
