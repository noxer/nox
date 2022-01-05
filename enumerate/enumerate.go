package enumerate

import (
	"constraints"
	"sync"

	. "github.com/noxer/nox/dot"
	"github.com/noxer/nox/set"
	"github.com/noxer/nox/tuple"
)

type enumMapper[T, S any] struct {
	e Enumerable[T]
	f func(T) S
}

func (e *enumMapper[T, S]) Next() bool {
	return e.e.Next()
}

func (e *enumMapper[T, S]) Value() S {
	return e.f(e.e.Value())
}

func Map[T, S any](e Enumerable[T], f func(T) S) Enumerable[S] {
	return &enumMapper[T, S]{e, f}
}

type enumFilter[T any] struct {
	e Enumerable[T]
	f func(T) bool
}

func (e *enumFilter[T]) Next() bool {
	for e.e.Next() {
		if e.f(e.e.Value()) {
			return true
		}
	}

	return false
}

func (e *enumFilter[T]) Value() T {
	return e.e.Value()
}

func Filter[T any](e Enumerable[T], f func(T) bool) Enumerable[T] {
	return &enumFilter[T]{e, f}
}

type enumZipper[T, S any] struct {
	a Enumerable[T]
	b Enumerable[S]
}

func (e *enumZipper[T, S]) Next() bool {
	return e.a.Next() && e.b.Next()
}

func (e *enumZipper[T, S]) Value() tuple.T2[T, S] {
	return tuple.T2[T, S]{A: e.a.Value(), B: e.b.Value()}
}

func Zip[T, S any](a Enumerable[T], b Enumerable[S]) Enumerable[tuple.T2[T, S]] {
	return &enumZipper[T, S]{a, b}
}

func Sum[T Number](e Enumerable[T]) (sum T) {
	for e.Next() {
		sum += e.Value()
	}
	return
}

func Max[T constraints.Ordered](e Enumerable[T]) Optional[T] {
	if !e.Next() {
		return Failure[T]()
	}
	max := e.Value()

	for e.Next() {
		if val := e.Value(); val > max {
			max = val
		}
	}
	return Success(max)
}

func Min[T constraints.Ordered](e Enumerable[T]) Optional[T] {
	if !e.Next() {
		return Failure[T]()
	}
	min := e.Value()

	for e.Next() {
		if val := e.Value(); val < min {
			min = val
		}
	}
	return Success(min)
}

func Count[T any](e Enumerable[T]) int {
	i := 0
	for e.Next() {
		i++
	}
	return i
}

func Drain[T any](e Enumerable[T]) {
	for e.Next() {
	}
}

func Consume[T any](e Enumerable[T], n int) {
	for i := 0; i < n; i++ {
		if !e.Next() {
			return
		}
	}
}

func Histogram[T comparable, C Number](e Enumerable[T]) map[T]C {
	h := make(map[T]C)
	for e.Next() {
		h[e.Value()]++
	}
	return h
}

func Unique[T comparable](e Enumerable[T]) set.Set[T] {
	s := make(set.Set[T])
	for e.Next() {
		s.Put(e.Value())
	}
	return s
}

type enumMutex[T any] struct {
	m sync.RWMutex
	e Enumerable[T]
}

func (e *enumMutex[T]) Next() bool {
	e.m.Lock()
	defer e.m.Unlock()

	return e.e.Next()
}

func (e *enumMutex[T]) Value() T {
	e.m.RLock()
	defer e.m.RUnlock()

	return e.e.Value()
}

func Synchronize[T any](e Enumerable[T]) Enumerable[T] {
	if _, ok := e.(*enumMutex[T]); ok {
		return e
	}

	return &enumMutex[T]{e: e}
}
