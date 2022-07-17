package option

type None struct{}
type Some[T any] struct {
	v T
}

func (s Some[T]) Unwrap() (t T) {
	return s.v
}

func (s *Some[T]) Wrap(t T) {
	s.v = t
}

type Option interface{}
