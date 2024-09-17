package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

var validOperation = map[string]bool{
	"plus":       true,
	"minus":      true,
	"multiplied": true,
	"divided":    true,
}

func Answer(question string) (int, bool) {
	arrayQuestion := strings.Split(question, " ")
	isNumber := regexp.MustCompile(`\d`)
	var nums []int
	var operations []string
	shouldPush := "number"

	lastWord := strings.ReplaceAll(arrayQuestion[len(arrayQuestion)-1], "?", "")
	if !isNumber.MatchString(lastWord) {
		return 0, false
	}

	for _, word := range arrayQuestion {
		word = strings.ReplaceAll(word, "?", "")
		if isNumber.MatchString(word) {
			num, _ := strconv.Atoi(word)
			nums = append(nums, num)
			if shouldPush != "number" {
				return 0, false
			}
			shouldPush = "operation"
		} else if validOperation[word] {
			operations = append(operations, word)
			if shouldPush != "operation" {
				return 0, false
			}
			shouldPush = "number"
		}
	}

	sum := nums[0]
	nums = nums[1:]
	for _, op := range operations {
		num := nums[0]
		nums = nums[1:]
		switch op {
		case "plus":
			sum += num
		case "minus":
			sum -= num
		case "multiplied":
			sum *= num
		case "divided":
			sum /= num
		}
	}

	return sum, true
}

func Plus(a *int, b int) {
	*a = *a + b
}

func Minus(a *int, b int) {
	*a = *a - b
}

func Multiplied(a *int, b int) {
	*a = *a * b
}

func Divided(a *int, b int) {
	*a = *a / b
}
