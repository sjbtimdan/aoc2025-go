package days

import (
	"bytes"
	"sjbtimdan/aoc2025-go/utils"
	"sort"
)

func Day5(contents []byte) utils.Answers {
	freshRanges, ingredients := parseInputs(contents)
	part1 := uint64(0)
	for _, ingredient := range ingredients {
		for _, freshRange := range freshRanges {
			if freshRange.Contains(ingredient) {
				part1++
				break
			}
		}
	}
	mergedRanges := mergeRanges(freshRanges)
	part2 := uint64(0)
	for _, mergedRange := range mergedRanges {
		part2 += mergedRange.Size()
	}
	return utils.Uint64Answers(part1, part2)
}

func mergeRanges(ranges []Range) []Range {
	sortRangesByStart(ranges)
	merged := []Range{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		lastMerged := merged[len(merged)-1]
		next := ranges[i]
		if mergedRange, ok := lastMerged.Merge(next); ok {
			merged[len(merged)-1] = mergedRange
		} else {
			merged = append(merged, next)
		}
	}
	return merged
}

type Range struct {
	Start uint64
	End   uint64
}

func (r Range) Size() uint64 {
	return r.End - r.Start + 1
}

func (r Range) Contains(value uint64) bool {
	return r.Start <= value && value <= r.End
}

func (r Range) Merge(other Range) (Range, bool) {
	if r.End < other.Start || other.End < r.Start {
		return other, false
	}
	return Range{
		Start: min(r.Start, other.Start),
		End:   max(r.End, other.End),
	}, true
}

func sortRangesByStart(ranges []Range) {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})
}

func parseInputs(contents []byte) ([]Range, []uint64) {
	lines := bytes.Split(contents, []byte("\n"))
	ranges := []Range{}
	i := 0
	for {
		line := lines[i]
		if len(line) == 0 {
			i++
			break
		}
		parts := bytes.Split(line, []byte("-"))
		ranges = append(ranges, Range{
			Start: utils.ParseUint64OrPanic(string(parts[0])),
			End:   utils.ParseUint64OrPanic(string(parts[1])),
		})
		i++
	}
	ingredients := []uint64{}
	for {
		if i == len(lines) {
			break
		}
		ingredient := utils.ParseUint64OrPanic(string(lines[i]))
		ingredients = append(ingredients, ingredient)
		i++
	}
	return ranges, ingredients
}
