package grep

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	var results []string
	isMultipleFiles := len(files) > 1

	isPrintLineNumber := slices.Contains(flags, "-n")
	isCaseInsensitive := slices.Contains(flags, "-i")
	isPrintFileName := slices.Contains(flags, "-l")
	isMatchEntireLine := slices.Contains(flags, "-x")
	isInverseMatch := slices.Contains(flags, "-v")

	re := pattern
	if isCaseInsensitive {
		re = "(?i)" + pattern
	}
	if isMatchEntireLine {
		re = "^" + re + "$"
	}
	regexPattern := regexp.MustCompile(re)

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		for i, sentence := range strings.Split(string(content), "\n") {
			tempSentence := sentence

			if isPrintLineNumber {
				tempSentence = fmt.Sprintf("%d:%s", i+1, tempSentence)
			}
			if isMultipleFiles {
				tempSentence = file + ":" + tempSentence
			}

			if isPrintFileName {
				if regexPattern.MatchString(sentence) {
					results = append(results, file)
					break
				}
			}

			if isInverseMatch {
				if !regexPattern.MatchString(sentence) && sentence != "" {
					results = append(results, tempSentence)
				}
				continue
			}

			if regexPattern.MatchString(sentence) {
				results = append(results, tempSentence)
			}
		}
	}

	return results
}
