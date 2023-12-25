package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var colors = []string{"red", "green", "blue"}

func getStruct() map[string]int {
	return map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
}

func solve(line string) (int, int) {
	fewestCubes := getStruct()

	power := 0
	possibleConfiguration := true

	game := strings.Split(line, ":")
	rounds := strings.Split(game[1], ";")
	for j := range rounds {
		cubesInRound := getStruct()

		cubes := strings.Split(rounds[j], ",")
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			parts := strings.Split(cube, " ")
			num, _ := strconv.Atoi(parts[0])
			cubesInRound[parts[1]] += num
		}

		if cubesInRound["red"] > 12 ||
			cubesInRound["green"] > 13 ||
			cubesInRound["blue"] > 14 {
			possibleConfiguration = false
		}

		for _, color := range colors {
			if cubesInRound[color] > fewestCubes[color] {
				fewestCubes[color] = cubesInRound[color]
			}
		}

		power = fewestCubes["red"] * fewestCubes["green"] * fewestCubes["blue"]
	}

	if !possibleConfiguration {
		return 0, power
	}

	gid := strings.Split(game[0], " ")
	ID, _ := strconv.Atoi(gid[1])
	return ID, power
}

func main() {
	readFile, err := os.Open("assets/day02.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	part1 := 0
	part2 := 0
	for fileScanner.Scan() {
		sumIDs, sumOfPowerOfIDs := solve(fileScanner.Text())
		part1 += sumIDs
		part2 += sumOfPowerOfIDs
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}
