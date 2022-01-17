package set

import (
	. "github.com/noxer/nox/dot"
	"github.com/noxer/nox/maps"
	"github.com/noxer/nox/slice"
)

// Set offers set functionality. A set is a data structure that only contains unique, unsorted values.
type Set[T comparable] map[T]struct{}

// New creates a new set from a list of values.
func New[T comparable](from ...T) Set[T] {
	s := make(Set[T], len(from))
	for _, e := range from {
		s[e] = struct{}{}
	}
	return s
}

// Union takes two sets and returns a new set that contains all values from
// both sets.
func Union[T comparable](a, b Set[T]) Set[T] {
	u := make(Set[T])
	for k := range b {
		u[k] = struct{}{}
	}
	for k := range a {
		u[k] = struct{}{}
	}
	return u
}

// Intersection takes two sets and returns a new set containing only the values
// found in both sets.
func Intersection[T comparable](a, b Set[T]) Set[T] {
	u := make(Set[T])
	for k := range a {
		if _, ok := b[k]; ok {
			u[k] = struct{}{}
		}
	}
	return u
}

// Len returns the number of elements in a Set.
func (s Set[T]) Len() int {
	return len(s)
}

// Any returns any element from the set (useful if the set has only one value).
func (s Set[T]) Any() T {
	for k := range s {
		return k
	}
	return Default[T]()
}

// Put an element into the set.
func (s Set[T]) Put(e T) {
	s[e] = struct{}{}
}

// Has checks if the set contains a certain element.
func (s Set[T]) Has(e T) bool {
	_, ok := s[e]
	return ok
}

// Delete removes an element from the set.
func (s Set[T]) Delete(e T) {
	delete(s, e)
}

// Add a second set to the first one. This is like Union but the elements are
// added to the set this message was called on.
func (s Set[T]) Add(o Set[T]) {
	for k := range o {
		s[k] = struct{}{}
	}
}

// Substract removes all elements in set o from this set.
func (s Set[T]) Subtract(o Set[T]) {
	for k := range o {
		delete(s, k)
	}
}

// Slice returns the unsorted list of elements of this set.
func (s Set[T]) Slice() []T {
	return maps.Keys(s)
}

// Enumerate creates an enumerable from this set.
func (s Set[T]) Enumerate() Enumerable[T] {
	return slice.Enumerate(s.Slice())
}
