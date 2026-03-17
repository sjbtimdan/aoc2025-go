package days

import (
	_ "embed"
	"testing"

	"github.com/shoenig/test/must"
)

//go:embed test_resources/day12.txt
var day12_file []byte

func TestDay12(t *testing.T) {
	result := Day12(day8_file)
	must.Eq(t, "TODO", result.Part1)
	must.Eq(t, "TODO", result.Part2)
}
