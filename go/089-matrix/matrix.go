package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	matrix := Matrix{}
	lines := strings.Split(s, "\n")

	for i := 0; i < len(lines); i++ {
		currentRow := strings.Split(strings.Trim(lines[i], " "), " ")
		if i+1 < len(lines) {
			nextRow := strings.Split(strings.Trim(lines[i+1], " "), " ")
			if len(currentRow) != len(nextRow) {
				return nil, errors.New("uneven row")
			}
		}

		rowValues := make([]int, len(currentRow))
		for j := 0; j < len(currentRow); j++ {
			num, err := strconv.Atoi(currentRow[j])
			if err != nil {
				return nil, err
			}
			rowValues[j] = num
		}
		matrix = append(matrix, rowValues)
	}

	if len(lines) == 1 {
		cell := strings.Split(strings.Trim(lines[0], " "), " ")
		_, err := strconv.Atoi(lines[0])
		if err != nil && len(cell) == 1 {
			fmt.Println(err.Error())
			return nil, errors.New("invalid syntax")
		}
	}

	return matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	if len(m) == 0 {
		return nil
	}

	cols := make([][]int, len(m[0]))
	for i := range cols {
		cols[i] = make([]int, len(m))
		for j := range m {
			cols[i][j] = m[j][i]
		}
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	if len(m) == 0 {
		return nil
	}

	rows := make([][]int, len(m))
	for i := range m {
		rows[i] = make([]int, len(m[i]))
		copy(rows[i], m[i])
	}
	return rows
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[row]) {
		return false
	}

	m[row][col] = val
	return true
}
