package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type doRegion struct {
	start  int
	finish int
}

func main() {
	data, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	memory := string(data)

	rgx := regexp.MustCompile(`mul\((\d?\d?\d?,\d?\d?\d?)\)`)

	matches := rgx.FindAllStringSubmatchIndex(memory, -1)
	sum := 0

	regions := findDoRegions(memory)

	for _, m := range matches {
		if isInRegion(regions, m[0]) {
			digits := strings.Split(memory[m[2]:m[3]], ",")

			leftNum, err := strconv.Atoi(digits[0])
			if err != nil {
				panic(err)
			}
			rightNum, err := strconv.Atoi(digits[1])
			if err != nil {
				panic(err)
			}

			sum += leftNum * rightNum
		}
	}

	fmt.Println(sum)

}

func findDoRegions(memory string) []doRegion {
	shouldDo := true
	left := 0
	right := -1

	regions := []doRegion{}

	for i := range len(memory) - 6 {
		if shouldDo && memory[i:i+7] == "don't()" {
			shouldDo = false
			right = i
			regions = append(regions, doRegion{left, right})
		}
		if !shouldDo && memory[i:i+4] == "do()" {
			shouldDo = true
			left = i
			right = -1
		}
	}

	if right == -1 {
		regions = append(regions, doRegion{left, len(memory)})
	}

	return regions
}

func isInRegion(list []doRegion, num int) bool {
	for _, r := range list {
		if num >= r.start && num <= r.finish {
			return true
		}
	}
	return false

}
