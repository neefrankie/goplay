package basics

import "testing"

func TestSayHello(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test SayHello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SayHello()
		})
	}
}

func TestSalutation(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test Salutation",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Salutation()
		})
	}
}

func TestSalutations(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Captured value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Salutations()
		})
	}
}
