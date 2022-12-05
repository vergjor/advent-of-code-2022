package main

import (
	"bufio"
	"fmt"
	"os"
)

func pointsForMove(letter string) int {
	switch letter {
	case "A":
		return 1 // Rock
	case "B":
		return 2 // Paper
	default:
		return 3 // Scissors
	}
}

func calculatePlayerPoints(opponentMove int, code string) int {
	if code == "Y" {
		return 3 + opponentMove
	} else if code == "X" {
		if opponentMove == 1 {
			return 3
		} else {
			return opponentMove - 1
		}
	} else {
		if opponentMove == 3 {
			return 7
		} else {
			return 7 + opponentMove
		}
	}
}

func main() {
	totalPoints := 0

	readFile, _  := os.Open("data.txt")
	defer readFile.Close()
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fileLine := scanner.Text()

		opponentMove := pointsForMove(string(fileLine[0]))
		code := string(fileLine[2])

		totalPoints += calculatePlayerPoints(opponentMove, code)
	}

	fmt.Println(totalPoints)
}
