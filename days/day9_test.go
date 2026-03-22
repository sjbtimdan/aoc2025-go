package days

import (
	_ "embed"
	"testing"

	"github.com/shoenig/test/must"
)

//go:embed test_resources/day9.txt
var day9_file []byte

func TestDay9(t *testing.T) {
	result := Day9(day9_file)
	must.Eq(t, "50", result.Part1)
	must.Eq(t, "24", result.Part2)
}

func TestParseCorners(t *testing.T) {
	corners := parseCorners(day9_file)
	must.Eq(t, Corner{x: 7, y: 1}, corners[0])
	must.Eq(t, Corner{x: 7, y: 3}, corners[len(corners)-1])
}

func TestParseEdgesAndBoxes(t *testing.T) {
	corners := []Corner{
		{x: 0, y: 0},
		{x: 0, y: 1},
		{x: 5, y: 1},
		{x: 5, y: 0},
	}

	edges, boxes := parseEdgesAndBoxes(corners)
	must.Eq(t, 4, len(edges))
	must.Eq(t, 6, len(boxes))
	must.Eq(t, int64(12), boxes[0].area)
}

func TestEdgeCrosses(t *testing.T) {
	box := Box{minX: 1, maxX: 4, minY: 1, maxY: 4}

	horizontalCross := Edge{min: 0, max: 5, other: 2, orientation: Horizontal}
	horizontalOnBoundary := Edge{min: 0, max: 5, other: 1, orientation: Horizontal}
	verticalCross := Edge{min: 0, max: 5, other: 2, orientation: Vertical}
	verticalOnBoundary := Edge{min: 0, max: 5, other: 4, orientation: Vertical}

	must.True(t, horizontalCross.crosses(box))
	must.False(t, horizontalOnBoundary.crosses(box))
	must.True(t, verticalCross.crosses(box))
	must.False(t, verticalOnBoundary.crosses(box))
}
