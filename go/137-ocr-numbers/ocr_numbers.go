package ocr

import (
	"strings"
)

// recognizeDigit recognizes a single 3x4 digit
func recognizeDigit(lines []string) string {
	// Each digit is 3 characters wide
	if len(lines) != 4 {
		return "?"
	}

	// Ensure each line has at least 3 characters
	for i := range lines {
		if len(lines[i]) < 3 {
			// Pad with spaces if needed
			lines[i] += strings.Repeat(" ", 3-len(lines[i]))
		}
	}

	// Create digit signature from first 3 characters of each line
	signature := ""
	for i := 0; i < 3; i++ {
		signature += lines[i][:3]
	}

	// Map of digit signatures to their values
	digits := map[string]string{
		" _ | ||_|": "0",
		"     |  |": "1",
		" _  _||_ ": "2",
		" _  _| _|": "3",
		"   |_|  |": "4",
		" _ |_  _|": "5",
		" _ |_ |_|": "6",
		" _   |  |": "7",
		" _ |_||_|": "8",
		" _ |_| _|": "9",
	}

	if digit, ok := digits[signature]; ok {
		return digit
	}
	return "?"
}

func Recognize(input string) []string {
	// Split input into lines
	lines := strings.Split(input, "\n")

	// Remove first empty line if exists (from raw string literal)
	if len(lines) > 0 && lines[0] == "" {
		lines = lines[1:]
	}

	// Check if input is valid (must have height divisible by 4)
	if len(lines)%4 != 0 {
		return []string{"?"}
	}

	var results []string

	// Process each group of 4 lines (one number per group)
	for i := 0; i < len(lines); i += 4 {
		group := lines[i : i+4]

		// Find the width (max length of any line in this group)
		width := 0
		for _, line := range group {
			if len(line) > width {
				width = len(line)
			}
		}

		// Pad all lines to same width
		for j := range group {
			if len(group[j]) < width {
				group[j] += strings.Repeat(" ", width-len(group[j]))
			}
		}

		// Process each digit (3 characters wide)
		numDigits := (width + 2) / 3 // Round up
		digitStr := ""

		for d := 0; d < numDigits; d++ {
			start := d * 3
			end := start + 3

			// Extract the 4 lines for this digit
			digitLines := make([]string, 4)
			for j := 0; j < 4; j++ {
				if end <= len(group[j]) {
					digitLines[j] = group[j][start:end]
				} else if start < len(group[j]) {
					digitLines[j] = group[j][start:]
					// Pad to 3 characters
					digitLines[j] += strings.Repeat(" ", 3-len(digitLines[j]))
				} else {
					digitLines[j] = "   "
				}
			}

			digitStr += recognizeDigit(digitLines)
		}

		results = append(results, digitStr)
	}

	return results
}
