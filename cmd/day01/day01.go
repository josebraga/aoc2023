package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var numbers = map[string]int{
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

func isNumber(char byte) bool {
	return char >= '0' && char <= '9'
}

func spelledNumber(str string) (int, int, int, int) {
	first := 0
	firstIndex := -1
	last := 0
	lastIndex := -1

	for s, i := range numbers {
		if index := strings.Index(str, s); index != -1 {
			if index < firstIndex || firstIndex == -1 {
				first = i
				firstIndex = index
			}
		}
	}

	for s, i := range numbers {
		if index := strings.LastIndex(str, s); index != -1 {
			if index > lastIndex || lastIndex == -1 {
				last = i
				lastIndex = index
			}
		}
	}
	return first, firstIndex, last, lastIndex
}

func getNumber(str string) (int, int) {
	first := 0
	firstIndex := -1
	last := 0
	lastIndex := -1

	for i := 0; i < len(str); i++ {
		if isNumber(str[i]) {
			first = int(str[i] - '0')
			firstIndex = i
			break
		}
	}

	for i := len(str) - 1; i >= 0; i-- {
		if isNumber(str[i]) {
			last = int(str[i] - '0')
			lastIndex = i
			break
		}
	}

	answerPart1 := first*10 + last

	firstPart2, firstIndexPart2, lastPart2, lastIndexPart2 := spelledNumber(str)

	if (firstIndexPart2 < firstIndex && firstIndexPart2 != -1) || firstIndex == -1 {
		first = firstPart2
	}

	if (lastIndexPart2 > lastIndex && lastIndexPart2 != -1) || lastIndex == -1 {
		last = lastPart2
	}

	return answerPart1, first*10 + last
}

func main() {
	readFile, err := os.Open("assets/day01.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	totalPart1 := 0
	totalPart2 := 0
	for fileScanner.Scan() {
		t1, t2 := getNumber(fileScanner.Text())
		totalPart1 += t1
		totalPart2 += t2
	}
	fmt.Printf("Part 1: %d\nPart 2: %d\n", totalPart1, totalPart2)

	readFile.Close()
}
