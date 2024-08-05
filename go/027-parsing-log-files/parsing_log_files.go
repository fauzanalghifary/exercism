package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)]`)
	if err != nil {
		return false
	}

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`<[~*=-]*>`)
	if err != nil {
		return []string{text}
	}

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, err := regexp.Compile(`(?i)".*password.*"`)
	if err != nil {
		return 0
	}

	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`end-of-line\d+`)
	if err != nil {
		return text
	}

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, err := regexp.Compile(`User\s+(\S+)`)
	if err != nil {
		return lines
	}

	var result []string
	for _, line := range lines {
		substr := re.FindStringSubmatch(line)
		if len(substr) > 0 {
			newLine := fmt.Sprintf("[USR] %s %s", substr[1], line)
			result = append(result, newLine)
		} else {
			result = append(result, line)
		}
	}

	return result
}
