package result_option

import "fmt"

type Result[T any] interface {
	IsError() bool
	IsOk() bool
	Expect(msg string) T
	Unwrap() T
	UnwrapOr(def T) T
	Err() Option[error]
	Ok() Option[T]
}

type Error[T any] struct {
	err error
}

type Ok[T any] struct {
	v T
}

func (o Ok[T]) Unwrap() T {
	return o.v
}

func (o Ok[T]) Expect(_ string) T {
	return o.Unwrap()
}

func (o *Ok[T]) Wrap(t T) {
	o.v = t
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

func (_ Ok[T]) IsError() (ok bool) {
	return
}

func (_ Ok[T]) IsOk() (ok bool) {
	ok = true
	return
}

func (_ Ok[T]) Error() Error[T] {
	panic("invalid type")
}

func (n Ok[T]) Ok() (o Option[T]) {
	var s Some[T]
	s.Wrap(n.Unwrap())
	return s
}

func (_ Ok[T]) Err() Option[error] {
	var n None[error]
	return n
}

func (n Ok[T]) UnwrapOr(_ T) T {
	return n.Unwrap()
}

func (n Ok[T]) UnwrapOrElse(_ func() T) T {
	return n.Unwrap()
}
