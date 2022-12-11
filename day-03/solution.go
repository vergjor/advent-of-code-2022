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

func lineContainsRune(letter rune, line []rune) bool {
	containsRune := false

	for _, item := range line {
		if letter == item {
			containsRune = true
			break;
		}
 	}

	return containsRune
}

func main() {
	readFile, _  := os.Open("data.txt")
	defer readFile.Close()
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	commonRucksackItems := make(map[rune]int)
	var groups []string


	for scanner.Scan() {
		fileLine := string(scanner.Text())
		groups = append(groups, fileLine)
	}

	for i := 0; i + 3 <= len(groups); i += 3 {
		lineOneToRune := []rune(groups[i])
		lineTwoToRune := []rune(groups[i + 1])
		lineThreeToRune := []rune(groups[i + 2])

		for index := 0; index < len(lineOneToRune); index++ {
			isElementInSecondLine := lineContainsRune(lineOneToRune[index], lineTwoToRune)
			isElementInThirdLine := lineContainsRune(lineOneToRune[index], lineThreeToRune)

			if isElementInSecondLine && isElementInThirdLine {
				if _, ok := commonRucksackItems[lineOneToRune[index]]; ok {
					commonRucksackItems[lineOneToRune[index]] += getItemPriority(lineOneToRune[index])
				} else {
					commonRucksackItems[lineOneToRune[index]] = getItemPriority(lineOneToRune[index])
				}
				break
			}
		}
	}

	sumOfPriorities := 0
	for _, value := range commonRucksackItems {
		sumOfPriorities += value
	}

	fmt.Println(sumOfPriorities)
}
