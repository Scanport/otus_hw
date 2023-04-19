package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder
	var isEscaped bool
	r := []rune(s)
	for i := range r {
		if isEscaped {
			isEscaped = false
			continue
		}
		err := isValidFirstCharacter(i, r[i])
		if err != nil {
			return "", err
		}
		err = isValidNumber(i, r)
		if err != nil {
			return "", err
		}
		if string(r[i]) == "\\" {
			if isDigit(r[i+1]) || string(r[i+1]) == "\\" {
				result.WriteRune(r[i+1])
				isEscaped = true
				continue
			}
			return "", ErrInvalidString
		}
		count := getCountCurr(i, r)
		if count > 0 {
			result.WriteString(strings.Repeat(string(r[i-1]), count-1))
			continue
		}
		count = getCountNext(i, r)
		if count == 0 {
			continue
		}
		result.WriteRune(r[i])
	}

	return result.String(), nil
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isValidFirstCharacter(i int, r rune) error {
	if i == 0 && isDigit(r) {
		return ErrInvalidString
	}
	return nil
}

func isValidNumber(i int, r []rune) error {
	if i != len(r)-1 {
		if isDigit(r[i]) && isDigit(r[i+1]) {
			return ErrInvalidString
		}
	}
	return nil
}

func getCountCurr(i int, r []rune) int {
	if isDigit(r[i]) {
		count := int(r[i] - '0')
		if count == 0 {
			count++
		}
		return count
	}
	return -1
}

func getCountNext(i int, r []rune) int {
	if i != len(r)-1 && isDigit(r[i+1]) {
		count := int(r[i+1] - '0')
		return count
	}
	return -1
}
