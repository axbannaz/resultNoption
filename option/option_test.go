package option_test

import (
	"testing"

	"github.com/axbannaz/resultNoption/option"
)

func returnNone() option.Option {
	return option.None{}
}

func returnSome[T any](v T) option.Option {
	var some option.Some[T]
	some.Wrap(v)
	return some
}

func TestOptionNone(t *testing.T) {
	opt := returnNone()
	switch v := opt.(type) {
	case option.None:
		t.Logf("v=%v", v)
	case option.Some[int]:
		t.Fatal("not None")
	default:
		t.Fatal("not None")
	}
}

func TestOptionSome(t *testing.T) {
	someValue42 := 42
	opt := returnSome(someValue42)
	switch v := opt.(type) {
	case option.None:
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
	opt = returnSome(someValuePi)
	switch v := opt.(type) {
	case option.None:
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
