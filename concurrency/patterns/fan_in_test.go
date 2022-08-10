package patterns

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		r := n % i
		if r == 0 {
			return false
		}
	}

	return true
}

func primeFinder(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
	primeStream := make(chan interface{})
	go func() {
		defer close(primeStream)
		for v := range intStream {
			select {
			case <-done:
				return
			default:
				if isPrime(v) {
					primeStream <- v
				}
			}
		}
	}()

	return primeStream
}

func getRandInt() interface{} {
	return rand.Intn(1000000000)
}

func Test_isPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "5 is prime",
			args: args{
				n: 5,
			},
			want: true,
		},
		{
			name: "11 is prime",
			args: args{
				n: 11,
			},
			want: true,
		},
		{
			name: "13 is prime",
			args: args{
				n: 13,
			},
			want: true,
		},
		{
			name: "17 is prime",
			args: args{
				n: 17,
			},
			want: true,
		},
		{
			name: "19 is prime",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "23 is prime",
			args: args{
				n: 23,
			},
			want: true,
		},
		{
			name: "30 is not prime",
			args: args{
				n: 30,
			},
			want: false,
		},
		{
			name: "31 is prime",
			args: args{
				n: 31,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.n); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_primeFinder(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, RepeatFn(done, getRandInt))

	t.Logf("Primes:\n")
	for prime := range Take(done, primeFinder(done, randIntStream), 10) {
		t.Logf("\t%d\n", prime)
	}

	t.Logf("Search took: %v\n", time.Since(start))
}

func TestFanIn(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, RepeatFn(done, getRandInt))

	numFinders := runtime.NumCPU()
	t.Logf("Spinning up %d prime finders.\n", numFinders)

	finders := make([]<-chan interface{}, numFinders)
	t.Log("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	for prime := range Take(done, FanIn(done, finders...), 10) {
		t.Logf("\t%d\n", prime)
	}

	t.Logf("Search took: %v", time.Since(start))
}
