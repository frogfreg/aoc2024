package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}
	var leftList []int

	occurrences := map[int]int{}

	for _, line := range strings.Split(string(data), "\n") {
		bothListsLine := strings.Fields(line)
		if len(bothListsLine) < 2 {
			continue
		}

		left, right := bothListsLine[0], bothListsLine[1]

		leftNum, err := strconv.Atoi(left)
		if err != nil {
			panic(err)
		}
		rightNum, err := strconv.Atoi(right)
		if err != nil {
			panic(err)
		}

		leftList = append(leftList, leftNum)
		occurrences[rightNum] += 1

	}

	sum := 0

	for _, num := range leftList {
		sum += num * occurrences[num]
	}

	fmt.Println(sum)
}
