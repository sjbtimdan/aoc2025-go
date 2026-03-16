package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Answers struct {
	Part1 string
	Part2 string
}

type DayFunc func([]byte) Answers

type DayPuzzle struct {
	DayFunc DayFunc
	Input   []byte
}

func IntAnswers(part1, part2 int) Answers {
	return Answers{
		Part1: strconv.Itoa(part1),
		Part2: strconv.Itoa(part2),
	}
}

func StringAnswers(part1, part2 string) Answers {
	return Answers{
		Part1: part1,
		Part2: part2,
	}
}

func Part1OnlyIntAnswers(part1 int) Answers {
	return Answers{
		Part1: strconv.Itoa(part1),
		Part2: "TODO",
	}
}

func TodoAnswers() Answers {
	return Answers{
		Part1: "TODO",
		Part2: "TODO",
	}
}

func Run(dayPuzzles []DayPuzzle) {
	var selectedDays []int
	if len(os.Args) == 1 {
		selectedDays = rangeArrayFrom1(len(dayPuzzles))
	} else {
		selectedDays = parseDayArgs(os.Args[1], len(dayPuzzles))
	}
	channels := make([]chan DayFuncResult, len(selectedDays))
	for i, day := range selectedDays {
		channels[i] = make(chan DayFuncResult)
		dayPuzzle := dayPuzzles[day-1]
		go runDayFunc(dayPuzzle.DayFunc, day, dayPuzzle.Input, channels[i])
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
	answers Answers
	elapsed float64
}

func runDayFunc(dayFunc DayFunc, day int, input []byte, results chan DayFuncResult) {
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
