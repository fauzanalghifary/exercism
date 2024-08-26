package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var proverb []string

	if len(rhyme) == 0 {
		return proverb
	}

	for i := range rhyme {
		if i >= len(rhyme)-1 {
			break
		}

		proverb = append(
			proverb,
			fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]),
		)
	}

	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return proverb
}
