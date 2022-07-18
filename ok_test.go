package result_option_test

import (
	"testing"

	result "github.com/axbannaz/resultNoption"
)

func returnOk[T any](v T) result.Result[T] {
	var ok result.Ok[T]
	ok.Wrap(v)
	return ok
}

func TestResultOk(t *testing.T) {
	someValue42 := 42
	opt := returnOk(someValue42)
	switch v := opt.(type) {
	case result.Ok[int]:
		val := v.Unwrap()
		t.Logf("val=%v", val)
		if val != someValue42 {
			t.Fatalf("%T test failed", val)
		}
	default:
		t.Fatalf("not Some int: %T", v)
	}
}

func TestResultIsOk(t *testing.T) {
	someValuePi := 3.14
	opt := returnOk(someValuePi)
	if !opt.IsOk() {
		t.Fatalf("not Some %T", opt)
	}
	t.Logf("v=%v", opt)
	t.Logf("v=%v", opt.Ok().Unwrap())
}
