package patterns

import "sync"

// FanIn multiplexes or joins together multiple streams of
// data into a single stream.
func FanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	// Create a wait group so that we can wait until all channels have been drained.
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})

	// A function read from the passed channel,
	// and pass the value read onto the multiplexedStream channel.
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
