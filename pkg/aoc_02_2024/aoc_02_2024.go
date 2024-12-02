package aoc_02_2024

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()
	var result [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		//Extract each number seperated by spaces
		numbers := strings.Fields(text)
		var row []int
		for i := range numbers {
			value, err := strconv.Atoi(numbers[i])
			if err != nil {
				log.Fatalf("Error parsing number: %v", err)
			}
			row = append(row, value)
		}
		result = append(result, row)
	}
	return result
}

func isSafe(numbers []int) bool {
	isAscending := numbers[0] < numbers[1]
	for i := 0; i < len(numbers)-1; i++ {
		diff := math.Abs(float64(numbers[i]) - float64(numbers[i+1]))

		if diff < 1 || diff > 3 {
			return false
		}

		if numbers[i] == numbers[i+1] {
			return false
		}

		if numbers[i] > numbers[i+1] && isAscending {
			return false
		}

		if numbers[i] < numbers[i+1] && !isAscending {
			return false
		}

	}
	return true
}

func Solve() {
	numbers := readFile("data/02_2024.txt")
	safeCount := 0
	for i := range numbers {
		if isSafe(numbers[i]) {
			safeCount++
		}
	}
	fmt.Printf("Day 2: %d\n", safeCount)
}
