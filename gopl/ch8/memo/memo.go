package memo

import "sync"

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]*entry),
	}
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]

	if e == nil {
		// This is the first request to this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{
			res:   result{},
			ready: make(chan struct{}),
		}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // broadcast for this key.
	} else {
		// If there was an existing entry, its value
		// is not necessarily ready yet; another goroutine
		// could still be calling the slow function f.
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready //wait for ready condition.
	}

	return e.res.value, e.res.err
}
