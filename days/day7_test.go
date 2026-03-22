package days

import (
	_ "embed"
	"testing"

	"github.com/shoenig/test/must"
)

//go:embed test_resources/day7.txt
var day7_file []byte

func TestParseManifold(t *testing.T) {
	manifold := parseManifold(day7_file)
	must.Eq(t, 7, manifold.startIndex)
	must.Eq(t, 15, manifold.width)
	must.Eq(t, byte('S'), manifold.grid[0][7])
}

func TestDay7(t *testing.T) {
	result := Day7(day7_file)
	must.Eq(t, "21", result.Part1)
	must.Eq(t, "40", result.Part2)
}
