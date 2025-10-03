package stateoftictactoe

import "errors"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	// Count X and O marks
	xCount := 0
	oCount := 0
	for _, row := range board {
		for _, cell := range row {
			if cell == 'X' {
				xCount++
			} else if cell == 'O' {
				oCount++
			}
		}
	}

	// Validate turn order: X starts, so X should have the same or 1 more than O
	if xCount < oCount || xCount > oCount+1 {
		return "", errors.New("invalid turn order")
	}

	// Check for winners
	xWins := hasWon(board, 'X')
	oWins := hasWon(board, 'O')

	// Invalid: both players won
	if xWins && oWins {
		return "", errors.New("both players won")
	}

	// Invalid: O won but X has one more move (continued playing after O won)
	if oWins && xCount > oCount {
		return "", errors.New("continued playing after win")
	}

	// Invalid: X won, but O continued playing
	if xWins && xCount == oCount {
		return "", errors.New("continued playing after win")
	}

	// Check win state
	if xWins || oWins {
		return Win, nil
	}

	// Check draw: board is full (9 cells)
	if xCount+oCount == 9 {
		return Draw, nil
	}

	// Game is ongoing
	return Ongoing, nil
}

func hasWon(board []string, player byte) bool {
	// Check rows
	for _, row := range board {
		if len(row) == 3 && row[0] == player && row[1] == player && row[2] == player {
			return true
		}
	}

	// Check columns
	for col := 0; col < 3; col++ {
		if board[0][col] == player && board[1][col] == player && board[2][col] == player {
			return true
		}
	}

	// Check diagonals
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}
