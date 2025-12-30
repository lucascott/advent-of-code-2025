package day04

import (
	"testing"

	. "github.com/lucascott/advent-of-code-2025/lib"
	"github.com/stretchr/testify/assert"
)

const day = 4

func TestPartOne(t *testing.T) {
	var expected int
	expected = ParseInt(ReadFileForDay(day, "small_result_part_one.txt"))
	assert.Equal(t, expected, PartOne(SmallInput))

	expected = ParseInt(ReadFileForDay(day, "large_result_part_one.txt"))
	assert.Equal(t, expected, PartOne(LargeInput))
}

func TestPartTwo(t *testing.T) {
	var expected int
	expected = ParseInt(ReadFileForDay(day, "small_result_part_two.txt"))
	assert.Equal(t, expected, PartTwo(SmallInput))

	expected = ParseInt(ReadFileForDay(day, "large_result_part_two.txt"))
	assert.Equal(t, expected, PartTwo(LargeInput))
}
