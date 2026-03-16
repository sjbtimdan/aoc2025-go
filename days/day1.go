package days

import (
	"bytes"
	"sjbtimdan/aoc2025-go/utils"
)

func Day1(contents []byte) utils.Answers {
	dial := 50
	part1Count := 0
	part2Count := 0
	for _, line := range bytes.Split(contents, []byte("\n")) {
		rotation := parseInstruction(line)
		if dial != 0 {
			var partialRotation = dial + rotation%100
			if partialRotation <= 0 || partialRotation >= 100 {
				part2Count += 1
			}
		} else {
			part1Count += 1
		}
		dial = (RemEuclid(dial+rotation, 100))
		part2Count += AbsInt(rotation) / 100
	}
	return utils.IntAnswers(part1Count, part2Count)
}

func parseInstruction(instruction []byte) int {
	magnitudeStr := string(instruction[1:])
	magnitude := utils.AtoiOrPanic(magnitudeStr)
	if instruction[0] == 'L' {
		return -magnitude
	} else {
		return magnitude
	}
}

func RemEuclid(a, b int) int {
	return (a%b + b) % b
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
