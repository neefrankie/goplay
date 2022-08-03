package basics

import "testing"

func TestWaitComplete(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Wait to complete",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WaitComplete()
		})
	}
}

func TestWaitCompleteN(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Wait an array of goroutines",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WaitCompleteN()
		})
	}
}

func TestMutualExclusion(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test Mutex",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MutualExclusion()
		})
	}
}

func TestReadWriteMutex(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test Read Write Mutex",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadWriteMutex()
		})
	}
}
