package result_option

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
