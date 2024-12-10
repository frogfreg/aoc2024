package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct {
	row, col int
}

func main() {
	data, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(data), "\n")
	mat := [][]string{}
	initialPosition := position{}

	for rowIndex, row := range rows {
		if strings.TrimSpace(row) == "" {
			continue
		}
		mat = append(mat, strings.Split(row, ""))

		colIndex := strings.Index(row, "^")

		if colIndex != -1 {
			initialPosition = position{col: colIndex, row: rowIndex}
		}
	}

	fmt.Println(analyzePath(mat, initialPosition))

}

func analyzePath(mat [][]string, pos position) int {
	rowIncrement := -1
	colIncrement := 0

	visited := map[position]bool{pos: true}

	for {
		newPosRow := pos.row + (rowIncrement * 1)
		newPosCol := pos.col + (colIncrement * 1)

		if newPosRow < 0 || newPosRow > (len(mat)-1) {
			break
		}
		if newPosCol < 0 || newPosCol > (len(mat[0])-1) {
			break
		}

		if mat[newPosRow][newPosCol] == "#" {
			newRI, newCI, err := updatedIncrements(rowIncrement, colIncrement)
			if err != nil {
				panic(err)
			}
			rowIncrement, colIncrement = newRI, newCI
		} else {
			pos = position{row: newPosRow, col: newPosCol}
			visited[pos] = true
		}
	}

	return len(visited)
}

func updatedIncrements(rowIncrement, colIncrement int) (int, int, error) {
	if rowIncrement == -1 {
		return 0, 1, nil
	}
	if colIncrement == 1 {
		return 1, 0, nil
	}
	if rowIncrement == 1 {
		return 0, -1, nil
	}
	if colIncrement == -1 {
		return -1, 0, nil
	}
	return 0, 0, fmt.Errorf("invalid state")
}
