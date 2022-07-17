package result_option

import "errors"

type Option[T any] interface {
	IsNone() bool
	IsSome() bool
	Expect(err string) T
	Unwrap() T
	UnwrapOr(def T) T
	UnwrapOrDefault() T
	UnwrapOrElse(fn func() T) T
	OkOr(msg string) Result[T]
	OkOrElse(fn func() string) Result[T]
}

type None[T any] struct{}

type Some[T any] struct {
	v T
}

func (n None[T]) IsNone() (ok bool) {
	ok = true
	return
}

func (n None[T]) IsSome() (ok bool) {
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

func (s Some[T]) Expect(_ string) T {
	return s.Unwrap()
}

func (s Some[T]) OkOr(_ string) (r Result[T]) {
	var ok Ok[T]

	ok.Wrap(s.Unwrap())

	return ok
}

func (s Some[T]) OkOrElse(_ func() string) (r Result[T]) {
	r = s.OkOr("")

	return
}

func (s Some[T]) Unwrap() T {
	return s.v
}

func (s Some[T]) UnwrapOr(_ T) T {
	return s.Unwrap()
}

func (s Some[T]) UnwrapOrDefault() T {
	return s.Unwrap()
}

func (s Some[T]) UnwrapOrElse(_ func() T) T {
	return s.Unwrap()
}

func (s *Some[T]) Wrap(t T) {
	s.v = t
}

func (_ Some[T]) IsNone() (ok bool) {
	return
}

func (_ Some[T]) IsSome() (ok bool) {
	ok = true
	return
}
