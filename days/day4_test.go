package days

import (
	_ "embed"
	"testing"

	"github.com/shoenig/test/must"
)

//go:embed test_resources/day4.txt
var day4_file []byte

func TestReadGrid(t *testing.T) {
	grid := readGrid(day4_file)
	must.Eq(t, 10, len(grid))
	must.Eq(t, 10, len(grid[0]))
	must.Eq(t, '.', grid[0][0])
	must.Eq(t, '@', grid[1][0])
}

func TestCountSurroundingRolls(t *testing.T) {
	grid := readGrid(day4_file)
	tests := []struct {
		x, y uint
		want uint
	}{
		{0, 0, 2},
		{2, 0, 3},
		{2, 2, 6},
		{9, 9, 2},
	}
	for _, test := range tests {
		must.Eq(t, test.want, countSurroundingRolls(grid, test.x, test.y))
	}
}

func TestDat4(t *testing.T) {
	result := Day4(day4_file)
	must.Eq(t, "13", result.Part1)
	must.Eq(t, "43", result.Part2)
}
