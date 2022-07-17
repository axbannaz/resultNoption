package result

type Result interface{}

type Err error

type Ok[T any] struct {
	v T
}

func (o Ok[T]) Unwrap() (t T) {
	return o.v
}

func (o *Ok[T]) Wrap(t T) {
	o.v = t
}
