package days

import (
	"bytes"
	"sjbtimdan/aoc2025-go/utils"
)

func Day4(contents []byte) utils.Answers {
	grid := readGrid(contents)
	part1 := uint(0)
	part2 := uint(0)
	forEachRemovableRoll(grid, func(x, y uint) {
		part1++
	})
	for {
		var removeableRolls []struct{ x, y uint }
		forEachRemovableRoll(grid, func(x, y uint) {
			removeableRolls = append(removeableRolls, struct{ x, y uint }{x, y})
		})
		if len(removeableRolls) == 0 {
			break
		}
		for _, roll := range removeableRolls {
			grid[roll.y][roll.x] = '.'
			part2++
		}
	}
	return utils.Uint64Answers(uint64(part1), uint64(part2))
}

func forEachRemovableRoll(grid [][]rune, f func(x, y uint)) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '@' && countSurroundingRolls(grid, uint(x), uint(y)) < 4 {
				f(uint(x), uint(y))
			}
		}
	}
}

func countSurroundingRolls(grid [][]rune, x, y uint) uint {
	xLen := len(grid[0])
	yLen := len(grid)
	containsRoll := func(dx, dy int) uint {
		x1 := int(x) + dx
		y1 := int(y) + dy
		if x1 < 0 || y1 < 0 || x1 >= xLen || y1 >= yLen {
			return 0
		}
		return utils.BoolToUint(grid[y1][x1] == '@')
	}
	return containsRoll(-1, -1) + containsRoll(0, -1) + containsRoll(1, -1) +
		containsRoll(-1, 0) + containsRoll(1, 0) +
		containsRoll(-1, 1) + containsRoll(0, 1) + containsRoll(1, 1)
}

func readGrid(contents []byte) [][]rune {
	var grid [][]rune
	for _, line := range bytes.Split(contents, []byte("\n")) {
		row := []rune{}
		for _, ch := range line {
			row = append(row, rune(ch))
		}
		grid = append(grid, row)
	}
	return grid
}
