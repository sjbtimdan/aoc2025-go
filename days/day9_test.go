package days

import (
	_ "embed"
	"testing"

	"github.com/shoenig/test/must"
)

//go:embed test_resources/day9.txt
var day9_file []byte

func TestDay9(t *testing.T) {
	result := Day9(day8_file)
	must.Eq(t, "TODO", result.Part1)
	must.Eq(t, "TODO", result.Part2)
}
