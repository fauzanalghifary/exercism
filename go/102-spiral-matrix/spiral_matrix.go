package spiralmatrix

func SpiralMatrix(size int) [][]int {
	arr := make([][]int, size)
	for i := 0; i < size; i++ {
		arr[i] = make([]int, size)
	}

	directions := [][]int{
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
		{-1, 0}, // Up
	}

	row, col, dir := 0, 0, 0
	for i := 1; i <= size*size; i++ {
		arr[row][col] = i

		nextRow := row + directions[dir][0]
		nextCol := col + directions[dir][1]

		if nextRow < 0 || nextRow >= size || nextCol < 0 || nextCol >= size || arr[nextRow][nextCol] != 0 {
			dir = (dir + 1) % 4
			nextRow = row + directions[dir][0]
			nextCol = col + directions[dir][1]
		}

		row, col = nextRow, nextCol
	}

	return arr
}
