package minesweeper

var neighboursShift = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

// Annotate returns an annotated board
func Annotate(board []string) []string {
	boardLen := len(board)
	if boardLen == 0 {
		return []string{}
	}
	rowLen := len(board[0])
	newBoard := make([]string, boardLen)
	newRow := make([]rune, rowLen)
	for i, row := range board {
		copy(newRow, []rune(row))
		for x, c := range row {
			if c == '*' {
				continue
			} else {
				neighboursCount := 0
				for _, shift := range neighboursShift {
					sY := i + shift[0]
					sX := x + shift[1]
					if sY >= 0 && sY < boardLen && sX >= 0 && sX < rowLen {
						if board[sY][sX] == '*' {
							neighboursCount++
						}
					}
				}
				if neighboursCount != 0 {
					newRow[x] = rune(neighboursCount + 48)
				}
			}
		}
		newBoard[i] = string(newRow)
	}
	return newBoard
}
