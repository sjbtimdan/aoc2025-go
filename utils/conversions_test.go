package utils

import (
	"testing"

	"github.com/shoenig/test/must"
)

func TestParseIntOrPanic(t *testing.T) {
	must.Eq(t, int64(42), ParseIntOrPanic("42"))
	must.Eq(t, int64(-7), ParseIntOrPanic("-7"))
}

func TestParseIntOrPanic_Invalid(t *testing.T) {
	must.Panic(t, func() { ParseIntOrPanic("abc") })
}

func TestParseUintOrPanic(t *testing.T) {
	must.Eq(t, uint(42), ParseUintOrPanic("42"))
}

func TestParseUintOrPanic_Invalid(t *testing.T) {
	must.Panic(t, func() { ParseUintOrPanic("abc") })
}

func TestAtoiOrPanic(t *testing.T) {
	must.Eq(t, 42, AtoiOrPanic("42"))
	must.Eq(t, -7, AtoiOrPanic("-7"))
}

func TestAtoiOrPanic_Invalid(t *testing.T) {
	must.Panic(t, func() { AtoiOrPanic("abc") })
}
