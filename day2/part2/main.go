package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	reports := strings.Split(string(data), "\n")

	safeCount := 0

	for _, r := range reports {
		levels := strings.Fields(r)
		if len(levels) < 1 {
			continue
		}
		numLevels := []int{}

		for _, l := range levels {
			n, err := strconv.Atoi(l)
			if err != nil {
				panic(err)
			}
			numLevels = append(numLevels, n)
		}

		if anyValid(numLevels) {
			safeCount++
		}

	}
	fmt.Println(safeCount)
}

func anyValid(originalLevels []int) bool {

	if isSafe(originalLevels) {
		return true
	}

	for i := range len(originalLevels) {
		copy := slices.Clone(originalLevels)
		if isSafe(append(copy[:i], copy[i+1:]...)) {
			return true
		}
	}

	return false

}

func isSafe(levels []int) bool {

	decreasing := false
	if levels[0] > levels[1] {
		decreasing = true
	}

	for i := range len(levels) - 1 {
		difference := levels[i] - levels[i+1]
		if decreasing && (difference > 3 || difference < 1) {
			return false
		}
		if !decreasing && (-difference > 3 || -difference < 1) {
			return false
		}
	}

	return true
}
