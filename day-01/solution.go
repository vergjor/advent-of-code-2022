package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
)

func isCalorieStackPopulated(calorieStack [3]int) bool {
	return calorieStack[0] != 0 && calorieStack[1] != 0 && calorieStack[2] != 0
}

func sortCalories(calorieStack *[3]int) {
	calories := *calorieStack
	
	for index := 2; index >= 0; index-- {
		for i, value := range calories[:index] {
			if value < calories[index] {
				tmp := calories[index]
				calories[index] = value
				calories[i] = tmp
			}
		}
	}

	*calorieStack = calories
}

func populateCalorieStack(calorieCount int, calorieStack *[3]int) {
	calories := *calorieStack
	
	for index, value := range calories {
		if value == 0 && (calorieStack[0] != calorieCount || calorieStack[1] != calorieCount || calorieStack[2] != calorieCount) {
			calories[index] = calorieCount
			sortCalories(&calories)
			*calorieStack = calories
			break;
		}
	}
}

func updateTopThreeElvenCalorieCount(calorieCount int, calorieStack *[3]int) {
	calories := *calorieStack

	for index := 2; index >= 0; index-- {
		if calories[index] < calorieCount {
			calories[index] = calorieCount
			break;
		}
	}

	sortCalories(&calories)
	*calorieStack = calories
}

func main() {
	var topThreeStackedElves [3]int
	currentElfCalorieCount, totalCaloriesInStack := 0, 0

	readFile, _  := os.Open("data.txt")
	defer readFile.Close()
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fileLine := fmt.Sprintf("%v", scanner.Text())

		if len(fileLine) == 0 {
			if isCalorieStackPopulated(topThreeStackedElves) {
				updateTopThreeElvenCalorieCount(currentElfCalorieCount, &topThreeStackedElves)
			} else {
				populateCalorieStack(currentElfCalorieCount, &topThreeStackedElves)
			}
			currentElfCalorieCount = 0
			continue;
		}

		currentFoodCalories, _ := strconv.Atoi(fileLine)
		currentElfCalorieCount += currentFoodCalories
	}

	for _, totalCalories := range topThreeStackedElves {
		totalCaloriesInStack += totalCalories
	}

	fmt.Println(totalCaloriesInStack)
}
