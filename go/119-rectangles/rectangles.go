package rectangles

func Count(diagram []string) int {
	count := 0
	corners := findCorners(diagram)

	for i := 0; i < len(corners); i++ {
		for j := i + 1; j < len(corners); j++ {
			if isValidRectangle(diagram, corners[i], corners[j]) {
				count++
			}
		}
	}

	return count
}

func findCorners(diagram []string) [][2]int {
	var corners [][2]int
	for y, row := range diagram {
		for x, char := range row {
			if char == '+' {
				corners = append(corners, [2]int{x, y})
			}
		}
	}

	return corners
}

func isValidRectangle(diagram []string, topLeft, bottomRight [2]int) bool {
	if topLeft[0] >= bottomRight[0] || topLeft[1] >= bottomRight[1] {
		return false
	}

	for x := topLeft[0]; x <= bottomRight[0]; x++ {
		if diagram[topLeft[1]][x] != '+' && diagram[topLeft[1]][x] != '-' {
			return false
		}

		if diagram[bottomRight[1]][x] != '+' && diagram[bottomRight[1]][x] != '-' {
			return false
		}
	}

	for y := topLeft[1]; y <= bottomRight[1]; y++ {
		if diagram[y][topLeft[0]] != '+' && diagram[y][topLeft[0]] != '|' {
			return false
		}

		if diagram[y][bottomRight[0]] != '+' && diagram[y][bottomRight[0]] != '|' {
			return false
		}
	}

	return true
}
