package utils

import (
	"testing"

	"github.com/shoenig/test/must"
)

func TestStrconvOrPanic(t *testing.T) {
	must.Eq(t, int64(42), StrconvOrPanic("42"))
	must.Eq(t, int64(-7), StrconvOrPanic("-7"))
}

func TestStrconvOrPanic_Invalid(t *testing.T) {
	must.Panic(t, func() { StrconvOrPanic("abc") })
}

func TestAtoiOrPanic(t *testing.T) {
	must.Eq(t, 42, AtoiOrPanic("42"))
	must.Eq(t, -7, AtoiOrPanic("-7"))
}

func TestAtoiOrPanic_Invalid(t *testing.T) {
	must.Panic(t, func() { AtoiOrPanic("abc") })
}
