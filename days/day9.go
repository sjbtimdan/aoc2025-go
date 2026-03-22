package days

import (
	"bytes"
	"sjbtimdan/aoc2025-go/utils"
	"slices"
)

func Day9(contents []byte) utils.Answers {
	corners := parseCorners(contents)
	edges, boxes := parseEdgesAndBoxes(corners)
	part1 := boxes[0].area
	var part2 int64
	for _, box := range boxes {
		var crosses bool
		for edge := range edges {
			if edges[edge].crosses(box) {
				crosses = true
				break
			}
		}
		if !crosses {
			part2 = box.area
			break
		}
	}
	return utils.Int64Answers(part1, part2)
}

func parseEdgesAndBoxes(corners []Corner) ([]Edge, []Box) {
	numCorners := len(corners)
	edges := make([]Edge, 0, numCorners)
	boxes := []Box{}
	for i := range corners {
		c1 := corners[i]
		c2 := corners[(i+1)%numCorners]
		var edge Edge
		if c1.x == c2.x {
			edge = Edge{
				min:         utils.MinInt64(c1.y, c2.y),
				max:         utils.MaxInt64(c1.y, c2.y),
				other:       c1.x,
				orientation: Vertical,
			}
		} else {
			edge = Edge{
				min:         utils.MinInt64(c1.x, c2.x),
				max:         utils.MaxInt64(c1.x, c2.x),
				other:       c1.y,
				orientation: Horizontal,
			}
		}
		edges = append(edges, edge)
		for j := i + 1; j < numCorners; j++ {
			c1Join := corners[j]
			area := c1.areaTo(c1Join)
			box := Box{
				area: area,
				minX: utils.MinInt64(c1.x, c1Join.x),
				maxX: utils.MaxInt64(c1.x, c1Join.x),
				minY: utils.MinInt64(c1.y, c1Join.y),
				maxY: utils.MaxInt64(c1.y, c1Join.y),
			}
			boxes = append(boxes, box)
		}
	}
	slices.SortFunc(boxes, func(b1, b2 Box) int {
		return -int(b1.area - b2.area)
	})
	return edges, boxes
}

type Corner struct {
	x, y int64
}

func (c Corner) areaTo(other Corner) int64 {
	return (1 + utils.AbsInt64(c.x-other.x)) * (1 + utils.AbsInt64(c.y-other.y))
}

type Orientation int

const (
	Horizontal Orientation = iota
	Vertical
)

type Edge struct {
	min, max    int64
	other       int64
	orientation Orientation
}

type Box struct {
	area, minX, maxX, minY, maxY int64
}

func (e Edge) crosses(box Box) bool {
	if e.orientation == Horizontal {
		return e.max > box.minX && e.min < box.maxX && (box.minY < e.other && e.other < box.maxY)
	} else {
		return e.max > box.minY && e.min < box.maxY && (box.minX < e.other && e.other < box.maxX)
	}
}

func parseCorners(contents []byte) []Corner {
	var corners []Corner
	lines := bytes.Split(contents, []byte("\n"))
	for _, line := range lines {
		parts := bytes.Split(line, []byte(","))
		corner := Corner{
			x: utils.ParseInt64OrPanic(string(parts[0])),
			y: utils.ParseInt64OrPanic(string(parts[1])),
		}
		corners = append(corners, corner)
	}
	return corners
}
