package result_option_test

import (
	"errors"
	"fmt"
	"testing"

	result "github.com/axbannaz/resultNoption"
)

func returnErr() result.Result[error] {
	err := errors.New("I got an error")
	var e result.Error[error]
	e.Wrap(err)
	return e
}

func returnError[T any](t *testing.T, v T) result.Result[T] {
	t.Logf("T=%T", v)
	err := fmt.Errorf("I got an error with v=%v", v)
	var e result.Error[T]
	e.Wrap(err)
	return e
}

func TestResultErr(t *testing.T) {
	opt := returnErr()
	switch opt.(type) {
	case result.Error[error]:
	default:
		t.Fatal("not Error")
	}

	t.Logf("v=%v", opt.Err().Unwrap())
}

func TestResultErrorFloat(t *testing.T) {
	opt := returnError(t, 90.4)
	switch v := opt.(type) {
	case result.Error[float64]:
		t.Logf("v=%v", v)
	default:
		t.Fatal("not Error")
	}
}

func TestResultErrorString(t *testing.T) {
	opt := returnError(t, "hell no")
	switch v := opt.(type) {
	case result.Error[string]:
		t.Logf("v=%v", v)
	default:
		t.Fatal("not Error")
	}
}

func TestOptionIsError(t *testing.T) {
	opt := returnErr()
	if !opt.IsError() {
		t.Fatal("not Error")
	}
	t.Logf("v=%v", opt)
	t.Logf("v=%v", opt.Err().Unwrap())
}
