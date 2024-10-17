package foodchain

import (
	"fmt"
	"strings"
)

const wriggle = " wriggled and jiggled and tickled inside her"

var verse = []struct{ eaten, comment string }{
	{"", ""},
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", "It" + wriggle + ".\n"},
	{"bird", "How absurd to swallow a bird!\n"},
	{"cat", "Imagine that, to swallow a cat!\n"},
	{"dog", "What a hog, to swallow a dog!\n"},
	{"goat", "Just opened her throat and swallowed a goat!\n"},
	{"cow", "I don't know how she swallowed a cow!\n"},
	{"horse", "She's dead, of course!"},
}

func Verse(v int) string {
	if v < 1 || v > 8 {
		return ""
	}

	output := fmt.Sprintf(
		"I know an old lady who swallowed a %s.\n%s",
		verse[v].eaten,
		verse[v].comment,
	)

	if v == 1 || v == 8 {
		return output
	}

	for i := v; i > 1; i-- {
		output += fmt.Sprintf(
			"She swallowed the %s to catch the %s",
			verse[i].eaten,
			verse[i-1].eaten,
		)

		if i == 3 {
			output += " that" + wriggle
		}

		output += ".\n"
	}

	return output + verse[1].comment
}

func Verses(start, end int) string {
	var output []string
	for i := start; i <= end; i++ {
		output = append(output, Verse(i))
	}
	return strings.Join(output, "\n\n")
}

func Song() string {
	return Verses(1, 8)
}
