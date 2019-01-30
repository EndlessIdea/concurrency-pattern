package concurrency_pattern

import (
	"fmt"
	"testing"
	"time"
)

func orDoneProducer() <-chan interface{} {
	stream := make(chan interface{})
	go func() {
		defer close(stream)
		for i := 0; i < 5; i++ {
			stream <- i
			time.Sleep(1 * time.Second)
		}
	}()
	return stream
}

func getOrDone() <-chan interface{} {
	done := make(chan interface{})
	go func() {
		defer close(done)
		time.Sleep(3 * time.Second)
	}()
	return done
}

func TestOrDone(t *testing.T) {
	stream := orDoneProducer()
	done := getOrDone()

	msgs := OrDone(done, stream)

	for msg := range msgs {
		fmt.Println(msg)
	}
}
