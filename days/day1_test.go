package days

import (
	"os"
	"testing"

	"github.com/shoenig/test/must"
)

func TestParseInstruction(t *testing.T) {
	must.Eq(t, -123, parseInstruction([]byte("L123")))
	must.Eq(t, 456, parseInstruction([]byte("R456")))
}

func TestDay1(t *testing.T) {
	contents, err := os.ReadFile("../test_resources/day1.txt") // Go 1.16+
	if err != nil {
		panic(err) // abort on error
	}
	result := Day1(contents)
	must.Eq(t, "3", result.Part1)
	must.Eq(t, "6", result.Part2)
}
