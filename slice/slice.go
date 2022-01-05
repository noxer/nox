package slice

import (
	"constraints"

	. "github.com/noxer/nox/dot"
)

func Append[T any](sl []T, items ...T) []T {
	return append(sl, items...)
}

func Insert[T constraints.Ordered](sl []T, e T) []T {
	index := Search(sl, e)
	sl = append(sl, e)

	if index == len(sl)-1 {
		return sl
	}

	copy(sl[index+1:], sl[index:])
	sl[index] = e
	return sl
}

func InsertBy[T any, P constraints.Ordered](sl []T, f func(T) P, e T) []T {
	index := SearchBy(sl, f, e)
	sl = append(sl, e)

	if index == len(sl)-1 {
		return sl
	}

	copy(sl[index+1:], sl[index:])
	sl[index] = e
	return sl
}

func RemoveFirst[T comparable](sl []T, e T) []T {
	for i, t := range sl {
		if t == e {
			return RemoveIndex(sl, i)
		}
	}

	return sl
}

func RemoveAll[T comparable](sl []T, e T) []T {
	newSl := sl[:0]

	for _, t := range sl {
		if t != e {
			newSl = append(newSl, t)
		}
	}

	return newSl
}

func RemoveIndex[T any](sl []T, index int) []T {
	copy(sl[index:], sl[index+1:])
	return sl[:len(sl)-1]
}

func Map[T, S any](sl []T, f func(T) S) []S {
	out := make([]S, len(sl))

	for i, t := range sl {
		out[i] = f(t)
	}

	return out
}

func MapSelf[T any](sl []T, f func(T) T) {
	for i, t := range sl {
		sl[i] = f(t)
	}
}

func Each[T any](sl []T, f func(T)) {
	for _, t := range sl {
		f(t)
	}
}

func Sum[T Number](sl ...T) T {
	if len(sl) == 0 {
		return Default[T]()
	}

	sum := sl[0]
	for _, t := range sl[1:] {
		sum += t
	}

	return sum
}

func Prod[T Number](sl ...T) T {
	if len(sl) == 0 {
		return Default[T]()
	}

	sum := sl[0]
	for _, t := range sl[1:] {
		sum *= t
	}

	return sum
}

type sliceEnumerator[T any] struct {
	sl []T
}

func (e *sliceEnumerator[T]) Next() bool {
	if len(e.sl) == 0 {
		return false
	}

	e.sl = e.sl[1:]
	return len(e.sl) > 0
}

func (e *sliceEnumerator[T]) Value() T {
	if len(e.sl) > 0 {
		return e.sl[0]
	}

	return Default[T]()
}

func Enumerate[T any](sl []T) Enumerable[T] {
	return &sliceEnumerator[T]{sl}
}
