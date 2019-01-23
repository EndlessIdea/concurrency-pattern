package concurrency_pattern

import (
	"testing"
)

func BenchmarkRepeatGenerator(b *testing.B) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for range TakeNGenerator(done, RepeatGenerator(done, "hello ", "world"), b.N){}
}
