package aoc_03_2024

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseString(input string) int {
	// Convert each character to array
	mulArr := strings.Split(input, "mul(")
	total := 0
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
