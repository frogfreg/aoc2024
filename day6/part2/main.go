package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type position struct {
	row, col int
}
type direction struct {
	rowDirection, colDirection int
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

	possiblePositions := map[position]bool{}
	walkPath(mat, initialPosition, possiblePositions)

	count := 0

	for pos := range possiblePositions {
		mat[pos.row][pos.col] = "#"
		if hasLoop(mat, initialPosition) {
			count++
		}
		mat[pos.row][pos.col] = "."
	}

	fmt.Println(count)

}

func hasLoop(mat [][]string, pos position) bool {
	rowIncrement := -1
	colIncrement := 0
	loopDict := map[position][]direction{}

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
			if dirList, exists := loopDict[pos]; exists && slices.ContainsFunc(dirList, func(dir direction) bool {
				return dir.rowDirection == rowIncrement && dir.colDirection == colIncrement
			}) {
				return true
			}
			loopDict[pos] = append(loopDict[pos], direction{rowDirection: rowIncrement, colDirection: colIncrement})

			newRI, newCI, err := updatedIncrements(rowIncrement, colIncrement)
			if err != nil {
				panic(err)
			}
			rowIncrement, colIncrement = newRI, newCI
		} else {
			pos = position{row: newPosRow, col: newPosCol}
		}
	}

	return false
}

func walkPath(mat [][]string, pos position, possiblePositions map[position]bool) {
	rowIncrement := -1
	colIncrement := 0

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
			possiblePositions[pos] = true
		}
	}
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
