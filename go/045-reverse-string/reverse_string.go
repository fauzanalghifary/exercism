package reverse

func Reverse(input string) string {
	reverseString := ""
	for _, char := range input {
		reverseString = string(char) + reverseString
	}
	return reverseString
}
