package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	sum := 0
	for _, cell := range cb[file] {
		if cell {
			sum += 1
		}
	}

	return sum
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	sum := 0
	for _, ranks := range cb {
		if ranks[rank-1] {
			sum++
		}
	}
	return sum
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	sum := 0
	for _, ranks := range cb {
		sum += len(ranks)
	}

	return sum
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	sum := 0
	for file := range cb {
		sum += CountInFile(cb, file)
	}
	return sum
}
