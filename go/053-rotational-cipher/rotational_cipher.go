package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	output := ""

	for _, char := range plain {
		if unicode.IsLetter(char) {
			shiftedUnicode := char + rune(shiftKey)
			if unicode.IsUpper(char) && shiftedUnicode > 'Z' ||
				unicode.IsLower(char) && shiftedUnicode > 'z' {
				shiftedUnicode -= 26
			}
			output += string(shiftedUnicode)
		} else {
			output += string(char)
		}
	}

	return output
}
