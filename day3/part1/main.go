package main

import (
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("../test-input")
	if err != nil {
		panic(err)
	}

	memory := string(data)
	sum := 0

	left := 0
	right := 0

	left = strings.Index(memory, "mul(")
	right = strings.Index(memory[left+4:], ")")

}
