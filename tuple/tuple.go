package tuple

type T2[T, U any] struct {
	A T
	B U
}

func ToList2[T any](t T2[T, T]) []T {
	return []T{t.A, t.B}
}

type T3[T, U, V any] struct {
	A T
	B U
	C V
}

func ToList3[T any](t T3[T, T, T]) []T {
	return []T{t.A, t.B, t.C}
}

type T4[T, U, V, W any] struct {
	A T
	B U
	C V
	D W
}

func ToList4[T any](t T4[T, T, T, T]) []T {
	return []T{t.A, t.B, t.C, t.D}
}
