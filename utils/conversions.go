package utils

import "strconv"

func StrconvOrPanic(s string) int64 {
	value, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		panic(err)
	}
	return value
}

func AtoiOrPanic(s string) int {
	value, err := strconv.Atoi(string(s))
	if err != nil {
		panic(err)
	}
	return value
}
