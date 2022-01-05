package set

import (
	"github.com/noxer/nox/channel"
	. "github.com/noxer/nox/dot"
	"github.com/noxer/nox/maps"
)

type Set[T comparable] map[T]struct{}

func New[T comparable](from ...T) Set[T] {
	s := make(Set[T], len(from))
	for _, e := range from {
		s[e] = struct{}{}
	}
	return s
}

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

func Intersection[T comparable](a, b Set[T]) Set[T] {
	u := make(Set[T])
	for k := range a {
		if _, ok := b[k]; ok {
			u[k] = struct{}{}
		}
	}
	return u
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Any() T {
	for k := range s {
		return k
	}
	return Default[T]()
}

func (s Set[T]) Put(e T) {
	s[e] = struct{}{}
}

func (s Set[T]) Has(e T) bool {
	_, ok := s[e]
	return ok
}

func (s Set[T]) Delete(e T) {
	delete(s, e)
}

func (s Set[T]) Add(o Set[T]) {
	for k := range o {
		s[k] = struct{}{}
	}
}

func (s Set[T]) Subtract(o Set[T]) {
	for k := range o {
		delete(s, k)
	}
}

func (s Set[T]) Slice() []T {
	return maps.Keys(s)
}

func (s Set[T]) Enumerate() Enumerable[T] {
	ch := make(chan T)
	go func() {
		for k := range s {
			ch <- k
		}
		close(ch)
	}()

	return channel.Enumerate(ch)
}
