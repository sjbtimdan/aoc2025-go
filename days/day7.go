package days

import (
	"bytes"
	"sjbtimdan/aoc2025-go/utils"
)

func Day7(contents []byte) utils.Answers {
	manifold := parseManifold(contents)
	beams := make([][]int, len(manifold.grid))
	for y := range beams {
		beams[y] = make([]int, len(manifold.grid[y]))
	}
	for x := 0; x < manifold.width; x++ {
		if manifold.grid[0][x] == 'S' {
			beams[0][x] = 1
			break
		}
	}
	for y := 1; y < len(manifold.grid); y++ {
		for x := range manifold.grid[y] {
			manifold.markBeams(beams, x, y)
		}
	}
	part1, part2 := 0, 1
	for row := 0; row < len(beams)-1; row++ {
		for col := 0; col < manifold.width; col++ {
			if beams[row][col] > 0 && manifold.grid[row+1][col] == '^' {
				part1++
				part2 += beams[row][col]
			}
		}
	}
	return utils.IntAnswers(part1, part2)
}

type Manifold struct {
	width      int
	startIndex int
	grid       [][]byte
}

func (m Manifold) markBeams(beams [][]int, x, y int) {
	if m.grid[y][x] != '.' {
		return
	}
	aboveBeamsRow := beams[y-1]
	if aboveBeamsRow[x] > 0 {
		beams[y][x] += aboveBeamsRow[x]
	}
	if x < m.width-1 {
		right := m.grid[y][x+1]
		if right == '^' && aboveBeamsRow[x+1] > 0 {
			beams[y][x] += aboveBeamsRow[x+1]
		}
	}
	if x > 0 {
		left := m.grid[y][x-1]
		if left == '^' && aboveBeamsRow[x-1] > 0 {
			beams[y][x] += aboveBeamsRow[x-1]
		}
	}
}

func (m Manifold) countBeams(beams [][]int) int {
	count := 0
	for row := 0; row < len(beams)-1; row++ {
		for col := 0; col < m.width; col++ {
			if beams[row][col] > 0 && m.grid[row+1][col] == '^' {
				count++
			}
		}
	}
	return count
}

func parseManifold(contents []byte) Manifold {
	grid := bytes.Split(contents, []byte{'\n'})
	width := len(grid[0])
	startIndex := bytes.Index(grid[0], []byte{'S'})
	return Manifold{
		width:      width,
		startIndex: startIndex,
		grid:       grid,
	}
}
