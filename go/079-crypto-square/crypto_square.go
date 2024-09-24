package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	if pt == "" {
		return ""
	}

	var cleanString string
	for _, char := range pt {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleanString += string(unicode.ToLower(char))
		}
	}

	sideLength := math.Sqrt(float64(len(cleanString)))
	rowNum := int(math.Floor(sideLength))
	colNum := int(math.Ceil(sideLength))
	if rowNum*colNum < len(cleanString) {
		rowNum++
	}

	grid := make([]string, colNum)
	for i := 0; i < len(cleanString); i++ {
		grid[i%colNum] += string(cleanString[i])
	}

	for i := 0; i < len(grid); i++ {
		for len(grid[i]) < rowNum {
			grid[i] += " "
		}
	}

	return strings.Join(grid, " ")
}
