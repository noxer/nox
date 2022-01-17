package list

import (
	"fmt"

	. "github.com/noxer/nox/dot"
)

// Linked defines a generic, single-linked list
type Linked[T any] struct {
	size  int
	first *link[T]
}

type link[T any] struct {
	val  T
	next *link[T]
}

// New creates a new single-linked list from a number of elements
func New[T any](from ...T) *Linked[T] {
	head := &Linked[T]{size: len(from)}
	cur := &head.first
	for _, e := range from {
		*cur = &link[T]{val: e}
		cur = &(*cur).next
	}
	return head
}

// Len returns the lenght of the list. Complexity: O(1).
func (l *Linked[T]) Len() int {
	return l.size
}

func (l *Linked[T]) byIndex(index int) *link[T] {
	lnk := l.first
	for i := 0; i < index; i++ {
		if lnk == nil {
			return nil
		}
		lnk = lnk.next
	}
	return lnk
}

func (l *Linked[T]) last() **link[T] {
	last := &l.first
	for (*last).next != nil {
		last = &(*last).next
	}
	return last
}

// Get returns the element at index i in the list. Complexity: O(n).
func (l *Linked[T]) Get(i int) Result[T] {
	if i >= l.size {
		return Err[T](fmt.Errorf("index out of range %d of %d", i, l.size))
	}
	return OK(l.byIndex(i).val)
}
