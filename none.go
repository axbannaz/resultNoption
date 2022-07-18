package result_option

import "errors"

type None[T any] struct{}

func (_ None[T]) IsNone() (ok bool) {
	ok = true
	return
}

func (_ None[T]) IsSome() (ok bool) {
	return
}

func (_ None[T]) Unwrap() T {
	panic("invalid type")
}

func (_ None[T]) Expect(err string) T {
	panic(err)
}

func (_ None[T]) OkOr(msg string) (r Result[T]) {
	var err Error[T]
	r = err
	err.Wrap(errors.New(msg))

	return
}

func (n None[T]) OkOrElse(fn func() string) (r Result[T]) {
	r = n.OkOr(fn())

	return
}

func (_ None[T]) UnwrapOr(defaultT T) T {
	return defaultT
}

func (_ None[T]) UnwrapOrDefault() T {
	var n T
	return n
}

func (_ None[T]) UnwrapOrElse(f func() T) T {
	return f()
}
