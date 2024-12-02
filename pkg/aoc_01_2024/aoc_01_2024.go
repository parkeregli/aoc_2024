package aoc_01_2024

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type NumberPair struct {
	first  float64
	second float64
}

func readNumberPairs(filename string) ([]NumberPair, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pairs []NumberPair
	for scanner.Scan() {
		text := scanner.Text()
		midPoint := len(text) / 2
		first, err := strconv.ParseFloat(strings.TrimSpace(text[:midPoint]), 64)
		if err != nil {
			return nil, fmt.Errorf("Parsing first number: %w", err)
		}
		second, err := strconv.ParseFloat(strings.TrimSpace(text[midPoint:]), 64)
		if err != nil {
			return nil, fmt.Errorf("Parsing second number: %w", err)
		}
		pairs = append(pairs, NumberPair{first, second})
	}
	return pairs, nil
}

func calculateTotalDistance(pairs []NumberPair) int {
	if len(pairs) == 0 {
		return 0
	}
	firsts := make([]float64, len(pairs))
	seconds := make([]float64, len(pairs))
	for i, pair := range pairs {
		firsts[i] = pair.first
		seconds[i] = pair.second
	}
	sort.Float64s(firsts)
	sort.Float64s(seconds)
	var totalDis float64
	for i := range firsts {
		totalDis += math.Abs(firsts[i] - seconds[i])
	}
	return int(totalDis)
}

func Solve() {
	pairs, err := readNumberPairs("data/01_2024.txt")
	if err != nil {
		log.Fatalf("Error processing number pairs: %v", err)
	}
	totalDis := calculateTotalDistance(pairs)
	fmt.Printf("Day 1: %d\n", totalDis)
}
