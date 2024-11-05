package matrix

import (
	"slices"
	"strconv"
	"strings"
)

type Matrix [][]int

type Pair [2]int

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	matrix := make(Matrix, len(rows))
	for i, row := range rows {
		cells := strings.Fields(row)
		matrix[i] = make([]int, len(cells))
		for j, cell := range cells {
			val, _ := strconv.Atoi(cell)
			matrix[i][j] = val
		}
	}
	return &matrix, nil
}

func (m *Matrix) Saddle() []Pair {
	var result []Pair

	if len(*m) == 1 {
		for _, row := range *m {
			if len(row) == 0 {
				return result
			}
		}
	}

	transposedMatrix := m.Transpose()

	for i, row := range *m {
		maxRow := slices.Max(row)
		for j, _ := range row {
			minColumn := slices.Min(transposedMatrix[j])
			if maxRow == minColumn {
				result = append(result, Pair{i + 1, j + 1})
			}
		}
	}

	return result
}

func (m *Matrix) Transpose() Matrix {
	transposed := make(Matrix, len((*m)[0]))
	for i := range transposed {
		transposed[i] = make([]int, len(*m))
		for j := range *m {
			transposed[i][j] = (*m)[j][i]
		}
	}
	return transposed
}
