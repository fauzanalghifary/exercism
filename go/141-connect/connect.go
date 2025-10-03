package connect

func ResultOf(lines []string) (string, error) {
	if len(lines) == 0 {
		return "", nil
	}

	// Check if X wins (connects left to right)
	if hasWon(lines, 'X') {
		return "X", nil
	}

	// Check if O wins (connects top to bottom)
	if hasWon(lines, 'O') {
		return "O", nil
	}

	return "", nil
}

func hasWon(board []string, player byte) bool {
	if len(board) == 0 {
		return false
	}

	height := len(board)
	width := len(board[0])

	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	if player == 'X' {
		// X needs to connect left to right
		// Start from all positions in the leftmost column
		for row := 0; row < height; row++ {
			if row < len(board) && 0 < len(board[row]) && board[row][0] == player {
				if dfs(board, row, 0, player, visited, width) {
					return true
				}
			}
		}
	} else {
		// O needs to connect top to bottom
		// Start from all positions in the top row
		for col := 0; col < width; col++ {
			if 0 < len(board) && col < len(board[0]) && board[0][col] == player {
				if dfs(board, 0, col, player, visited, width) {
					return true
				}
			}
		}
	}

	return false
}

func dfs(board []string, row, col int, player byte, visited [][]bool, targetWidth int) bool {
	height := len(board)

	// Check bounds
	if row < 0 || row >= height || col < 0 || col >= len(board[row]) {
		return false
	}

	// Check if already visited or not the player's stone
	if visited[row][col] || board[row][col] != player {
		return false
	}

	// Mark as visited
	visited[row][col] = true

	// Check win condition
	if player == 'X' {
		// X wins if reaching the rightmost column
		if col == len(board[row])-1 {
			return true
		}
	} else {
		// O wins if reaching the bottom row
		if row == height-1 {
			return true
		}
	}

	// Explore neighbors (6 directions in hex grid)
	// In this representation, the hex grid is laid out where each row has the same positions
	// but they are visually offset. The connections are:
	// - Same row: left (-1), right (+1)
	// - Previous row: same col, col-1
	// - Next row: same col, col+1
	directions := []struct{ dr, dc int }{
		{0, -1}, // left
		{0, 1},  // right
		{-1, 0}, // up-left (previous row, same column)
		{-1, 1}, // up-right (previous row, right column)
		{1, -1}, // down-left (next row, left column)
		{1, 0},  // down-right (next row, same column)
	}

	for _, dir := range directions {
		newRow := row + dir.dr
		newCol := col + dir.dc
		if dfs(board, newRow, newCol, player, visited, targetWidth) {
			return true
		}
	}

	return false
}
