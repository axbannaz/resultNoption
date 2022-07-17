package option

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

type Option[T any] interface {
	IsNone() bool
	IsSome() bool
	Unwrap() T
	UnwrapOr(def T) T
	UnwrapOrDefault() T
	UnwrapOrElse(fn func() T) T
}
