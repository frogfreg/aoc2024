package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("../test-input")
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

	count := 0
	rowLength := len(mat[0])
	colLength := len(mat)

	// right+down
	if j < rowLength-3 && i < colLength-3 && joinDiagonal(mat, i, j, 1, 1) == "XMAS" {
		count++
	}
	// right+up
	if j < rowLength-3 && i > 2 && joinDiagonal(mat, i, j, -1, 1) == "XMAS" {
		count++
	}
	// left+down
	if j > 2 && i < colLength-3 && joinDiagonal(mat, i, j, 1, -1) == "XMAS" {
		count++
	}
	// left+up
	if j > 2 && i > 2 && joinDiagonal(mat, i, j, -1, -1) == "XMAS" {
		count++
	}
	return count

}

func reversed(str string) string {
	final := []byte{}
	for i := len(str) - 1; i >= 0; i-- {
		final = append(final, str[i])

	}
	return string(final)
}

func joinCol(mat [][]string, colIndex, start, finish int) string {
	final := ""
	for i := start; i < finish; i++ {
		final += mat[i][colIndex]
	}
	return final
}
func joinDiagonal(mat [][]string, rowIndex, colIndex, rowDirection, colDirection int) string {
	final := ""
	for i := range 4 {
		final += mat[rowIndex+(rowDirection*i)][colIndex+(colDirection*i)]
	}

	return final
}