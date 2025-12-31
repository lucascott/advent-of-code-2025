package day05

import (
	"log"
	"slices"
	"strings"

	. "github.com/lucascott/advent-of-code-2025/lib"
)

type FreshRange struct {
	start int
	end   int
}

func (r FreshRange) contains(value int) bool {
	return value >= r.start && value <= r.end
}
func (r FreshRange) size() int {
	return r.end - r.start + 1
}

func readIngredients(input string) ([]FreshRange, []int) {
	rawSplit := strings.SplitN(input, "\n\n", 2)
	rawFreshRanges, rawIngredients := rawSplit[0], rawSplit[1]
	return parseRanges(rawFreshRanges), parseIngredients(rawIngredients)
}

func parseIngredients(ingredientsInput string) []int {
	rawIngredients := strings.Split(ingredientsInput, "\n")
	ingredients := make([]int, len(rawIngredients))
	for i, rawIngredient := range rawIngredients {
		ingredients[i] = ParseInt(rawIngredient)
	}
	return ingredients
}

func parseRanges(rangesInput string) []FreshRange {
	rawRanges := strings.Split(rangesInput, "\n")
	ranges := make([]FreshRange, len(rawRanges))
	for i, rawRange := range rawRanges {
		parts := strings.SplitN(rawRange, "-", 2)
		ranges[i] = FreshRange{start: ParseInt(parts[0]), end: ParseInt(parts[1])}
	}
	return ranges
}

func isFresh(id int, ranges *[]FreshRange) bool {
	for _, r := range *ranges {
		if r.contains(id) {
			return true
		}
	}
	return false
}

func lastItem[T any](collection *[]T) (*T, bool) {
	var zero T
	if collection == nil || len(*collection) == 0 {
		return &zero, false
	}
	return &(*collection)[len(*collection)-1], true
}

func PartOne(inputFileName string) int {
	input := ReadFileForDay(5, inputFileName)
	freshRanges, ings := readIngredients(input)
	freshIngredients := 0
	for _, id := range ings {
		if isFresh(id, &freshRanges) {
			freshIngredients++
		}
	}
	log.Printf("Total fresh ingredients available: %d", freshIngredients)
	return freshIngredients
}

func PartTwo(inputFileName string) int {
	input := ReadFileForDay(5, inputFileName)
	freshRanges, _ := readIngredients(input)
	freshIngredientsInDatabase := countFreshIngredientsInDatabase(&freshRanges)
	log.Printf("Total fresh ingredients in database: %d", freshIngredientsInDatabase)
	return freshIngredientsInDatabase
}

func countFreshIngredientsInDatabase(ranges *[]FreshRange) int {
	if len(*ranges) == 0 {
		panic("No fresh ingredients range available")
	}
	// sort the ranges list based on the ranges' lowest boundary
	slices.SortFunc(*ranges, func(a, b FreshRange) int {
		if a.start < b.start {
			return -1
		} else if a.start == b.start {
			return 0
		}
		return 1
	})
	mergedRanges := []FreshRange{(*ranges)[0]}
	for _, freshRange := range *ranges {
		last, _ := lastItem(&mergedRanges)
		// we can use last.end+1 as the boundaries are inclusive
		// so that non overlapping contiguous ranges are merged together
		// i.e. {1, 2} and {3, 4} are merged into {1, 4}
		if freshRange.start > last.end+1 {
			mergedRanges = append(mergedRanges, freshRange)
		} else if freshRange.end > last.end {
			last.end = freshRange.end
		}
	}
	freshIngredientsInDatabase := 0
	for _, mergedRange := range mergedRanges {
		freshIngredientsInDatabase += mergedRange.size()
	}
	return freshIngredientsInDatabase
}
