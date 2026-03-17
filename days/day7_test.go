package days

import (
	_ "embed"
	"testing"

	"github.com/shoenig/test/must"
)

//go:embed test_resources/day7.txt
var day7_file []byte

func TestDay7(t *testing.T) {
	result := Day7(day7_file)
	must.Eq(t, "TODO", result.Part1)
	must.Eq(t, "TODO", result.Part2)
}
