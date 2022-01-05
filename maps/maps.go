package maps

import (
	"constraints"

	. "github.com/noxer/nox/dot"
	"github.com/noxer/nox/slice"
	"github.com/noxer/nox/tuple"
)

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func SortedKeys[K constraints.Ordered, V any](m map[K]V) []K {
	keys := Keys(m)
	slice.Sort(keys)
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func SortedValues[K comparable, V constraints.Ordered](m map[K]V) []V {
	values := Values(m)
	slice.Sort(values)
	return values
}

func ValuesBySortedKeys[K constraints.Ordered, V any](m map[K]V) []V {
	keys := SortedKeys(m)
	values := make([]V, len(m))
	for i, k := range keys {
		values[i] = m[k]
	}
	return values
}

func KeyValues[K comparable, V any](m map[K]V) []tuple.T2[K, V] {
	tuples := make([]tuple.T2[K, V], 0, len(m))
	for k, v := range m {
		tuples = append(tuples, tuple.T2[K, V]{A: k, B: v})
	}
	return tuples
}

func SortedKeyValues[K constraints.Ordered, V any](m map[K]V) []tuple.T2[K, V] {
	keyValues := KeyValues(m)
	slice.SortBy(keyValues, func(t tuple.T2[K, V]) K { return t.A })
	return keyValues
}

func SumKeys[K Number, V any](m map[K]V) (sum K) {
	for k := range m {
		sum += k
	}
	return
}

func SumValues[K comparable, V Number](m map[K]V) (sum V) {
	for _, v := range m {
		sum += v
	}
	return
}

func ProdKeys[K Number, V any](m map[K]V) (sum K) {
	for k := range m {
		sum *= k
	}
	return
}

func ProdValues[K comparable, V Number](m map[K]V) (sum V) {
	for _, v := range m {
		sum *= v
	}
	return
}
