package main

import (
	_ "embed"
	"sjbtimdan/aoc2025-go/days"
	"sjbtimdan/aoc2025-go/utils"
)

//go:embed resources/day1.txt
var day1_file []byte

//go:embed resources/day2.txt
var day2_file []byte

//go:embed resources/day3.txt
var day3_file []byte

func main() {
	utils.Run([]utils.DayPuzzle{
		{DayFunc: days.Day1, Input: day1_file},
		{DayFunc: days.Day2, Input: day2_file},
		{DayFunc: days.Day3, Input: day3_file},
	})
}
