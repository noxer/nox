package channel

import (
	. "github.com/noxer/nox/dot"
)

func TryGet[T any](ch chan T) Optional[T] {
	select {
	case t := <-ch:
		return Success(t)
	default:
		return Failure[T]()
	}
}

func TryPut[T any](ch chan T, t T) bool {
	select {
	case ch <- t:
		return true
	default:
		return false
	}
}

func Enumerate[T any](ch chan T) Enumerable[T] {
	return &chanEnumerator[T]{ch: ch}
}

type chanEnumerator[T any] struct {
	ch  chan T
	cur T
}

func (e *chanEnumerator[T]) Next() (ok bool) {
	e.cur, ok = <-e.ch
	return
}

func (e *chanEnumerator[T]) Value() T {
	return e.cur
}
