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

	stringMatrix := string(data)
	fmt.Println(countVertical(stringMatrix))

}

func countHorizontal(strMat string) int {
	count := 0
	lines := strings.Split(strMat, "\n")

	for _, l := range lines {
		count += countSub(l, "XMAS")
		count += countSub(reversed(l), "XMAS")
	}

	return count
}
func countVertical(strMat string) int {
	count := 0

	rows := strings.Split(strMat, "\n")

	columns := []string{}
	colBytes := make([][]byte, len(rows[0]))

	for _, r := range rows {
		for i := range len(r) {
			colBytes[i] = append(colBytes[i], r[i])
		}
	}

	for _, colB := range colBytes {
		columns = append(columns, string(colB))
	}

	for _, c := range columns {
		count += countSub(c, "XMAS")
		count += countSub(reversed(c), "XMAS")
	}

	return count
}

func countDiagonal(strMat string) {

}

func countSub(str, sub string) int {
	sum := 0
	for i := range len(str) - len(sub) + 1 {
		if str[i:i+len(sub)] == sub {
			sum++
		}
	}
	return sum
}

func reversed(str string) string {
	new := []byte{}
	for i := len(str) - 1; i >= 0; i-- {
		new = append(new, str[i])
	}
	return string(new)
}
