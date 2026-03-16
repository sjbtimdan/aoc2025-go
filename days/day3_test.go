package days

import (
	_ "embed"
	"testing"

	"github.com/shoenig/test/must"
)

func TestReadBanks(t *testing.T) {
	banks := readBanks([]byte("123"))
	must.Eq(t, [][]uint{{1, 2, 3}}, banks)
}

func TestLargestJoltage(t *testing.T) {
	must.Eq(t, uint(89), largestJoltage([]uint{8, 1, 1, 9}, 2))
	must.Eq(t, uint(434234234278), largestJoltage([]uint{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, 12))
}

//go:embed test_resources/day3.txt
var day3_file []byte

func TestDay3(t *testing.T) {
	result := Day3(day3_file)
	must.Eq(t, "357", result.Part1)
	must.Eq(t, "3121910778619", result.Part2)
}
