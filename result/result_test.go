package result_test

import (
	"errors"
	"testing"

	"github.com/axbannaz/resultNoption/result"
)

func returnErr() result.Result {
	return result.Err(
		errors.New("error"),
	)
}

func returnOk[T any](v T) result.Result {
	var ok result.Ok[T]
	ok.Wrap(v)
	return ok
}

func TestResultErr(t *testing.T) {
	opt := returnErr()
	switch v := opt.(type) {
	case result.Err:
		t.Logf("v=%v", v)
	default:
		t.Fatal("not None")
	}
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
