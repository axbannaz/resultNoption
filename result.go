package result_option

import "fmt"

type Result[T any] interface {
	IsError() bool
	IsOk() bool
	Error() Error[T]
	Ok() Ok[T]
	Expect(msg string) T
	Unwrap() T
	UnwrapOr(def T) T
}

type Error[T any] struct {
	Err error
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
	return o.Err
}

func (o Error[T]) Unwrap() T {
	panic(o.unwrap())
}

func (o Error[T]) Expect(msg string) T {
	m := fmt.Sprintf("%s: %s", o.unwrap(), msg)
	panic(m)
}

func (o *Error[T]) Wrap(t error) {
	o.Err = t
}

func (_ Error[T]) IsError() (ok bool) {
	ok = true
	return
}

func (_ Error[T]) IsOk() (ok bool) {
	return
}

func (n Error[T]) Error() Error[T] {
	return n
}

func (_ Error[T]) Ok() Ok[T] {
	panic("invalid type")
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

func (n Ok[T]) Ok() Ok[T] {
	return n
}

func (n Ok[T]) UnwrapOr(_ T) T {
	return n.Unwrap()
}

func (n Ok[T]) UnwrapOrElse(_ func() T) T {
	return n.Unwrap()
}
