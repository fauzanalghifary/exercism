package house

import (
	"fmt"
	"strings"
)

var verse = []struct{ animal, action string }{
	{"house that Jack built", "lay in"},
	{"malt", "ate"},
	{"rat", "killed"},
	{"cat", "worried"},
	{"dog", "tossed"},
	{"cow with the crumpled horn", "milked"},
	{"maiden all forlorn", "kissed"},
	{"man all tattered and torn", "married"},
	{"priest all shaven and shorn", "woke"},
	{"rooster that crowed in the morn", "kept"},
	{"farmer sowing his corn", "belonged to"},
	{"horse and the hound and the horn", ""},
}

func Verse(v int) string {
	var output string

	for i := v; i >= 1; i-- {
		if i == v {
			output += fmt.Sprintf("This is the %s", verse[v-1].animal)
		} else {
			output += fmt.Sprintf("\nthat %s the %s", verse[i-1].action, verse[i-1].animal)
		}
	}

	return output + "."
}

func Song() string {
	var verses []string

	for i := 1; i < 13; i++ {
		verses = append(verses, Verse(i))
	}

	return strings.Join(verses, "\n\n")
}
