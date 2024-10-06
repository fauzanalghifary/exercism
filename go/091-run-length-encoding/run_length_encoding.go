package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	var output strings.Builder
	var lastChar rune
	charCount := make(map[rune]int)

	for i, char := range input {
		charCount[char]++
		if i > 0 && lastChar != char {
			WriteToOutput(charCount, lastChar, &output)
		}

		lastChar = char

		if i == len(input)-1 {
			WriteToOutput(charCount, char, &output)
		}
	}

	return output.String()
}

func RunLengthDecode(input string) string {
	var output strings.Builder
	var stack strings.Builder

	for _, char := range input {
		if unicode.IsDigit(char) {
			stack.WriteRune(char)
		}

		if !unicode.IsDigit(char) {
			count, _ := strconv.Atoi(stack.String())
			if count == 0 {
				count = 1
			}
			for i := 0; i < count; i++ {
				output.WriteRune(char)
			}

			stack.Reset()
		}
	}

	return output.String()
}

func WriteToOutput(charCount map[rune]int, char rune, output *strings.Builder) {
	count := charCount[char]
	if count > 1 {
		output.WriteString(strconv.Itoa(count))
	}
	output.WriteRune(char)
	charCount[char] = 0
}
