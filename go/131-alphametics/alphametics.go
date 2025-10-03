package alphametics

import (
	"errors"
	"regexp"
	"strings"
)

func Solve(puzzle string) (map[string]int, error) {
	// Parse the puzzle
	parts := strings.Split(puzzle, "==")
	if len(parts) != 2 {
		return nil, errors.New("invalid puzzle format")
	}

	leftSide := strings.TrimSpace(parts[0])
	result := strings.TrimSpace(parts[1])

	// Extract words (addends and result)
	re := regexp.MustCompile(`[A-Z]+`)
	leftWords := re.FindAllString(leftSide, -1)

	// Get unique letters
	letters := make(map[rune]bool)
	for _, word := range leftWords {
		for _, ch := range word {
			letters[ch] = true
		}
	}
	for _, ch := range result {
		letters[ch] = true
	}

	// Check if too many unique letters (max 10 digits)
	if len(letters) > 10 {
		return nil, errors.New("too many unique letters")
	}

	// Get leading letters (cannot be zero)
	leadingLetters := make(map[rune]bool)
	for _, word := range leftWords {
		if len(word) > 1 {
			leadingLetters[rune(word[0])] = true
		}
	}
	if len(result) > 1 {
		leadingLetters[rune(result[0])] = true
	}

	// Convert map to slice for iteration
	uniqueLetters := make([]rune, 0, len(letters))
	for letter := range letters {
		uniqueLetters = append(uniqueLetters, letter)
	}

	// Try all permutations using backtracking
	assignment := make(map[rune]int)
	usedDigits := make(map[int]bool)

	if backtrack(uniqueLetters, 0, assignment, usedDigits, leadingLetters, leftWords, result) {
		// Convert rune map to string map
		solution := make(map[string]int)
		for letter, digit := range assignment {
			solution[string(letter)] = digit
		}
		return solution, nil
	}

	return nil, errors.New("no solution found")
}

func backtrack(letters []rune, index int, assignment map[rune]int, usedDigits map[int]bool,
	leadingLetters map[rune]bool, leftWords []string, result string) bool {

	if index == len(letters) {
		// Check if current assignment is valid
		return isValidSolution(assignment, leftWords, result)
	}

	letter := letters[index]

	for digit := 0; digit <= 9; digit++ {
		// Skip if digit is already used
		if usedDigits[digit] {
			continue
		}

		// Skip if leading letter and digit is 0
		if digit == 0 && leadingLetters[letter] {
			continue
		}

		// Try this assignment
		assignment[letter] = digit
		usedDigits[digit] = true

		if backtrack(letters, index+1, assignment, usedDigits, leadingLetters, leftWords, result) {
			return true
		}

		// Backtrack
		delete(assignment, letter)
		delete(usedDigits, digit)
	}

	return false
}

func isValidSolution(assignment map[rune]int, leftWords []string, result string) bool {
	sum := 0

	// Calculate sum of left side
	for _, word := range leftWords {
		value := wordToNumber(word, assignment)
		sum += value
	}

	// Calculate result value
	resultValue := wordToNumber(result, assignment)

	return sum == resultValue
}

func wordToNumber(word string, assignment map[rune]int) int {
	value := 0
	for _, ch := range word {
		value = value*10 + assignment[ch]
	}
	return value
}
