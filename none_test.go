package result_option_test

import (
	"testing"

	option "github.com/axbannaz/resultNoption"
)

func returnNone() option.Option[any] {
	return option.None[any]{}
}

func TestOptionNone(t *testing.T) {
	opt := returnNone()
	switch v := opt.(type) {
	case option.None[any]:
		t.Logf("v=%v", v)
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
