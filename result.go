package result_option

type Result[T any] interface {
	IsError() bool
	IsOk() bool
	Error() Error[T]
	Ok() Ok[T]
}

type Error[T any] struct {
	Err error
}

type Ok[T any] struct {
	v T
}

func (o Ok[T]) Unwrap() (t T) {
	return o.v
}

func (o *Ok[T]) Wrap(t T) {
	o.v = t
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
