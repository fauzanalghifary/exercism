package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fd FodderCalculator, cowsQty int) (float64, error) {
	totalFodders, err := fd.FodderAmount(cowsQty)
	if err != nil {
		return 0, err
	}

	fatFactor, err := fd.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return totalFodders * fatFactor / float64(cowsQty), nil
}

func ValidateInputAndDivideFood(fd FodderCalculator, cowsQty int) (float64, error) {
	if cowsQty > 0 {
		return DivideFood(fd, cowsQty)
	}

	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	message string
	cowsQty int
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cowsQty, e.message)
}

func ValidateNumberOfCows(cowsQty int) error {
	if cowsQty < 0 {
		return &InvalidCowsError{
			message: "there are no negative cows",
			cowsQty: cowsQty,
		}
	} else if cowsQty == 0 {
		return &InvalidCowsError{
			message: "no cows don't need food",
			cowsQty: cowsQty,
		}
	}

	return nil
}
