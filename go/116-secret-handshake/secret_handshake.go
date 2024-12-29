package secret

import (
	"fmt"
	"slices"
)

func Handshake(code uint) []string {
	num := code
	var binary string
	for num > 0 {
		remainder := num % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		num = num / 2
	}

	// binary := fmt.Sprintf("%b", code)

	var result []string
	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == '1' {
			switch len(binary) - 1 - i {
			case 0:
				result = append(result, "wink")
			case 1:
				result = append(result, "double blink")
			case 2:
				result = append(result, "close your eyes")
			case 3:
				result = append(result, "jump")
			case 4:
				slices.Reverse(result)
				return result
			}
		}
	}

	return result
}
