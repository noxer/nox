package slice

import (
	"constraints"
	"sort"
)

type sorter[T constraints.Ordered] []T

func (s sorter[T]) Len() int {
	return len(s)
}

func (s sorter[T]) Less(a, b int) bool {
	return s[a] < s[b]
}

func (s sorter[T]) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

func Sort[T constraints.Ordered](sl []T) {
	sort.Sort(sorter[T](sl))
}

func Stable[T constraints.Ordered](sl []T) {
	sort.Stable(sorter[T](sl))
}

func ReverseSort[T constraints.Ordered](sl []T) {
	sort.Sort(sort.Reverse(sorter[T](sl)))
}

func ReverseStable[T constraints.Ordered](sl []T) {
	sort.Stable(sort.Reverse(sorter[T](sl)))
}

func Search[T constraints.Ordered](sl []T, el T) int {
	return sort.Search(len(sl), func(i int) bool {
		return sl[i] > el
	})
}

func SearchBy[T any, P constraints.Ordered](sl []T, f func(T) P, el T) int {
	ref := f(el)

	return sort.Search(len(sl), func(i int) bool {
		return f(sl[i]) > ref
	})
}

func IsSorted[T constraints.Ordered](sl []T) bool {
	return sort.IsSorted(sorter[T](sl))
}

func IsReverseSorted[T constraints.Ordered](sl []T) bool {
	return sort.IsSorted(sort.Reverse(sorter[T](sl)))
}

func IsSortedBy[T any, P constraints.Ordered](sl []T, f func(T) P) bool {
	return sort.IsSorted(sorterBy[T, P]{sl, f})
}

func IsReverseSortedBy[T any, P constraints.Ordered](sl []T, f func(T) P) bool {
	return sort.IsSorted(sort.Reverse(sorterBy[T, P]{sl, f}))
}

type sorterBy[T any, P constraints.Ordered] struct {
	sl []T
	f  func(T) P
}

func (s sorterBy[T, P]) Len() int {
	return len(s.sl)
}

func (s sorterBy[T, P]) Less(a, b int) bool {
	return s.f(s.sl[a]) < s.f(s.sl[b])
}

func (s sorterBy[T, P]) Swap(a, b int) {
	s.sl[a], s.sl[b] = s.sl[b], s.sl[a]
}

func SortBy[T any, P constraints.Ordered](sl []T, f func(T) P) {
	sort.Sort(sorterBy[T, P]{sl, f})
}

func StableBy[T any, P constraints.Ordered](sl []T, f func(T) P) {
	sort.Stable(sorterBy[T, P]{sl, f})
}

func ReverseSortBy[T any, P constraints.Ordered](sl []T, f func(T) P) {
	sort.Sort(sort.Reverse(sorterBy[T, P]{sl, f}))
}

func ReverseStableBy[T any, P constraints.Ordered](sl []T, f func(T) P) {
	sort.Stable(sort.Reverse(sorterBy[T, P]{sl, f}))
}
