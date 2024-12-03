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

	memory := string(data)
	sum := 0

	left := 0
	right := 0

	for {
		left = strings.Index(memory, "mul(")
		right = strings.Index(memory, ")")

		if left == -1 || right == -1 {
			break
		}

		if right < left {
			memory = memory[left+4:]
			continue
		}

		fmt.Printf("memory: %v, left: %v, right: %v\n", memory, left, right)

		digits := strings.Split(memory[left+4:right], ",")
		if len(digits) != 2 {
			memory = memory[left+4:]
			continue
		}

		leftDigit, rightDigit := digits[0], digits[1]
		leftNum, err := strconv.Atoi(leftDigit)
		if err != nil {
			memory = memory[left+4:]
			continue
		}
		rightNum, err := strconv.Atoi(rightDigit)
		if err != nil {
			memory = memory[left+4:]
			continue
		}

		if leftNum > 999 || rightNum > 999 {
			memory = memory[left+4:]
			continue
		}

		fmt.Printf("digits: %#v, %v * %v= %v\n", digits, leftNum, rightNum, leftNum*rightNum)

		sum += leftNum * rightNum

		memory = memory[right+1:]
	}

	fmt.Println(sum)
}
