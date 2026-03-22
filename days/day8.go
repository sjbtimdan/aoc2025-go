package days

import (
	"bytes"
	"math"
	"sjbtimdan/aoc2025-go/utils"
	"sort"
)

func Day8(contents []byte) utils.Answers {
	return Day8WithCount(contents, 1000)
}

func Day8WithCount(contents []byte, limit int) utils.Answers {
	positions := readJunctionBoxPositions(contents)
	partitions_count := len(positions)
	dsu := utils.NewUnionFind(partitions_count)
	distances := calculateDistances(positions)
	var sizes []int
	for i, d := range distances {
		if dsu.Union(d.i, d.j) {
			partitions_count--
			if partitions_count == 2 {
				break
			}
		}
		if i == limit-1 {
			partitions := dsu.PartitionInfos()
			sizes = make([]int, 0, len(partitions))
			for _, partition := range partitions {
				sizes = append(sizes, partition.Size)
			}
		}
	}
	var part2 int64
	ps := dsu.PartitionInfos()
	for _, partitionInfo := range ps {
		if partitionInfo.Size == 1 {
			lastPosition := positions[partitionInfo.Representative]
			var secondLastPositionIndex int
			minDist := int64(math.MaxInt64)
			for _, d := range distances {
				if d.i == partitionInfo.Representative || d.j == partitionInfo.Representative {
					if d.d < minDist {
						minDist = d.d
						if d.i == partitionInfo.Representative {
							secondLastPositionIndex = d.j
						} else {
							secondLastPositionIndex = d.i
						}
					}
				}
			}
			secondLastPosition := positions[secondLastPositionIndex]
			part2 = lastPosition.x * secondLastPosition.x
			break
		}
	}
	top3 := largestThree(sizes)
	part1 := top3[0] * top3[1] * top3[2]
	return utils.Uint64Answers(uint64(part1), uint64(part2))
}

func largestThree(nums []int) [3]int {
	first, second, third := math.MinInt, math.MinInt, math.MinInt
	for _, num := range nums {
		if num > first {
			third = second
			second = first
			first = num
		} else if num > second {
			third = second
			second = num
		} else if num > third {
			third = num
		}
	}

	return [3]int{first, second, third}
}

type Position struct {
	x, y, z int64
}

type Distance struct {
	i, j int
	d    int64
}

func calculateDistances(positions []Position) []Distance {
	distances := make([]Distance, 0, len(positions)*(len(positions)-1)/2)
	for i := range positions {
		for j := i + 1; j < len(positions); j++ {
			distances = append(distances, Distance{i: i, j: j, d: positions[i].distanceTo(positions[j])})
		}
	}
	sort.Slice(
		distances,
		func(i, j int) bool {
			return distances[i].d < distances[j].d
		},
	)
	return distances
}

func (p Position) distanceTo(other Position) int64 {
	dx := p.x - other.x
	dy := p.y - other.y
	dz := p.z - other.z
	return utils.AbsInt64(dx*dx) + utils.AbsInt64(dy*dy) + utils.AbsInt64(dz*dz)
}

func readJunctionBoxPositions(contents []byte) []Position {
	lines := bytes.Split(contents, []byte{'\n'})
	positions := make([]Position, len(lines))
	for i, line := range lines {
		coords := bytes.Split(line, []byte{','})
		positions[i] = Position{
			x: utils.ParseInt64OrPanic(string(coords[0])),
			y: utils.ParseInt64OrPanic(string(coords[1])),
			z: utils.ParseInt64OrPanic(string(coords[2])),
		}
	}
	return positions
}
