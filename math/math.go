package math

import "constraints"

// Min returns the smallest item from a list of items.
func Min[T constraints.Ordered](vs ...T) (min T) {
	if len(vs) == 0 {
		return
	}

	min = vs[0]
	for _, v := range vs[1:] {
		if v < min {
			min = v
		}
	}

	return
}

// Max returns the biggest item from a list of items.
func Max[T constraints.Ordered](vs ...T) (max T) {
	if len(vs) == 0 {
		return
	}

	max = vs[0]
	for _, v := range vs[1:] {
		if v > max {
			max = v
		}
	}

	return
}

// MinMax returns the smallest and biggest element from a list of items.
func MinMax[T constraints.Ordered](vs ...T) (min, max T) {
	if len(vs) == 0 {
		return
	}

	max = vs[0]
	min = vs[0]
	for _, v := range vs[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return
}

// Abs returns the absolute (unsigned) value of a possibly signed value.
func Abs[T constraints.Signed](n T) T {
	if n < 0 {
		return -n
	}
	return n
}
