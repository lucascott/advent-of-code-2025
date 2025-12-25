package day03

import (
	"cmp"
	"log"
	"slices"
	"strconv"
	"strings"

	. "github.com/lucascott/advent-of-code-2025/lib"
)

func PartOne(inputFileName string) int {
	input := ReadFileForDay(3, inputFileName)
	banks := readBanks(input)
	log.Print(banks)
	var sum = 0
	for _, bank := range banks {
		sum += getBankJoltagePartOne(bank)
		// alternatively, with solution of part two
		// sum += getBankJoltagePartTwo(bank, 2)
	}
	log.Printf("Total sum: %d", sum)
	return sum
}

func PartTwo(inputFileName string) int {
	input := ReadFileForDay(3, inputFileName)
	banks := readBanks(input)
	log.Print(banks)
	var sum = 0
	for _, bank := range banks {
		sum += getBankJoltagePartTwo(bank, 12)
	}
	log.Printf("Total sum: %d", sum)
	return sum
}

func readBanks(input string) [][]int {
	banksStr := strings.Split(input, "\n")
	var banks [][]int
	for _, bankStr := range banksStr {
		var bank []int
		for _, battery := range []byte(bankStr) {
			bank = append(bank, ParseInt(string(battery)))
		}
		banks = append(banks, bank)
	}
	return banks
}

func maxWithIndex[S interface{ ~[]E }, E cmp.Ordered](x S) (int, E) {
	if len(x) < 1 {
		panic("maxWithIndex: empty list")
	}

	maxVal := x[0]
	atIndex := 0

	for index, value := range x {
		if value > maxVal {
			maxVal = value
			atIndex = index
		}
	}
	return atIndex, maxVal
}

func getBankJoltagePartOne(bank []int) int {
	indexFirst, first := maxWithIndex(bank[:len(bank)-1])
	last := slices.Max(bank[indexFirst+1:])
	joltage := ParseInt(strconv.Itoa(first) + strconv.Itoa(last))
	log.Print("Bank: ", bank, " Joltage: ", joltage)
	return joltage
}

func getBankJoltagePartTwo(bank []int, digits int) int {
	bankSize := len(bank)
	baseIndex := 0
	joltageStr := ""
	for digitsLeft := digits; digitsLeft > 0; digitsLeft-- {
		index, maxDigit := maxWithIndex(bank[baseIndex : bankSize-digitsLeft+1])
		joltageStr += strconv.Itoa(maxDigit)
		baseIndex += index + 1
	}
	joltage := ParseInt(joltageStr)
	log.Print("Bank: ", bank, " Joltage: ", joltage)
	return joltage
}
