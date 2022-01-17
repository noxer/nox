package tuple

// T2 is a tuple with two elements
type T2[T, U any] struct {
	A T
	B U
}

// Unpack this tuple into individual values.
func (t T2[T, U]) Unpack() (T, U) {
	return t.A, t.B
}

// ToList2 converts a T2 into a slice of its elements. This can only be used for tuples with matching element types.
func ToList2[T any](t T2[T, T]) []T {
	return []T{t.A, t.B}
}

// T3 is a tuple with three elements
type T3[T, U, V any] struct {
	A T
	B U
	C V
}

// Unpack this tuple into individual values.
func (t T3[T, U, V]) Unpack() (T, U, V) {
	return t.A, t.B, t.C
}

// ToList3 converts a T3 into a slice of its elements. This can only be used for tuples with matching element types.
func ToList3[T any](t T3[T, T, T]) []T {
	return []T{t.A, t.B, t.C}
}

// T3 is a tuple with four elements
type T4[T, U, V, W any] struct {
	A T
	B U
	C V
	D W
}

// Unpack this tuple into individual values.
func (t T4[T, U, V, W]) Unpack() (T, U, V, W) {
	return t.A, t.B, t.C, t.D
}

// ToList4 converts a T4 into a slice of its elements. This can only be used for tuples with matching element types.
func ToList4[T any](t T4[T, T, T, T]) []T {
	return []T{t.A, t.B, t.C, t.D}
}
