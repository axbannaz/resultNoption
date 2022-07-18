package result_option

type Result[T any] interface {
	IsError() bool
	IsOk() bool
	Expect(msg string) T
	Unwrap() T
	UnwrapOr(def T) T
	Err() Option[error]
	Ok() Option[T]
}
