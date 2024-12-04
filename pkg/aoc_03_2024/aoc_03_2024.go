package aoc_03_2024

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseDoDonts(input string) []string {
	var matchedStr []string
	dontArr := strings.Split(input, "don't()")
	matchedStr = append(matchedStr, dontArr[0])
	for i := 1; i < len(dontArr); i++ {
		doArr := strings.Split(dontArr[i], "do()")
		fmt.Printf("Do Arr Size: %v\n", len(doArr))
		if len(doArr) == 1 {
			continue
		} else {
			for j := 1; j < len(doArr); j++ {
				matchedStr = append(matchedStr, doArr[j])
			}
		}
	}
	return matchedStr
}

func parseString(input string) int {
	total := 0
	doDonts := parseDoDonts(input)
	// Convert each character to array
	for i := 0; i < len(doDonts); i++ {
		mulArr := strings.Split(doDonts[i], "mul(")
		for i := 0; i < len(mulArr); i++ {
			firstNum, secondNum := 0, 0
			commaArr := strings.Split(mulArr[i], ",")
			if len(commaArr) < 2 {
				continue
			}
			firstNum, err := strconv.Atoi(commaArr[0])
			if err != nil {
				continue
			}
			closingArr := strings.Split(commaArr[1], ")")
			secondNum, err = strconv.Atoi(closingArr[0])
			if err != nil {
				continue
			}
			total = total + (firstNum * secondNum)
		}
	}
	return total
}

func Solve() {
	data, err := os.ReadFile("data/03_2024.txt")
	if err != nil {
		log.Fatalf("Error processing file: %v", err)
	}
	string := string(data)
	total := parseString(string)
	fmt.Printf("Day 3: %d\n", total)
}
