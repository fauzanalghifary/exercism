package forth

import (
	"errors"
	"strconv"
	"strings"
)

// Forth evaluates Forth commands and returns the resulting stack
func Forth(input []string) ([]int, error) {
	stack := []int{}
	definitions := make(map[string][]string)

	for _, line := range input {
		words := strings.Fields(strings.ToLower(line))
		if len(words) == 0 {
			continue
		}

		// Check if this is a definition
		if words[0] == ":" {
			if err := defineWord(words, definitions); err != nil {
				return nil, err
			}
			continue
		}

		// Process words
		var err error
		stack, err = processWords(words, stack, definitions)
		if err != nil {
			return nil, err
		}
	}

	return stack, nil
}

// defineWord defines a new word
func defineWord(words []string, definitions map[string][]string) error {
	if len(words) < 3 {
		return errors.New("invalid definition")
	}

	if words[len(words)-1] != ";" {
		return errors.New("invalid definition")
	}

	name := words[1]

	// Check if trying to redefine a number
	if _, err := strconv.Atoi(name); err == nil {
		return errors.New("illegal operation")
	}

	// Get definition (everything between name and ;)
	definition := words[2 : len(words)-1]

	// Expand any custom words in the definition
	expanded := []string{}
	for _, word := range definition {
		if def, exists := definitions[word]; exists {
			expanded = append(expanded, def...)
		} else {
			expanded = append(expanded, word)
		}
	}

	definitions[name] = expanded
	return nil
}

// processWords processes a list of words
func processWords(words []string, stack []int, definitions map[string][]string) ([]int, error) {
	for _, word := range words {
		var err error

		// Check if it's a custom word
		if def, exists := definitions[word]; exists {
			stack, err = processWords(def, stack, definitions)
			if err != nil {
				return nil, err
			}
			continue
		}

		// Check if it's a number
		if num, err := strconv.Atoi(word); err == nil {
			stack = append(stack, num)
			continue
		}

		// Process built-in operations
		stack, err = processOperation(word, stack)
		if err != nil {
			return nil, err
		}
	}
	return stack, nil
}

// processOperation processes a single operation
func processOperation(op string, stack []int) ([]int, error) {
	switch op {
	case "+":
		if len(stack) < 2 {
			return nil, errors.New("stack underflow")
		}
		a, b := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		stack = append(stack, a+b)

	case "-":
		if len(stack) < 2 {
			return nil, errors.New("stack underflow")
		}
		a, b := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		stack = append(stack, a-b)

	case "*":
		if len(stack) < 2 {
			return nil, errors.New("stack underflow")
		}
		a, b := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		stack = append(stack, a*b)

	case "/":
		if len(stack) < 2 {
			return nil, errors.New("stack underflow")
		}
		a, b := stack[len(stack)-2], stack[len(stack)-1]
		if b == 0 {
			return nil, errors.New("divide by zero")
		}
		stack = stack[:len(stack)-2]
		stack = append(stack, a/b)

	case "dup":
		if len(stack) < 1 {
			return nil, errors.New("stack underflow")
		}
		stack = append(stack, stack[len(stack)-1])

	case "drop":
		if len(stack) < 1 {
			return nil, errors.New("stack underflow")
		}
		stack = stack[:len(stack)-1]

	case "swap":
		if len(stack) < 2 {
			return nil, errors.New("stack underflow")
		}
		stack[len(stack)-1], stack[len(stack)-2] = stack[len(stack)-2], stack[len(stack)-1]

	case "over":
		if len(stack) < 2 {
			return nil, errors.New("stack underflow")
		}
		stack = append(stack, stack[len(stack)-2])

	default:
		return nil, errors.New("undefined operation")
	}

	return stack, nil
}
