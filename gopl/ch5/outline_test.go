package ch5

import (
	"bytes"
	"testing"
)

func Test_htmlOutline(t *testing.T) {
	err := htmlOutline(bytes.NewReader(getRustHome()))
	if err != nil {
		t.Error(err)
	}
}
