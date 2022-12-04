package main

import (
	"bufio"
	"fmt"
	"os"
)

func decodeMove(letter string) int {
	switch letter {
		case "A", "X":
			return 1 // Rock
		case "B", "Y":
			return 2 // Paper
		default:
			return 3 // Scissors
	}
}

func findWinner(playerMove int, opponentMove int) int {
	if playerMove == opponentMove {
		return 3
	} else if opponentMove < playerMove && playerMove - opponentMove == 1 || playerMove == 1 && opponentMove == 3 {
		return 6
	}
	return 0
}



func main() {
	totalPoints := 0

	readFile, _  := os.Open("data.txt")
	defer readFile.Close()
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fileLine := scanner.Text()

		playerMove := decodeMove(string(fileLine[2]))
		opponentMove := decodeMove(string(fileLine[0]))

		totalPoints += (findWinner(playerMove, opponentMove) + playerMove)
	}

	fmt.Println(totalPoints)
}
