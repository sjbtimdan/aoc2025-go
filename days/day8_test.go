package days

import (
	_ "embed"
	"testing"

	"github.com/shoenig/test/must"
)

//go:embed test_resources/day8.txt
var day8_file []byte

func TestDay8(t *testing.T) {
	result := Day8WithCount(day8_file, 10)
	must.Eq(t, "40", result.Part1)
	must.Eq(t, "25272", result.Part2)
}

func TestPositionDistanceTo(t *testing.T) {
	p1 := Position{x: 1, y: 2, z: 3}
	p2 := Position{x: 4, y: 6, z: 8}

	must.Eq(t, int64(50), p1.distanceTo(p2))
}

func TestReadJunctionBoxPositions(t *testing.T) {
	positions := readJunctionBoxPositions(day8_file)

	must.Len(t, 20, positions)
	must.Eq(t, Position{x: 62, y: 817, z: 812}, positions[0])
	must.Eq(t, Position{x: 425, y: 690, z: 689}, positions[19])
}
