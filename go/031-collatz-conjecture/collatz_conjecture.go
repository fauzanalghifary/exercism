package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("input must be a positive integer")
	}
	return collatzSteps(n), nil
}

func collatzSteps(n int) int {
	if n == 1 {
		return 0
	} else if n%2 == 0 {
		return 1 + collatzSteps(n/2)
	} else {
		return 1 + collatzSteps(3*n+1)
	}
}
