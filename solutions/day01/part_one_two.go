package day01

import (
	"log"
	"strconv"
	"strings"

	. "github.com/lucascott/advent-of-code-2025/lib"
)

const (
	knobStepsSize   = 100
	initialPosition = 50
)

func PartOne(inputFileName string) int {
	input := ReadFileForDay(1, inputFileName)
	stopsAtZeroCount, _ := rotateKnob(input)
	log.Printf("Stops at zero count = %d", stopsAtZeroCount)
	return stopsAtZeroCount
}

func PartTwo(inputFileName string) int {
	input := ReadFileForDay(1, inputFileName)
	_, passAcrossZeroCount := rotateKnob(input)
	log.Printf("Pass across zero count = %d", passAcrossZeroCount)
	return passAcrossZeroCount
}

func rotateKnob(input string) (int, int) {
	var stopsAtZeroCount = 0
	var passAcrossZeroCount = 0
	var currentPosition = initialPosition
	actions := strings.SplitSeq(input, "\n")
	for row := range actions {
		var sign int
		var passesAtZero = 0
		if string(row[0]) == "L" {
			sign = -1
		} else {
			sign = 1
		}
		steps, err := strconv.Atoi(row[1:])
		if err != nil {
			log.Fatalf("Error when parsing: %s", err)
		}
		normalizedSteps := steps % knobStepsSize
		passesAtZero += steps / knobStepsSize
		newPosition := currentPosition + sign*normalizedSteps
		if (newPosition <= 0 || newPosition >= knobStepsSize) && normalizedSteps != 0 && currentPosition != 0 {
			passesAtZero++
		}
		currentPosition = (newPosition + knobStepsSize) % knobStepsSize
		log.Printf("The dial is rotated %s to point at %d (%d passes at 0)", row, currentPosition, passesAtZero)
		if currentPosition == 0 {
			stopsAtZeroCount++
		}
		passAcrossZeroCount += passesAtZero
	}
	return stopsAtZeroCount, passAcrossZeroCount
}
