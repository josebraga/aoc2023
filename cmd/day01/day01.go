package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var numbers = map[string]int{
	"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

var spelledNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func findOccurrences(line string, num map[string]int) (int, int, int, int) {
	first := 0
	firstIndex := -1
	last := 0
	lastIndex := -1

	for substr, number := range num {
		if index := strings.Index(line, substr); index != -1 {
			if index < firstIndex || firstIndex == -1 {
				first = number
				firstIndex = index
			}
		}
	}

	for substr, number := range num {
		if index := strings.LastIndex(line, substr); index != -1 {
			if index > lastIndex || lastIndex == -1 {
				last = number
				lastIndex = index
			}
		}
	}
	return first, firstIndex, last, lastIndex
}

func solve(line string) (int, int) {
	first, firstIndexPart1, last, lastIndexPart1 := findOccurrences(line, numbers)
	answerPart1 := first*10 + last

	firstPart2, firstIndexPart2, lastPart2, lastIndexPart2 := findOccurrences(line, spelledNumbers)

	if (firstIndexPart2 < firstIndexPart1 && firstIndexPart2 != -1) || firstIndexPart1 == -1 {
		first = firstPart2
	}

	if lastIndexPart2 > lastIndexPart1 && lastIndexPart2 != -1 {
		last = lastPart2
	}

	return answerPart1, first*10 + last
}

func main() {
	readFile, err := os.Open("assets/day01.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	part1 := 0
	part2 := 0
	for fileScanner.Scan() {
		t1, t2 := solve(fileScanner.Text())
		part1 += t1
		part2 += t2
	}
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}
