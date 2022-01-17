package maps

import (
	"constraints"

	. "github.com/noxer/nox/dot"
	"github.com/noxer/nox/slice"
	"github.com/noxer/nox/tuple"
)

// Keys returns an unsorted slice of the keys of m.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// SortedKeys returns a sorted slice of the keys of m.
func SortedKeys[K constraints.Ordered, V any](m map[K]V) []K {
	keys := Keys(m)
	slice.Sort(keys)
	return keys
}

// Values returns an unsorted slice of values of m.
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// SortedValues returns a sorted slice of values of m.
func SortedValues[K comparable, V constraints.Ordered](m map[K]V) []V {
	values := Values(m)
	slice.Sort(values)
	return values
}

// ValuesBySortedKeys returns a slice of values of m sorted by the keys of m.
func ValuesBySortedKeys[K constraints.Ordered, V any](m map[K]V) []V {
	keys := SortedKeys(m)
	values := make([]V, len(m))
	for i, k := range keys {
		values[i] = m[k]
	}
	return values
}

// KeyValues returns an unsorted slice of Key-Value pairs.
func KeyValues[K comparable, V any](m map[K]V) []tuple.T2[K, V] {
	tuples := make([]tuple.T2[K, V], 0, len(m))
	for k, v := range m {
		tuples = append(tuples, tuple.T2[K, V]{A: k, B: v})
	}
	return tuples
}

// SortedKeyValues returns a slice of Key-Value pairs sorted by the keys.
func SortedKeyValues[K constraints.Ordered, V any](m map[K]V) []tuple.T2[K, V] {
	keyValues := KeyValues(m)
	slice.SortBy(keyValues, func(t tuple.T2[K, V]) K { return t.A })
	return keyValues
}

// SumKeys sums up the keys of m.
func SumKeys[K Number, V any](m map[K]V) (sum K) {
	for k := range m {
		sum += k
	}
	return
}

// SumValues sums up the values of m.
func SumValues[K comparable, V Number](m map[K]V) (sum V) {
	for _, v := range m {
		sum += v
	}
	return
}

// ProdKeys multiplies all the keys of m.
func ProdKeys[K Number, V any](m map[K]V) (sum K) {
	for k := range m {
		sum *= k
	}
	return
}

// ProdValues multiplies all the values of m.
func ProdValues[K comparable, V Number](m map[K]V) (sum V) {
	for _, v := range m {
		sum *= v
	}
	return
}
