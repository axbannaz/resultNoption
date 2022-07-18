package result_option

import "fmt"

type Error[T any] struct {
	err error
}

func (o Error[T]) unwrap() error {
	return o.err
}

func (o Error[T]) Unwrap() T {
	panic(o.unwrap())
}

func (o Error[T]) Expect(msg string) T {
	m := fmt.Sprintf("%s: %s", o.unwrap(), msg)
	panic(m)
}

func (o *Error[T]) Wrap(t error) {
	o.err = t
}

func (_ Error[T]) IsError() (ok bool) {
	ok = true
	return
}

func (_ Error[T]) IsOk() (ok bool) {
	return
}

func (_ Error[T]) Ok() Option[T] {
	var n None[T]
	return n
}

func (n Error[T]) Err() Option[error] {
	var o Some[error]
	o.Wrap(n.unwrap())

	return o
}

func (_ Error[T]) UnwrapOr(defaultT T) T {
	return defaultT
}

func (n Error[T]) UnwrapOrElse(f func(arg error) T) T {
	return f(n.unwrap())
}
