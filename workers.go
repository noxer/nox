package nox

import (
	"runtime"
)

type Result[T any] struct {
	r T
	e error
}

func NewResult[T any](r T, e error) Result[T] {
	return Result[T]{
		r: r,
		e: e,
	}
}

func (r Result[T]) OK() bool {
	return r.e == nil
}

func ConcurrentNow[T, S any](tasks []T, f func(T) S) []S {
	return ConcurrentNowN(tasks, f, runtime.NumCPU())
}

func ConcurrentNowN[T, S any](tasks []T, f func(T) S, workers int) []S {

}

func ConcurrentN[T, S any](tasks []T, f func(T) S, workers int) Future[S] {

}
