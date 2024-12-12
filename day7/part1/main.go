package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("../test-input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}

	}

}

func checkOperatorTest(test string) int {
	parts := strings.Split(test, ":")

	objective, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	numStringList := strings.Fields(parts[1])

	nums := []int{}

	for _, numString := range numStringList {
		num, err := strconv.Atoi(numString)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
}
