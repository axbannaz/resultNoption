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

func (n None[T]) None() None[T] {
	return n
}

func (n None[T]) Some() Some[T] {
	panic("invalid type")
}

func (s Some[T]) Unwrap() (t T) {
	return s.v
}

func (s *Some[T]) Wrap(t T) {
	s.v = t
}

func (s Some[T]) IsNone() (ok bool) {
	return
}

func (s Some[T]) IsSome() (ok bool) {
	ok = true
	return
}

func (s Some[T]) None() None[T] {
	panic("invalid type")
}

func (s Some[T]) Some() Some[T] {
	return s
}

type Option[T any] interface {
	IsNone() bool
	IsSome() bool
	Some() Some[T]
	None() None[T]
}
