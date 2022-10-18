package ch5

import (
	"bytes"
	"goplay/gopl/pkg"
	"testing"
)

func getRustHome() []byte {
	b, err := pkg.Fetch("https://www.rust-lang.org/")
	if err != nil {
		panic(err)
	}

	return b
}

func TestFindLinks(t *testing.T) {
	links, err := findLinks(bytes.NewReader(getRustHome()))
	if err != nil {
		t.Error(err)
		return
	}

	for _, link := range links {
		t.Logf("%s\n", link)
	}
}
