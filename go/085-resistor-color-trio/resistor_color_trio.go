package resistorcolortrio

import (
	"fmt"
	"math"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	COLOR_CODE := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}

	val := (COLOR_CODE[colors[0]]*10 + COLOR_CODE[colors[1]]) * int(
		math.Pow(
			10,
			float64(COLOR_CODE[colors[2]]),
		),
	)

	if val >= 1000000000 {
		return fmt.Sprintf("%d gigaohms", val/1000000000)
	}

	if val >= 1000000 {
		return fmt.Sprintf("%d megaohms", val/1000000)
	}

	if val >= 1000 {
		return fmt.Sprintf("%d kiloohms", val/1000)
	}

	return fmt.Sprintf("%d ohms", val)
}
