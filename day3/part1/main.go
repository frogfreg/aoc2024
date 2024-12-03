package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	memory := string(data)

	rgx := regexp.MustCompile(`mul\((\d?\d?\d?,\d?\d?\d?)\)`)

	matches := rgx.FindAllStringSubmatch(memory, -1)
	sum := 0

	for _, m := range matches {
		digits := strings.Split(m[1], ",")

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

	fmt.Println(sum)
}
