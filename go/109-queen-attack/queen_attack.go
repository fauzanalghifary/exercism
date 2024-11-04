package queenattack

import (
	"errors"
	"math"
)

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if whitePosition == blackPosition || len(whitePosition) != 2 || len(blackPosition) != 2 {
		return false, errors.New("invalid input")
	}

	whiteX, whiteY := whitePosition[0], whitePosition[1]
	blackX, blackY := blackPosition[0], blackPosition[1]

	if whiteX < 'a' || whiteX > 'h' || blackX < 'a' || blackX > 'h' ||
		whiteY > '8' || blackY > '8' || whiteY < '1' || blackY < '1' {
		return false, errors.New("invalid input")
	}

	diffX := math.Abs(float64(int64(whiteX) - int64(blackX)))
	diffY := math.Abs(float64(int64(whiteY) - int64(blackY)))

	if whiteX == blackX || whiteY == blackY || diffX == diffY {
		return true, nil
	}

	return false, nil
}
