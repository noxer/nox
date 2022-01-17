package dot

import (
	"constraints"
)

type Optional[T any] struct {
	val T
	set bool
}

func Success[T any](t T) Optional[T] {
	return Optional[T]{val: t, set: true}
}

func Failure[T any]() Optional[T] {
	return Optional[T]{}
}

func (o Optional[T]) HasValue() bool {
	return o.set
}

func (o Optional[T]) Value() T {
	return o.val
}

type Result[T any] struct {
	val T
	err error
}

func OK[T any](t T) Result[T] {
	return Result[T]{val: t}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (r Result[T]) Value() T {
	return r.val
}

func (r Result[T]) Success() bool {
	return r.err == nil
}

func (r Result[T]) Error() error {
	return r.err
}

type noxPanic struct {
	err error
}

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(noxPanic{r.err})
	}

	return r.val
}

func IgnoreError() {
	e := recover()
	if e == nil {
		return
	}

	if _, ok := e.(noxPanic); ok {
		return
	}

	panic(e)
}

func CatchError(into *error) {
	e := recover()
	if e == nil {
		return
	}

	if ne, ok := e.(noxPanic); ok {
		*into = ne.err
		return
	}

	panic(e)
}

func HandleError(f func(error)) {
	e := recover()
	if e == nil {
		return
	}

	if ne, ok := e.(noxPanic); ok {
		f(ne.err)
		return
	}

	panic(e)
}

func Default[T any]() (t T) {
	return
}

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

type NumberSigned interface {
	constraints.Signed | constraints.Float | constraints.Complex
}

type Enumerable[T any] interface {
	Next() bool
	Value() T
}
