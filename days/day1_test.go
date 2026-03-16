package days

import (
	_ "embed"
	"testing"

	"github.com/shoenig/test/must"
)

func TestParseInstruction(t *testing.T) {
	must.Eq(t, -123, parseInstruction([]byte("L123")))
	must.Eq(t, 456, parseInstruction([]byte("R456")))
}

//go:embed test_resources/day1.txt
var day1_file []byte

func TestDay1(t *testing.T) {
	result := Day1(day1_file)
	must.Eq(t, "3", result.Part1)
	must.Eq(t, "6", result.Part2)
}
