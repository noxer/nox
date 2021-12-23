package nox

import "constraints"

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

func Abs[T constraints.Signed](n T) T {
	if n < 0 {
		return -n
	}
	return n
}
