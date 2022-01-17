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

// Map applies a function f to every element of the enumerable e.
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

// Filter lets only the elements pass that f returns true for.
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

// Zip takes two enumerables and returns a enumerable with tuples of them.
func Zip[T, S any](a Enumerable[T], b Enumerable[S]) Enumerable[tuple.T2[T, S]] {
	return &enumZipper[T, S]{a, b}
}

// Sum reads all values from an enumerable and sums them up.
func Sum[T Number](e Enumerable[T]) (sum T) {
	for e.Next() {
		sum += e.Value()
	}
	return
}

// Max reads all values from an enumerable and returns the biggest one.
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

// Min reads all values from an enumerable and returns the smallest one.
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

// Count reads all values from an enumerable and returns the number of values.
func Count[T any](e Enumerable[T]) int {
	i := 0
	for e.Next() {
		i++
	}
	return i
}

// Drain reads all values from an enumerable.
func Drain[T any](e Enumerable[T]) {
	for e.Next() {
	}
}

// Consume reads n values from an enumerable.
func Consume[T any](e Enumerable[T], n int) {
	for i := 0; i < n; i++ {
		if !e.Next() {
			return
		}
	}
}

// Histogram reads all values from an enumerable and returns a map of values and their respective counts.
func Histogram[T comparable, C Number](e Enumerable[T]) map[T]C {
	h := make(map[T]C)
	for e.Next() {
		h[e.Value()]++
	}
	return h
}

// Unique reads all values from an enumerable and returns a set of unique values.
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

// Synchronize makes reading the enumerable thread safe.
func Synchronize[T any](e Enumerable[T]) Enumerable[T] {
	if _, ok := e.(*enumMutex[T]); ok {
		// don't wrap an existing synchronized mutex
		return e
	}

	return &enumMutex[T]{e: e}
}
