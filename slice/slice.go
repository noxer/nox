package slice

import (
	"constraints"

	. "github.com/noxer/nox/dot"
)

// MakeRoom extends the length of sl by add items. It will allocate a new slice
// if cap(sl) is too small. If not it will reslice sl to include the additional
// items. This works not unlike `append` but doesn't copy in any new elements.
func MakeRoom[T any](sl []T, add int) []T {
	available := cap(sl) - len(sl)
	if available >= add {
		return sl[:len(sl)+add]
	}

	newSl := make([]T, len(sl)+add)
	copy(newSl, sl)

	return newSl
}

// Insert item e into sl at position i.
func Insert[T any](sl []T, e T, i int) []T {
	sl = MakeRoom(sl, 1)

	if i != len(sl)-1 {
		copy(sl[i+1:], sl[i:])
	}

	sl[i] = e
	return sl
}

// InsertSorted an element into a sorted list.
func InsertSorted[T constraints.Ordered](sl []T, e T) []T {
	index := Search(sl, e)
	return Insert(sl, e, index)
}

// InsertSorted an element into a sorted list by f(e).
func InsertSortedBy[T any, P constraints.Ordered](sl []T, f func(T) P, e T) []T {
	index := SearchBy(sl, f, e)
	return Insert(sl, e, index)
}

// RemoveFirst removes the first occurrance of e in sl.
func RemoveFirst[T comparable](sl []T, e T) []T {
	for i, t := range sl {
		if t == e {
			return RemoveIndex(sl, i)
		}
	}

	return sl
}

// RemoveAll instances of e from sl.
func RemoveAll[T comparable](sl []T, e T) []T {
	newSl := sl[:0]

	for _, t := range sl {
		if t != e {
			newSl = append(newSl, t)
		}
	}

	return newSl
}

// RemoveIndex removes the element at index i.
func RemoveIndex[T any](sl []T, i int) []T {
	copy(sl[i:], sl[i+1:])
	return sl[:len(sl)-1]
}

// Map applies the function f to all elements in sl and returns a new slice
// with the results.
func Map[T, S any](sl []T, f func(T) S) []S {
	out := make([]S, len(sl))

	for i, t := range sl {
		out[i] = f(t)
	}

	return out
}

// MapSelf applies the function f to all elements in sl and puts the result into sl.
func MapSelf[T any](sl []T, f func(T) T) {
	for i, t := range sl {
		sl[i] = f(t)
	}
}

// Each calls f for each element in sl.
func Each[T any](sl []T, f func(T)) {
	for _, t := range sl {
		f(t)
	}
}

// Sum all the elements in sl.
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

// Prod multiplies all elements of sl.
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

// Enumerate creates an enumerator for sl.
func Enumerate[T any](sl []T) Enumerable[T] {
	return &sliceEnumerator[T]{sl}
}
