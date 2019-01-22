package concurrency_pattern

import (

)

// Generating channels by values
func CommonGenerator(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	stream := make(chan interface{})

	go func() {
		defer close(stream)

		for _, value := range values {
			select {
			case <-done:
				return
			case stream <- value:
			}
		}
	}()

	return stream
}

// Generating values repeatedly
func RepeatGenerator(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	stream := make(chan interface{})

	go func() {
		defer close(stream)

		for {
			for _, value := range values {
				select {
				case <-done:
					return
				case stream <- value:
				}
			}
		}
	}()

	return stream
}

// Get n value from inputStream, then stop and return
func TakeNGenerator(done <-chan interface{}, inputStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})

	go func() {
		defer close(takeStream)

		//cnt := 0
		//for v := range inputStream {
		//	select {
		//	case <-done:
		//		return
		//	case takeStream <- v:
		//	}
		//	cnt++
		//	if cnt >= num {
		//		return
		//	}
		//}
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- inputStream:
			}
		}
	}()

	return takeStream
}
