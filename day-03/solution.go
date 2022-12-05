package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func getItemPriority(letter rune) int {
	if unicode.IsUpper(letter) {
		return int(letter) - 38
	}
    return int(letter) - 96
}

func main() {
	readFile, _  := os.Open("data.txt")
	defer readFile.Close()
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	commonRucksackItems := make(map[rune]int)

	for scanner.Scan() {
		fileLine := string(scanner.Text())
		firstRucksack := fileLine[:len(fileLine)/2]
		secondRucksack := fileLine[len(fileLine)/2:]
		
		for _, itemFromFirstRucksack := range firstRucksack {
			isCommonItemFound := false

			for _, itemFromSecondRucksack := range secondRucksack {

				if itemFromFirstRucksack == itemFromSecondRucksack {
					if _, ok := commonRucksackItems[itemFromFirstRucksack]; ok {
						commonRucksackItems[itemFromFirstRucksack] += getItemPriority(itemFromFirstRucksack) 
					} else {
						commonRucksackItems[itemFromFirstRucksack] = getItemPriority(itemFromFirstRucksack) 
					}
					
					isCommonItemFound = true
					break;
				}
			}

			if isCommonItemFound {
				break;
			}
		}
	}

	sumOfPriorities := 0
	for _, value := range commonRucksackItems {
		sumOfPriorities += value
	}

	fmt.Println(sumOfPriorities)
}
