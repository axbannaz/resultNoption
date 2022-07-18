package result_option

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
