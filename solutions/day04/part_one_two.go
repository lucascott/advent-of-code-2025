package day04

import (
	"log"
	"strings"

	. "github.com/lucascott/advent-of-code-2025/lib"
)

const (
	// nonAccessibilityThreshold is the minimum number of rolls surrounding a roll of paper
	// that make it not accessible by the forklift
	nonAccessibilityThreshold = 4
	// freeSpace marks a space without a roll of paper in it in the map
	freeSpace = rune(46) // . char
	// paperRollRune marks a roll of paper in the map
	paperRollRune = rune(64) // @ char
	// paperRollAccessibleRune represents a roll of paper in the map. the roll is accessible by the forklift.
	paperRollAccessibleRune = rune(164) // ¤ char
)

// Cell represents a cell in the map.
// The cell struct has a reference to its map for advanced operations.
type Cell struct {
	x    int
	y    int
	map_ *[][]rune
}

// isAccessible verifies if the cell is not surrounded by nonAccessibilityThreshold number of paper rolls
func (cell *Cell) isAccessible() bool {
	map_ := *(*cell).map_
	shapeN := len(map_)
	shapeM := len(map_[0])
	if shapeN == 0 || shapeM == 0 {
		log.Fatal("The map is empty")
	}

	paperCount := 0
	for offsetX := -1; offsetX <= 1; offsetX++ {
		for offsetY := -1; offsetY <= 1; offsetY++ {
			c := Cell{(*cell).x + offsetX, (*cell).y + offsetY, (*cell).map_}
			if c.isValid() && !c.equals(cell) && c.isPaper() {
				paperCount++
			}
			if paperCount >= nonAccessibilityThreshold {
				return false
			}
		}
	}
	return true
}

// isValid checks whether the cell is within the map
func (cell *Cell) isValid() bool {
	map_ := *(*cell).map_
	maxRows := len(map_)
	maxCols := len(map_[0])
	if cell.x >= 0 && cell.y >= 0 && cell.x < maxRows && cell.y < maxCols {
		return true
	}
	return false
}

// isPaper checks if the cell contains a roll of paper
func (cell *Cell) isPaper() bool {
	c := *cell
	return (*c.map_)[c.x][c.y] == paperRollRune || (*c.map_)[c.x][c.y] == paperRollAccessibleRune
}

// equals checks whether two cells are actually the same cell (same coordinates)
func (cell *Cell) equals(cell2 *Cell) bool {
	return (*cell).x == (*cell2).x && (*cell).y == (*cell2).y
}

func readMap(input string) [][]rune {
	rows := strings.Split(input, "\n")
	var map_ [][]rune
	for _, row := range rows {
		var bank []rune
		for _, item := range []rune(row) {
			bank = append(bank, item)
		}
		map_ = append(map_, bank)
	}
	return map_
}

func printMap(map_ [][]rune) {
	for _, row := range map_ {
		log.Println(string(row))
	}
}

func findAccessibleRolls(paperMap *[][]rune) int {
	accessibleRolls := 0
	for m := 0; m < len(*paperMap); m++ {
		for n := 0; n < len((*paperMap)[0]); n++ {
			c := Cell{m, n, paperMap}
			if c.isPaper() && c.isAccessible() {
				(*paperMap)[m][n] = paperRollAccessibleRune
				accessibleRolls++
			}
		}
	}
	return accessibleRolls
}

func removeAccessibleRolls(paperMap *[][]rune) {
	for m := 0; m < len(*paperMap); m++ {
		for n := 0; n < len((*paperMap)[0]); n++ {
			if (*paperMap)[m][n] == paperRollAccessibleRune {
				(*paperMap)[m][n] = freeSpace
			}
		}
	}
}

func PartOne(inputFileName string) int {
	input := ReadFileForDay(4, inputFileName)
	paperMap := readMap(input)
	printMap(paperMap)
	log.Println("")
	accessibleRolls := findAccessibleRolls(&paperMap)
	printMap(paperMap)
	log.Println("")
	log.Printf("Total accessible rolls: %d", accessibleRolls)
	return accessibleRolls
}

func PartTwo(inputFileName string) int {
	input := ReadFileForDay(4, inputFileName)
	paperMap := readMap(input)
	printMap(paperMap)
	log.Println("")
	totalRemovedRolls := 0
	removedRolls := -1
	for removedRolls != 0 {
		removedRolls = findAccessibleRolls(&paperMap)
		removeAccessibleRolls(&paperMap)
		totalRemovedRolls += removedRolls
	}
	printMap(paperMap)
	log.Println("")
	log.Printf("Total removed rolls: %d", totalRemovedRolls)
	return totalRemovedRolls
}
