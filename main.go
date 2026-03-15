package main

import (
	_ "embed"
	"fmt"
	"os"
	"sjbtimdan/aoc2025-go/days"
	"sjbtimdan/aoc2025-go/utils"
	"strconv"
	"strings"
	"time"
)

func main() {
	dayFuncs := []utils.DayFunc{
		days.Day1,
	}
	inputs := inputs()
	var selectedDays []int
	if len(os.Args) == 1 {
		selectedDays = rangeArrayFrom1(len(dayFuncs))
	} else {
		selectedDays = parseDayArgs(os.Args[1], len(dayFuncs))
	}
	channels := make([]chan DayFuncResult, len(selectedDays))
	for i, day := range selectedDays {
		channels[i] = make(chan DayFuncResult)
		go runDayFunc(dayFuncs[day-1], day, inputs[day-1], channels[i])
	}
	totalCpu := 0.0
	for _, channel := range channels {
		result := <-channel
		totalCpu += result.elapsed
		fmt.Printf("Day %d [%.2f ms]: ", result.day, result.elapsed)
		fmt.Printf("Part 1: %s; Part 2: %s\n", result.answers.Part1, result.answers.Part2)
	}
	fmt.Printf("Total CPU time: %.2f ms\n", totalCpu)
}

type DayFuncResult struct {
	day     int
	answers utils.Answers
	elapsed float64
}

func runDayFunc(dayFunc utils.DayFunc, day int, input []byte, results chan DayFuncResult) {
	start := time.Now()
	answers := dayFunc(input)
	elapsed := time.Since(start).Seconds() * 1000
	results <- DayFuncResult{day: day, answers: answers, elapsed: elapsed}
}

func rangeArrayFrom1(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	return nums
}

func parseDayArgs(arg string, maxDay int) []int {
	daysList := []int{}
	for _, part := range strings.Split(arg, ",") {
		trimmed := strings.TrimSpace(part)
		day, err := strconv.Atoi(trimmed)
		if err != nil {
			panic(fmt.Sprintf("Invalid integer: %s", trimmed))
		}
		if day < 1 || day > maxDay {
			panic(fmt.Sprintf("Day out of range: %d. Please provide integers between 1 and %d", day, maxDay))
		}
		daysList = append(daysList, day)
	}
	if len(daysList) == 0 {
		panic("No valid day numbers provided.")
	}
	return daysList
}

//go:embed resources/day1.txt
var day1_file []byte

func inputs() [][]byte {
	return [][]byte{
		day1_file,
	}
}
