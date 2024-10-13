package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("char out of range")
	}

	var output []string
	num := int(char - 'A' + 1)

	for i := 0; i < num; i++ {
		var tempArray []string
		for j := 0; j < num; j++ {
			if j == num-i-1 {
				tempArray = append(tempArray, string(rune('A'+i)))
			} else {
				tempArray = append(tempArray, " ")
			}
		}
		output = append(output, strings.Join(AddSymmetric(tempArray), ""))
	}

	return strings.Join(AddSymmetric(output), "\n"), nil
}

func AddSymmetric(input []string) []string {
	originalLength := len(input)
	for i := originalLength - 2; i >= 0; i-- {
		input = append(input, input[i])
	}
	return input
}
