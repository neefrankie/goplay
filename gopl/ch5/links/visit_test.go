package links

import (
	"bufio"
	"strings"
	"testing"
)

func Test_countWords(t *testing.T) {
	text := "言语与语言处 Speech and Language Processing"
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)

	var n int
	for scanner.Scan() {
		w := scanner.Text()
		t.Logf("%s\n", w)
		n++
	}

	t.Logf("words: %d\n", n)
}
