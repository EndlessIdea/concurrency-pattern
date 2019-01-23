package concurrency_pattern

import (
	"runtime"
	"sync"
)

func FanInMine(done <-chan interface{}, inputStreams ...<-chan interface{}) <-chan interface{} {
	outputStream := make(chan interface{})

	flow := func(done <-chan interface{}, inputStream <-chan interface{}, wg *sync.WaitGroup) {
		defer wg.Done()
		select {
		case <-done:
			return
		case outputStream <- inputStream:
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(inputStreams))
	for _, inputStream := range inputStreams {
		flow(done, inputStream, &wg)
	}

	go func() {
		wg.Wait()
		close(outputStream)
	}()

	return outputStream
}

func FanIn(done <-chan interface{}, inputStreams ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	outputStream := make(chan interface{})

	multiplex := func(inputStream <-chan interface{}) {
		defer wg.Done()
		for i := range inputStream {
			select {
			case <-done:
				return
			case outputStream <- i:
			}
		}
	}

	wg.Add(len(inputStreams))
	for _, inputStream := range inputStreams {
		go multiplex(inputStream)
	}

	go func() {
		wg.Wait()
		close(outputStream)
	}()

	return outputStream
}

type Worker func(done <-chan interface{}, task <-chan interface{}) <-chan interface{}

func FanOut(done <-chan interface{}, task <-chan interface{}, fn Worker) []<-chan interface{} {
	processCnt := runtime.NumCPU()
	workers := make([]<-chan interface{}, processCnt)
	for i := 0; i < processCnt; i++ {
		workers[i] = fn(done, task)
	}
	return workers
}
