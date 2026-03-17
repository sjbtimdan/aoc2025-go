package days

import (
	_ "embed"
	"testing"

	"github.com/shoenig/test/must"
)

//go:embed test_resources/day6.txt
var day6_file []byte

func TestParseProblems(t *testing.T) {
	problems := parseProblems(day6_file)
	expected_numbers := [][]string{
		{"123", " 45", "  6"},
		{"328", "64", "98"},
		{" 51", "387", "215"},
		{"64", "23", "314"},
	}
	expected_operations := []byte{'*', '+', '*', '+'}
	must.Eq(t, expected_numbers, problems.numbers)
	must.Eq(t, expected_operations, problems.operations)
}

func TestCalculateColumnResult(t *testing.T) {
	numbers := []string{"2", "3", "5"}
	must.Eq(t, uint64(10), calculateColumnResult(numbers, '+'))
	must.Eq(t, uint64(30), calculateColumnResult(numbers, '*'))
}

func TestCalculateRightToLeftColumnResult(t *testing.T) {
	must.Eq(t, uint64(1058), calculateRightToLeftColumnResult([]string{"64", "23", "314"}, '+'))
	must.Eq(t, uint64(3253600), calculateRightToLeftColumnResult([]string{" 51", "387", "215"}, '*'))
}

func TestDay6(t *testing.T) {
	result := Day6(day6_file)
	must.Eq(t, "4277556", result.Part1)
	must.Eq(t, "3263827", result.Part2)
}
