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
	var leftList, rightList []int

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
		rightList = append(rightList, rightNum)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)
	sum := 0

	for i := range len(leftList) {
		sum += abs(leftList[i], rightList[i])
	}

	fmt.Println(sum)

}

func abs(a, b int) int {
	res := a - b
	if res > 0 {
		return res
	}
	return -res
}
