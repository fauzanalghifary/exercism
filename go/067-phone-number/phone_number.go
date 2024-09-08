package phonenumber

import (
	"errors"
	"fmt"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	var cleanNumber string

	for _, char := range phoneNumber {
		if unicode.IsDigit(char) {
			cleanNumber += string(char)
		}
	}

	if len(cleanNumber) < 10 || len(cleanNumber) > 11 {
		return "", errors.New("invalid length")
	}

	if len(cleanNumber) == 11 {
		if cleanNumber[0] != '1' {
			return "", errors.New("invalid number")
		} else {
			cleanNumber = cleanNumber[1:]
		}
	}

	if int(cleanNumber[0]-'0') < 2 {
		return "", errors.New("invalid area code")
	}

	if int(cleanNumber[3]-'0') < 2 {
		return "", errors.New("invalid exchange code")
	}

	return cleanNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	cleanNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return cleanNumber[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	cleanNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", cleanNumber[0:3], cleanNumber[3:6], cleanNumber[6:]), nil
}
