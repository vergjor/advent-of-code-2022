package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func toInt(value string) int {
	number, _ := strconv.Atoi(value)
	return number
}

func isInRange(firstRange []string, secondRange []string) bool {
	for number := toInt(firstRange[0]); number <= toInt(firstRange[1]); number++ {
		if number >= toInt(secondRange[0]) && number <= toInt(secondRange[1]) {
			return true
		}
	}
	return false
}

func main() {
	readFile, _  := os.Open("data.txt")
	defer readFile.Close()
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var sectionAssignmentPairs []string
	var totalOverlappingSections int = 0

	for scanner.Scan() {
		fileLine := string(scanner.Text())
		sectionAssignmentPairs = append(sectionAssignmentPairs, fileLine)
	}

	for _, pair := range sectionAssignmentPairs {
		sectionElves := strings.Split(pair, ",")
		firstElfSection := strings.Split(sectionElves[0], "-")
		secondElfSection := strings.Split(sectionElves[1], "-")

		if isInRange(firstElfSection, secondElfSection) {
			totalOverlappingSections++
		}
	}

	fmt.Println(totalOverlappingSections)
}
