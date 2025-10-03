package palindrome

import (
	"errors"
	"strconv"
)

// Product represents a palindrome product with its value and factorizations
type Product struct {
	Value          int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}

	palindromes := make(map[int][][2]int)

	// Find all palindrome products in the range
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			product := i * j
			if isPalindrome(product) {
				palindromes[product] = append(palindromes[product], [2]int{i, j})
			}
		}
	}

	if len(palindromes) == 0 {
		return Product{}, Product{}, errors.New("no palindromes")
	}

	// Find min and max palindromes
	var minVal, maxVal int
	first := true
	for val := range palindromes {
		if first {
			minVal = val
			maxVal = val
			first = false
		} else {
			if val < minVal {
				minVal = val
			}
			if val > maxVal {
				maxVal = val
			}
		}
	}

	return Product{minVal, palindromes[minVal]}, Product{maxVal, palindromes[maxVal]}, nil
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
