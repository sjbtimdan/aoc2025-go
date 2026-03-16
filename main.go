package main

import (
	_ "embed"
	"sjbtimdan/aoc2025-go/days"
	"sjbtimdan/aoc2025-go/utils"
)

//go:embed resources/day1.txt
var day1_file []byte

func main() {
	utils.Run([]utils.DayPuzzle{
		{DayFunc: days.Day1, Input: day1_file},
	})
}
