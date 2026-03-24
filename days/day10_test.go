package days

import (
	_ "embed"
	"testing"

	"github.com/shoenig/test/must"
)

//go:embed test_resources/day10.txt
var day10_file []byte

func TestDay10(t *testing.T) {
	result := Day10(day10_file)
	must.Eq(t, "TODO", result.Part1)
	must.Eq(t, "TODO", result.Part2)
}

func TestParseMachines(t *testing.T) {
	machines := parseMachines([]byte("[.##.] (3) (1,3) (2) {3,5,4,7}"))
	must.Len(t, 1, machines)
	expected := Machine{
		desiredState: 0b0110,
		buttons:      []uint16{0b0001, 0b0101, 0b0010},
	}
	must.Eq(t, expected, machines[0])
}
