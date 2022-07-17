package result_option_test

import (
	"testing"

	option "github.com/axbannaz/resultNoption"
)

func returnNone() option.Option[any] {
	return option.None[any]{}
}

func returnSome[T any](v T) option.Option[T] {
	var some option.Some[T]
	some.Wrap(v)
	return some
}

func TestOptionNone(t *testing.T) {
	opt := returnNone()
	switch v := opt.(type) {
	case option.None[any]:
		t.Logf("v=%v", v)
	case option.Some[any]:
		t.Fatal("not None")
	default:
		t.Fatal("not None")
	}
}

func TestOptionIsNone(t *testing.T) {
	opt := returnNone()
	if !opt.IsNone() {
		t.Fatal("not None")
	}
	t.Logf("v=%v", opt)
}

func TestOptionSome(t *testing.T) {
	someValue42 := 42
	opt := returnSome(someValue42)
	switch v := opt.(type) {
	case option.None[int]:
		t.Fatal("None")
	case option.Some[int]:
		val := v.Unwrap()
		t.Logf("val=%v", val)
		if val != someValue42 {
			t.Fatalf("%T test failed", val)
		}
	default:
		t.Fatalf("not Some int: %T", v)
	}

	someValuePi := 3.14
	opt1 := returnSome(someValuePi)
	switch v := opt1.(type) {
	case option.None[float64]:
		t.Fatal("None")
	case option.Some[float64]:
		val := v.Unwrap()
		t.Logf("val=%v", val)
		if val != someValuePi {
			t.Fatalf("%T test failed", val)
		}
	default:
		t.Fatalf("not Some float64: %T", v)
	}
}

func TestOptionIsSome(t *testing.T) {
	someValuePi := 3.14
	opt := returnSome(someValuePi)
	if !opt.IsSome() {
		t.Fatalf("not Some %T", opt)
	}
	t.Logf("v=%v", opt)
}
