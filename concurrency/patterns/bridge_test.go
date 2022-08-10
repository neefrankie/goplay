package patterns

import (
	"testing"
)

func genVals() <-chan <-chan interface{} {
	chanStream := make(chan (<-chan interface{}))
	go func() {
		defer close(chanStream)
		for i := 0; i < 10; i++ {
			stream := make(chan interface{}, 1)
			stream <- i
			close(stream)
			chanStream <- stream
		}
	}()

	return chanStream
}

func TestBridge(t *testing.T) {
	for v := range Bridge(nil, genVals()) {
		t.Logf("%v ", v)
	}
}
