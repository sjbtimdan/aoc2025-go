package days

import (
	"sjbtimdan/aoc2025-go/utils"
	"strings"
)

func Day6(contents []byte) utils.Answers {
	problems := parseProblems(contents)
	part1, part2 := uint64(0), uint64(0)
	for i, numbers := range problems.numbers {
		operation := problems.operations[i]
		part1 += calculateColumnResult(numbers, operation)
		part2 += calculateRightToLeftColumnResult(numbers, operation)
	}
	return utils.Uint64Answers(part1, part2)
}

type Problems struct {
	numbers    [][]string
	operations []byte
}

func calculateColumnResult(numbers []string, operation byte) uint64 {
	var result uint64
	switch operation {
	case '+':
		result = 0
	case '*':
		result = 1
	}
	for _, numberStr := range numbers {
		number := utils.ParseUint64OrPanic(strings.TrimSpace(numberStr))
		switch operation {
		case '+':
			result += number
		case '*':
			result *= number
		}
	}
	return result
}

func calculateRightToLeftColumnResult(numbers []string, operation byte) uint64 {
	var operationFn func(a, b uint64) uint64
	var defaultValue uint64
	if operation == '+' {
		operationFn = func(a, b uint64) uint64 { return a + b }
	} else {
		operationFn = func(a, b uint64) uint64 { return a * b }
		defaultValue = 1
	}
	result := defaultValue
	digitColumnIndex := 0
	for {
		digitColumnStr := ""
		for _, number := range numbers {
			if digitColumnIndex < len(number) {
				digitColumnStr += string(number[digitColumnIndex])
			}
		}
		if digitColumnStr == "" {
			break
		}
		result = operationFn(result, utils.ParseUint64OrPanic(strings.TrimSpace(digitColumnStr)))
		digitColumnIndex++
	}
	return result
}

func parseProblems(contents []byte) Problems {
	contentsStr := string(contents)
	lines := strings.Split(contentsStr, "\n")
	operationsLine := lines[len(lines)-1]
	operationIndices := findIndicesOfNonSpaceCharacters(operationsLine)
	numbers := make([][]string, len(operationIndices))
	for i := range numbers {
		numbers[i] = make([]string, 0, len(lines)-1)
	}
	for row := 0; row < len(lines)-1; row++ {
		for col := 0; col < len(operationIndices); col++ {
			sliced := sliceUntilNextSpace(lines[row], operationIndices[col])
			numbers[col] = append(numbers[col], sliced)
		}
	}
	operations := make([]byte, len(operationIndices))
	for i := range operationIndices {
		operations[i] = operationsLine[operationIndices[i]]
	}
	return Problems{
		numbers:    numbers,
		operations: operations,
	}
}

func findIndicesOfNonSpaceCharacters(s string) []int {
	indices := []int{}
	for i, char := range s {
		if char != ' ' {
			indices = append(indices, i)
		}
	}
	return indices
}

func sliceUntilNextSpace(s string, startIndex int) string {
	endIndex := startIndex
	for s[endIndex] == ' ' {
		endIndex++
	}
	for endIndex < len(s) && s[endIndex] != ' ' {
		endIndex++
	}
	return s[startIndex:endIndex]
}

func removeValue[T comparable](slice []T, value T) []T {
	n := 0
	for _, v := range slice {
		if v != value {
			slice[n] = v
			n++
		}
	}
	return slice[:n]
}
