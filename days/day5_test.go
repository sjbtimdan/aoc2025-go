package days

import (
	_ "embed"
	"testing"

	"github.com/shoenig/test/must"
)

//go:embed test_resources/day5.txt
var day5_file []byte

func TestParseInputs(t *testing.T) {
	ranges, igredients := parseInputs(day5_file)
	must.Eq(t, []Range{{3, 5}, {10, 14}, {16, 20}, {12, 18}}, ranges)
	must.Eq(t, []uint64{1, 5, 8, 11, 17, 32}, igredients)
}

func TestRangeContains(t *testing.T) {
	r := Range{3, 5}
	must.True(t, r.Contains(3))
	must.True(t, r.Contains(4))
	must.True(t, r.Contains(5))
	must.False(t, r.Contains(2))
	must.False(t, r.Contains(6))
}

func TestRangeMergeNoOverlap(t *testing.T) {
	r1 := Range{3, 5}
	r2 := Range{6, 7}
	r3, merged := r1.Merge(r2)
	must.False(t, merged)
	must.Eq(t, r2, r3)
}

func TestRangeMergeOverlap(t *testing.T) {
	r1 := Range{3, 5}
	r2 := Range{4, 7}
	r3, merged := r1.Merge(r2)
	must.True(t, merged)
	must.Eq(t, Range{3, 7}, r3)
}

func TestDay5(t *testing.T) {
	result := Day5(day5_file)
	must.Eq(t, "3", result.Part1)
	must.Eq(t, "14", result.Part2)
}
