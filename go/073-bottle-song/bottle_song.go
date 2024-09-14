package bottlesong

import (
	"fmt"
	"strings"
)

var NumWord = map[int]string{
	0:  "no",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
}

func Recite(startBottles, takeDown int) []string {
	var verses []string
	bottles := "bottles"
	bottlesNum := startBottles
	for i := 0; i < takeDown; i++ {
		if bottlesNum == 1 {
			bottles = "bottle"
		}
		verses = append(
			verses,
			fmt.Sprintf("%s green %s hanging on the wall,", NumWord[bottlesNum], bottles),
			fmt.Sprintf("%s green %s hanging on the wall,", NumWord[bottlesNum], bottles),
			"And if one green bottle should accidentally fall,",
		)

		bottlesNum--
		if bottlesNum == 1 {
			bottles = "bottle"
		} else if bottlesNum == 0 {
			bottles = "bottles"
		}
		verses = append(
			verses,
			fmt.Sprintf(
				"There'll be %s green %s hanging on the wall.",
				strings.ToLower(NumWord[bottlesNum]),
				bottles,
			),
		)

		if i < takeDown-1 {
			verses = append(verses, "")
		}
	}

	return verses
}
