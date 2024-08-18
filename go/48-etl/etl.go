package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int)
	for point, chars := range in {
		for _, char := range chars {
			result[strings.ToLower(char)] = point
		}
	}

	return result
}
