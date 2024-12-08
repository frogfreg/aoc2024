package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	mat := [][]string{}

	for _, l := range strings.Split(string(data), "\n") {
		row := strings.Split(l, "")
		if len(row) == 0 {
			continue
		}
		mat = append(mat, row)
	}

	count := 0

	for i := range len(mat) {
		for j := range len(mat[0]) {
			if mat[i][j] == "A" {
				count += checkAndCount(mat, i, j)
			}
		}
	}

	fmt.Println(count)
}

func checkAndCount(mat [][]string, i, j int) int {
	rowLength := len(mat[0])
	colLength := len(mat)

	diagOne := false
	diagTwo := false
	if j > 0 && j < rowLength-1 && i > 0 && i < colLength-1 {
		if mat[i-1][j-1] == "M" && mat[i+1][j+1] == "S" {
			diagOne = true
		}
		if mat[i-1][j-1] == "S" && mat[i+1][j+1] == "M" {
			diagOne = true
		}
		if mat[i-1][j+1] == "M" && mat[i+1][j-1] == "S" {
			diagTwo = true
		}
		if mat[i-1][j+1] == "S" && mat[i+1][j-1] == "M" {
			diagTwo = true
		}
	}

	if diagOne && diagTwo {
		return 1
	}

	return 0
}
