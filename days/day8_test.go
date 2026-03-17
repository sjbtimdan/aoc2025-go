package days

import (
	_ "embed"
	"testing"

	"github.com/shoenig/test/must"
)

//go:embed test_resources/day8.txt
var day8_file []byte

func TestDay8(t *testing.T) {
	result := Day8(day8_file)
	must.Eq(t, "TODO", result.Part1)
	must.Eq(t, "TODO", result.Part2)
}
