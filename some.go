package result_option

type Some[T any] struct {
	v T
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
