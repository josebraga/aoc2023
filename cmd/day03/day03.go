package main

import (
	"bufio"
	"fmt"
	"os"
)

var buffer = []string{"", "", ""}

func topOfThePuzzle() bool {
	return buffer[0] == "" && buffer[1] != ""
}

func bottomOfThePuzzle() bool {
	return buffer[1] != "" && buffer[2] == ""
}

func middleOfThePuzzle() bool {
	return buffer[0] != "" && buffer[1] != "" && buffer[2] != ""
}

type ANumber struct {
	value      int
	startIndex int
	endIndex   int
}

type AGear struct {
	lineNumber int
	index      int
	numbers    []int
}

var gears = map[int]map[int][]int{}

func parseNumbers(line string) []ANumber {
	var numbers []ANumber

	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			number := ANumber{startIndex: i}
			for ; i < len(line) && line[i] >= '0' && line[i] <= '9'; i++ {
				number.value = number.value*10 + int(line[i]-'0')
			}
			number.endIndex = i - 1
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func containsSymbols(line string, startIndex int, endIndex int, number int, lineNumber int) bool {
	start := startIndex - 1
	if start < 0 {
		start = 0
	}

	end := endIndex + 1
	if end >= len(line) {
		end = len(line) - 1
	}

	contains := false
	for i := start; i <= end; i++ {
		if (line[i] < '0' || line[i] > '9') && line[i] != '.' {
			// for part 2, track gears
			trackGears(line[i], i, number, lineNumber)
			contains = true
		}
	}
	return contains
}

func containsSymbolInPlace(line string, index int, number int, lineNumber int) bool {
	if index < 0 || index >= len(line) {
		return false
	}

	// for part 2
	trackGears(line[index], index, number, lineNumber)

	return (line[index] < '0' || line[index] > '9') && line[index] != '.'
}

func trackGears(char byte, index int, number int, lineNumber int) {
	if char == '*' {
		val, ok := gears[lineNumber]
		if !ok {
			gears[lineNumber] = make(map[int][]int)
			gears[lineNumber][index] = []int{number}
			return
		}

		nums, ok := val[index]
		if !ok {
			gears[lineNumber][index] = []int{number}
			return
		}

		gears[lineNumber][index] = append(nums, number)
	}
}

func process(line string, lineNumber int) int {
	buffer[0] = buffer[1]
	buffer[1] = buffer[2]
	buffer[2] = line

	lineSum := 0
	// top of the line
	if topOfThePuzzle() {
		numbers := parseNumbers(buffer[1])
		for _, number := range numbers {
			// line below
			if containsSymbols(buffer[2], number.startIndex, number.endIndex, number.value, lineNumber) {
				lineSum += number.value
			}

			// same line
			if (number.startIndex > 0 && containsSymbolInPlace(buffer[1], number.startIndex-1, number.value, lineNumber-1)) ||
				(number.endIndex < len(buffer[1])-1 && containsSymbolInPlace(buffer[1], number.endIndex+1, number.value, lineNumber-1)) {
				lineSum += number.value
			}
		}
	} else if middleOfThePuzzle() {
		numbers := parseNumbers(buffer[1])
		for _, number := range numbers {
			// line above and line below
			if containsSymbols(buffer[2], number.startIndex, number.endIndex, number.value, lineNumber) ||
				containsSymbols(buffer[0], number.startIndex, number.endIndex, number.value, lineNumber-2) {
				lineSum += number.value
			}

			// same line
			if (number.startIndex > 0 && containsSymbolInPlace(buffer[1], number.startIndex-1, number.value, lineNumber-1)) ||
				(number.endIndex < len(buffer[1])-1 && containsSymbolInPlace(buffer[1], number.endIndex+1, number.value, lineNumber-1)) {
				lineSum += number.value
			}
		}
	} else if bottomOfThePuzzle() {
		numbers := parseNumbers(buffer[1])
		for _, number := range numbers {
			// line above
			if containsSymbols(buffer[0], number.startIndex, number.endIndex, number.value, lineNumber-2) {
				lineSum += number.value
			}

			// same line
			if (number.startIndex > 0 && containsSymbolInPlace(buffer[1], number.startIndex-1, number.value, lineNumber-1)) ||
				(number.endIndex < len(buffer[1])-1 && containsSymbolInPlace(buffer[1], number.endIndex+1, number.value, lineNumber-1)) {
				lineSum += number.value
			}
		}
	}

	return lineSum
}

func main() {
	readFile, err := os.Open("assets/day03.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	lineNumber := 0
	part1 := 0
	for fileScanner.Scan() {
		part1 += process(fileScanner.Text(), lineNumber)
		lineNumber++
	}

	part1 += process("", lineNumber)

	part2 := 0
	for _, val := range gears {
		for _, nums := range val {
			if len(nums) == 2 {
				part2 += (nums[0] * nums[1])
			}
		}
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}
