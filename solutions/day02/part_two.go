package day02

import (
	"log"
	"strconv"

	. "github.com/lucascott/advent-of-code-2025/lib"
)

var sliceSizesToCheck = map[int][]int{}

func PartTwo(inputFileName string) int {
	input := ReadFileForDay(2, inputFileName)
	ranges := parseRanges(input)
	log.Print(ranges)
	var sum = 0
	for _, range_ := range ranges {
		sum += computeRangePartTwo(range_)
	}
	log.Printf("Total sum: %d", sum)
	return sum
}

func computeRangePartTwo(range_ Range) int {
	log.Printf("Processing range: %v", range_)
	var total = 0
	for i := range_.start; i <= range_.end; i++ {
		if isInvalidId(i) {
			log.Printf("Found value: %d", i)
			total += i
		}
	}
	return total
}

func divisors(n int) []int {
	var divs = []int{1}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			divs = append(divs, i)
		}
	}
	return divs
}

func areSlicesAllEqual(s string, sliceSize int) bool {
	slice := s[:sliceSize]
	for i := sliceSize; i < len(s); i += sliceSize {
		if slice != s[i:i+sliceSize] {
			return false
		}
	}
	return true
}

func isInvalidId(id int) bool {
	idStr := strconv.Itoa(id)
	digitsCount := len(idStr)
	// single digits are always valid
	if digitsCount == 1 {
		return false
	}
	validSizes := getSizes(digitsCount)
	for _, sliceSize := range validSizes {
		if areSlicesAllEqual(idStr, sliceSize) {
			return true
		}
	}
	return false

}

func getSizes(count int) []int {
	i, ok := sliceSizesToCheck[count]
	if ok {
		return i
	}
	copies := divisors(count)
	sliceSizesToCheck[count] = copies
	return copies
}
