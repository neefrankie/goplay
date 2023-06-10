package main

import (
	"testing"
)

func Test_connect(t *testing.T) {
	_, err := connect(buildDSN())
	if err != nil {
		t.Fatal(err)
	}
}
