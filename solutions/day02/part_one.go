package day02

import (
	"log"
	"math"
	"strconv"

	. "github.com/lucascott/advent-of-code-2025/lib"
)

func PartOne(inputFileName string) {
	input := ReadFileForDay(2, inputFileName)
	ranges := parseRanges(input)
	log.Print(ranges)
	var sum = 0
	for _, range_ := range ranges {
		sum += computeRangePartOne(range_)
	}
	log.Printf("Total sum: %d", sum)
}

func nextValue(v string) int {
	//log.Printf("processing  %s", v)
	digitsCount := len(v)
	if digitsCount%2 != 0 {
		return int(math.Pow10(digitsCount))
	}
	leftPart, rightPart := ParseInt(v[:digitsCount/2]), ParseInt(v[digitsCount/2:])
	if rightPart >= leftPart {
		leftPart++
	}
	nextDigitsCount := len(strconv.Itoa(leftPart))
	return leftPart*int(math.Pow10(nextDigitsCount)) + leftPart
}

func computeRangePartOne(range_ Range) int {
	log.Printf("Processing range: %v", range_)
	var total = 0
	var value = range_.start
	var next int
	for value <= range_.end {
		valueStr := strconv.Itoa(value)
		digitsCount := len(valueStr)
		if valueStr[:digitsCount/2] == valueStr[digitsCount/2:] {
			total += value
			//log.Printf("Found error id: %d", value)
		}
		next = nextValue(valueStr)
		value = next
	}
	return total
}
