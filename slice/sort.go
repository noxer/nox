package slice

import (
	"sort"

	"golang.org/x/exp/constraints"
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

// Sort a slice.
func Sort[T constraints.Ordered](sl []T) {
	sort.Sort(sorter[T](sl))
}

// Reverse sort a slice.
func ReverseSort[T constraints.Ordered](sl []T) {
	sort.Sort(sort.Reverse(sorter[T](sl)))
}

// Search an item in a sorted slice and return its index. If the exact item is
// not found the index points to the place it would be.
func Search[T constraints.Ordered](sl []T, el T) int {
	return sort.Search(len(sl), func(i int) bool {
		return sl[i] > el
	})
}

// SearchBy searches a slice by using a mapping function to convert an
// unordered type into an ordered one and returns its index. If the exact item
// is not found the index points to the place it would be.
func SearchBy[T any, P constraints.Ordered](sl []T, f func(T) P, el T) int {
	ref := f(el)

	return sort.Search(len(sl), func(i int) bool {
		return f(sl[i]) > ref
	})
}

// IsSorted checks if the slice is sorted.
func IsSorted[T constraints.Ordered](sl []T) bool {
	return sort.IsSorted(sorter[T](sl))
}

// IsReverseSorted checks if the slice is sorted in reverse.
func IsReverseSorted[T constraints.Ordered](sl []T) bool {
	return sort.IsSorted(sort.Reverse(sorter[T](sl)))
}

// IsSortedBy checks if the slice is sorted by f(e).
func IsSortedBy[T any, P constraints.Ordered](sl []T, f func(T) P) bool {
	return sort.IsSorted(sorterBy[T, P]{sl, f})
}

// IsReverseSorted checks if the slice is sorted in reverse by f(e).
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

// SortBy sorts the slice by f(e).
func SortBy[T any, P constraints.Ordered](sl []T, f func(T) P) {
	sort.Sort(sorterBy[T, P]{sl, f})
}

// StableBy stable sorts the slice by f(e). Stable means, that equal elements
// remain in the same order.
func StableBy[T any, P constraints.Ordered](sl []T, f func(T) P) {
	sort.Stable(sorterBy[T, P]{sl, f})
}

// StableBy stable sorts the slice by f(e) in reverse.
func ReverseSortBy[T any, P constraints.Ordered](sl []T, f func(T) P) {
	sort.Sort(sort.Reverse(sorterBy[T, P]{sl, f}))
}

// ReverseStableBy stable sorts the slice by f(e) in reverse. Stable means, that equal elements
// remain in the same order.
func ReverseStableBy[T any, P constraints.Ordered](sl []T, f func(T) P) {
	sort.Stable(sort.Reverse(sorterBy[T, P]{sl, f}))
}
