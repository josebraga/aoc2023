package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(line string) (int, int) {
	game := strings.Split(line, ":")
	numbers := strings.Split(game[1], "|")

	cardNumbersStr := strings.Split(numbers[0], " ")
	winningNumbersStr := strings.Split(numbers[1], " ")

	cardNumbers := make([]int, len(cardNumbersStr))
	winningNumbers := make([]int, len(winningNumbersStr))

	for i, num := range cardNumbersStr {
		cardNumbers[i], _ = strconv.Atoi(num)
	}

	for i, num := range winningNumbersStr {
		winningNumbers[i], _ = strconv.Atoi(num)
	}

	points := 0
	matchingNumbers := 0
	for _, num := range cardNumbers {
		if num == 0 {
			continue
		}

		for _, winningNum := range winningNumbers {
			if num == winningNum {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}

				matchingNumbers++
				break
			}
		}
	}

	return points, matchingNumbers
}

func main() {
	readFile, err := os.Open("assets/day04.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	totalPoints := 0
	cardIndex := 0
	scratchcards := make([]int, 500)
	for fileScanner.Scan() {
		scratchcards[cardIndex]++
		points, matchingNumbers := solve(fileScanner.Text())
		totalPoints += points

		copiesOfCard := scratchcards[cardIndex]

		// we spread the matching numbers to the next cards,
		// as many times as copies of cards we have
		cardIndex++
		for i := cardIndex; i < cardIndex+matchingNumbers; i++ {
			scratchcards[i] += copiesOfCard
		}
	}

	// now just sum the scratchcard total
	totalScratchcards := 0
	scratchcards = scratchcards[:cardIndex]
	for _, card := range scratchcards {
		totalScratchcards += card
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", totalPoints, totalScratchcards)
}
