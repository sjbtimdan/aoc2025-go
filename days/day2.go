package days

import (
	"bytes"
	"math"
	"sjbtimdan/aoc2025-go/utils"
	"strconv"
)

type Interval struct {
	start    int64
	startStr string
	end      int64
	endStr   string
}

func Day2(contents []byte) utils.Answers {
	part1 := int64(0)
	for _, line := range bytes.Split(contents, []byte(",")) {
		part1 += sumAllInvalidIds(parseInterval(line))
	}
	return utils.Int64Answers(part1, 0)
}

func sumAllInvalidIds(interval Interval) int64 {
	sum := int64(0)
	equalLengthIntervals := splitIntoEqualDigitLengthIntervals(interval)
	for _, interval := range equalLengthIntervals {
		sum += sumInvalidIds(interval)
	}
	return sum
}

func splitIntoEqualDigitLengthIntervals(interval Interval) []Interval {
	var intervals []Interval
	for length := len(interval.startStr); length <= len(interval.endStr); length++ {
		start := int64(math.Pow10(int(length - 1)))
		end := int64(math.Pow10(int(length))) - 1
		intervals = append(intervals, toIntervalFromInt64(start, end))
	}
	intervals[0] = toIntervalFromInt64(interval.start, intervals[0].end)
	intervals[len(intervals)-1] = toIntervalFromInt64(intervals[len(intervals)-1].start, interval.end)
	return intervals
}

func sumInvalidIds(interval Interval) int64 {
	if len(interval.startStr)%2 == 1 {
		return 0
	}
	startPrefixStr := interval.startStr[:len(interval.startStr)/2]
	endPrefixStr := interval.endStr[:len(interval.endStr)/2]
	incrementLen := len(interval.startStr) / 2
	increment := int64(math.Pow10(incrementLen)) + 1
	start := utils.StrconvOrPanic(startPrefixStr + startPrefixStr)
	if start < interval.start {
		start += increment
	}
	end := utils.StrconvOrPanic(endPrefixStr + endPrefixStr)
	if end > interval.end {
		end -= increment
	}
	n := (end-start)/increment + 1
	sum := n * (start + end) / 2
	return sum
}

func parseInterval(line []byte) Interval {
	parts := bytes.Split(line, []byte("-"))
	startStr := string(parts[0])
	endStr := string(parts[1])
	return toInterval(startStr, endStr)
}

func toInterval(startStr, endStr string) Interval {
	return Interval{
		start:    utils.StrconvOrPanic(startStr),
		startStr: startStr,
		end:      utils.StrconvOrPanic(endStr),
		endStr:   endStr,
	}
}

func toIntervalFromInt64(start, end int64) Interval {
	return Interval{
		start:    start,
		startStr: strconv.FormatInt(start, 10),
		end:      end,
		endStr:   strconv.FormatInt(end, 10),
	}
}
