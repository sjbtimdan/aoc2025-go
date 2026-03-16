package days

import (
	"os"
	"strconv"
	"testing"

	"github.com/shoenig/test/must"
)

func TestParseInterval(t *testing.T) {
	interval := parseInterval([]byte("123-456"))
	must.Eq(t, Interval{
		start:    123,
		startStr: "123",
		end:      456,
		endStr:   "456",
	}, interval)
}

func TestCountInvalidIdsPart1(t *testing.T) {
	must.Eq(t, int64(0), part1SumInvalidIds(1698522, 1698528))
	must.Eq(t, int64(1188511885), part1SumInvalidIds(1188511880, 1188511890))
	must.Eq(t, int64(1010), part1SumInvalidIds(998, 1012))
	must.Eq(t, int64(0), part1SumInvalidIds(1698522, 1698528))
}

func TestDay2(t *testing.T) {
	contents, err := os.ReadFile("../test_resources/day2.txt") // Go 1.16+
	if err != nil {
		panic(err) // abort on error
	}
	result := Day2(contents)
	must.Eq(t, "1227775554", result.Part1)
}

func part1SumInvalidIds(start, end int64) int64 {
	interval := toInterval(strconv.FormatInt(start, 10), strconv.FormatInt(end, 10))
	return sumAllInvalidIds(interval)
}
