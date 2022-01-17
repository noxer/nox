// Package dot defines functions and types for the nox library that are so common, that they should be imported as a dot import (`. "github.com/noxer/nox/dot"`) so they may be used without the package name.
package dot

import (
	"constraints"
)

// Optional allows functions to return an optional result. This is useful for situations where the result may not exist but the absence is not an error.
type Optional[T any] struct {
	val T
	set bool
}

// Success creates a new Optional with the result set.
func Success[T any](t T) Optional[T] {
	return Optional[T]{val: t, set: true}
}

// Failure creates a new Options with the result not set.
func Failure[T any]() Optional[T] {
	return Optional[T]{}
}

// HasValue checks if the Optional's value has been set.
func (o Optional[T]) HasValue() bool {
	return o.set
}

// Value returns the value stores in the Optional.
func (o Optional[T]) Value() T {
	return o.val
}

// Result allows functions to return a result or an error.
type Result[T any] struct {
	val T
	err error
}

// OK creates a new Result with the value set to t.
func OK[T any](t T) Result[T] {
	return Result[T]{val: t}
}

// Err creates a new Result with the error set to err.
func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

// Value returns the value of the result without an error check. You should use Unwrap instead.
func (r Result[T]) Value() T {
	return r.val
}

// Success checks if the function call was successful.
func (r Result[T]) Success() bool {
	return r.err == nil
}

// Error returns the error without checking if it exists.
func (r Result[T]) Error() error {
	return r.err
}

type noxPanic struct {
	err error
}

// Unwrap returns the returned value or panics if a error was encountered.
func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(noxPanic{r.err})
	}

	return r.val
}

// IgnoreError catches panics from Result.Unwrap calls and ignores them. Must be deferred.
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

// ReturnError returns the error from Result.Unwrap panics. This only makes
// sense if the function has a named error return value, a pointer to which
// should be passed into this function. Must be deferred.
func ReturnError(into *error) {
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

// HandleError calls f for the error returned from Result.Unwrap. This doesn't
// prevent the termination of the function, it just allows the function to use
// custom code to handle an error. Must be deferred.
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

// Default returns the default value for T. This allows functions to create new
// generic variables. This will be unnecessary in future Go versions.
func Default[T any]() (t T) {
	return
}

// Number defines an interface for all types that can be used with mathematical operators (+-*/%)
type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// NumberSigned defines an interface for all signed types that can be used with mathematical operators (+-*/%)
type NumberSigned interface {
	constraints.Signed | constraints.Float | constraints.Complex
}

// Enumerable defines an interface for enumerable types. See the nox/enumerable package for more functionality.
type Enumerable[T any] interface {
	Next() bool
	Value() T
}
