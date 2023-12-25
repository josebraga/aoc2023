package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func split(str string) (int, int) {
	fewestCubes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	power := 0
	possible := true

	game := strings.Split(str, ":")
	for i := range game {
		game[i] = strings.TrimSpace(game[i])
	}

	rounds := strings.Split(game[1], ";")
	for j := range rounds {
		cubesInRound := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		rounds[j] = strings.TrimSpace(rounds[j])

		cubes := strings.Split(rounds[j], ",")
		for k := range cubes {
			cubes[k] = strings.TrimSpace(cubes[k])

			cube := strings.Split(cubes[k], " ")
			num, err := strconv.Atoi(cube[0])
			if err != nil {
				fmt.Println(err)
			}

			cubesInRound[cube[1]] += num
		}

		if cubesInRound["red"] > 12 ||
			cubesInRound["green"] > 13 ||
			cubesInRound["blue"] > 14 {
			possible = false
		}

		if cubesInRound["red"] > fewestCubes["red"] {
			fewestCubes["red"] = cubesInRound["red"]
		}

		if cubesInRound["green"] > fewestCubes["green"] {
			fewestCubes["green"] = cubesInRound["green"]
		}

		if cubesInRound["blue"] > fewestCubes["blue"] {
			fewestCubes["blue"] = cubesInRound["blue"]
		}

		power = fewestCubes["red"] * fewestCubes["green"] * fewestCubes["blue"]
	}

	if possible {
		gameId := strings.Split(game[0], " ")
		ID, err := strconv.Atoi(gameId[1])
		if err != nil {
			fmt.Println(err)
		}

		return ID, power
	}

	return 0, power
}

func main() {
	readFile, err := os.Open("assets/day02.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	total := 0
	power := 0
	for fileScanner.Scan() {
		t, p := split(fileScanner.Text())
		total += t
		power += p
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", total, power)
	readFile.Close()
}
