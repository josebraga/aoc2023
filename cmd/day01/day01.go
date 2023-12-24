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

func getNumber(str string) int {
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

	f2, fi2, l2, li2 := spelledNumber(str)

	if (fi2 < firstIndex && fi2 != -1) || firstIndex == -1 {
		first = f2
	}

	if (li2 > lastIndex && li2 != -1) || lastIndex == -1 {
		last = l2
	}

	return first*10 + last
}

func main() {
	readFile, err := os.Open("assets/day01.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	total := 0
	for fileScanner.Scan() {
		total += getNumber(fileScanner.Text())
	}

	fmt.Println("Total is", total)

	readFile.Close()
}
