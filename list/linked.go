package list

import (
	"fmt"
)

type Linked[T any] struct {
	size  int
	first *link[T]
}

type link[T any] struct {
	val  T
	next *link[T]
}

func New[T any](from ...T) *Linked[T] {
	head := &Linked[T]{size: len(from)}
	cur := &head.first
	for _, e := range from {
		*cur = &link[T]{val: e}
		cur = &(*cur).next
	}
	return head
}

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

func (l *Linked[T]) Get(i int) T {
	if i >= l.size {
		panic(fmt.Sprintf("index out of range %d of %d", i, l.size))
	}
	return l.byIndex(i).val
}
