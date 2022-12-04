package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
)

func main() {
	var highestCalorieCount, currentElfCalorieCount int = 0, 0

	readFile, _  := os.Open("data.txt")
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fileLine := fmt.Sprintf("%v", scanner.Text())

		if len(fileLine) == 0 {
			if highestCalorieCount < currentElfCalorieCount {
				highestCalorieCount = currentElfCalorieCount
			}
			currentElfCalorieCount = 0
			continue;
		}

		currentFoodCalories, _ := strconv.Atoi(fileLine)
		currentElfCalorieCount += currentFoodCalories
	}

	fmt.Println(highestCalorieCount)
	readFile.Close()
}