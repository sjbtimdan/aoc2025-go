package days

import (
	"bytes"
	"sjbtimdan/aoc2025-go/utils"
)

func Day3(contents []byte) utils.Answers {
	banks := readBanks(contents)
	var sum1, sum2 uint
	for _, bank := range banks {
		sum1 += largestJoltage(bank, 2)
		sum2 += largestJoltage(bank, 12)
	}
	return utils.Uint64Answers(uint64(sum1), uint64(sum2))
}

func largestJoltage(banks []uint, count uint64) uint {
	var joltage uint
	steps := count
	if steps > uint64(len(banks)) {
		steps = uint64(len(banks))
	}
	index := 0
	for i := uint64(0); i < steps; i++ {
		remaining := steps - i
		nextIndex, digit := nextDigit(banks, index, int(remaining))
		joltage = joltage*10 + digit
		index = nextIndex
	}
	return joltage
}

func nextDigit(vs []uint, index, remaining int) (int, uint) {
	sliceLen := len(vs) - index - remaining + 1
	bestOffset := 0
	bestDigit := vs[index]
	bestScore := int(bestDigit) * 100

	for offset := 1; offset < sliceLen; offset++ {
		digit := vs[index+offset]
		score := int(digit)*100 - offset
		if score > bestScore {
			bestOffset = offset
			bestDigit = digit
			bestScore = score
		}
	}

	nextIndex := index + bestOffset + 1
	return nextIndex, bestDigit
}

func readBanks(contents []byte) [][]uint {
	var banks [][]uint
	for _, line := range bytes.Split(contents, []byte("\n")) {
		bank := []uint{}
		for _, joltageByte := range line {
			joltage := uint(joltageByte - '0')
			bank = append(bank, joltage)
		}
		banks = append(banks, bank)
	}
	return banks
}
