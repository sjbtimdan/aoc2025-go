package utils

import "strconv"

func ParseIntOrPanic(s string) int64 {
	return OrPanic(strconv.ParseInt(string(s), 10, 64))
}

func ParseUintOrPanic(s string) uint {
	return uint(OrPanic(strconv.ParseUint(string(s), 10, strconv.IntSize)))
}

func AtoiOrPanic(s string) int {
	return OrPanic(strconv.Atoi(string(s)))
}

func OrPanic[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
