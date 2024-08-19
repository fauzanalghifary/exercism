package bob

import (
	"regexp"
	"strings"
)

func Hey(remark string) string {
	isQuestion := regexp.MustCompile(`\?\s*$`).MatchString(remark)
	isYelling := regexp.MustCompile(`[a-zA-Z]`).MatchString(remark) && strings.ToUpper(remark) == remark
	isNotSilence := regexp.MustCompile(`\S`).MatchString(remark)

	if isQuestion && isYelling {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion {
		return "Sure."
	} else if isYelling {
		return "Whoa, chill out!"
	} else if !isNotSilence {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}
