package nox

import (
	"runtime"
	"sync"

	"github.com/noxer/nox/math"
)

// ConcurrentNow takes a list of tasks and starts a worker pool to process them
// with f.
func ConcurrentNow[T, S any](tasks []T, f func(T) S) []S {
	return ConcurrentNowN(tasks, f, runtime.NumCPU())
}

// ConcurrentNowN takes a list of tasks and starts a worker pool of size
// workers to process them with f.
func ConcurrentNowN[T, S any](tasks []T, f func(T) S, workers int) []S {
	switch len(tasks) {
	case 0:
		return nil

	case 1:
		return []S{f(tasks[0])}
	}

	workers = math.Min(len(tasks), workers)

	wg := &sync.WaitGroup{}
	wg.Add(workers)

	in := make(chan int, workers)
	results := make([]S, len(tasks))
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for index := range in {
				results[index] = f(tasks[index])
			}
		}()
	}

	go func() {
		for i := range tasks {
			in <- i
		}
		close(in)
	}()

	return results
}
