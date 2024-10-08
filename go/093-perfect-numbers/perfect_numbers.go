package perfect

import (
	"errors"
	"math"
)

type Classification int

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

var ErrOnlyPositive = errors.New("input must be a positive integer")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}

	var sum int64
	sqrtN := int64(math.Ceil(math.Sqrt(float64(n))))
	for i := int64(1); i < sqrtN; i++ {
		if n%i == 0 {
			sum += i
			if n != n/i {
				sum += n / i
			}
		}
	}

	switch {
	case sum == n:
		return ClassificationPerfect, nil
	case sum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}
