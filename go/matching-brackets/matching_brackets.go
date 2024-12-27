package brackets

func Bracket(input string) bool {
	bracketPairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []rune

	for _, char := range input {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

//func Bracket(input string) bool {
//	openBrackets := []string{
//		"[",
//		"{",
//		"(",
//	}
//
//	closeBrackets := []string{
//		"]",
//		"}",
//		")",
//	}
//
//	var temp []string
//
//	for _, char := range input {
//		stringChar := string(char)
//		if slices.Contains(openBrackets, stringChar) {
//			temp = append(temp, stringChar)
//		} else if slices.Contains(closeBrackets, stringChar) {
//			if len(temp) == 0 {
//				return false
//			}
//			lastChar := temp[len(temp)-1]
//			closeBracketIndex := slices.Index(closeBrackets, stringChar)
//			openBracketIndex := slices.Index(openBrackets, lastChar)
//			if closeBracketIndex == openBracketIndex {
//				temp = slices.Delete(temp, len(temp)-1, len(temp))
//			} else {
//				return false
//			}
//		}
//	}
//
//	if len(temp) == 0 {
//		return true
//	}
//
//	return false
//}
