package concurrency_pattern

import (
	"testing"
	"time"
	"fmt"
)

func TestMergeDone(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-MultiDone(
		sig(time.Second),
		sig(time.Second*2),
		sig(time.Minute),
		sig(time.Hour),
	)
	fmt.Printf("done after %v", time.Since(start))
}