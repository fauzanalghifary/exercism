package series

func All(n int, s string) []string {
	var result []string

	for i := range s {
		if i+n <= len(s) {
			result = append(result, s[i:i+n])
		}
	}

	return result
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
