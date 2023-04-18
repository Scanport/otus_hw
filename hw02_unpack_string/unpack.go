package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder
	var isEscaped bool
	for i := range s {
		if isEscaped {
			isEscaped = false
			continue
		}
		err := isValidFirstCharacter(i, s[i])
		if err != nil {
			return "", err
		}
		err = isValidNumber(i, s)
		if err != nil {
			return "", err
		}
		if string(s[i]) == "\\" {
			if isDigit(s[i+1]) || string(s[i+1]) == "\\" {
				result.WriteRune(rune(s[i+1]))
				isEscaped = true
				continue
			}
			return "", ErrInvalidString
		}
		count, err := getCountCurr(i, s)
		if err != nil {
			return "", err
		}
		if count > 0 {
			result.WriteString(strings.Repeat(string(s[i-1]), count-1))
			continue
		}
		count, err = getCountNext(i, s)
		if err != nil {
			return "", err
		}
		if count == 0 {
			continue
		}
		result.WriteRune(rune(s[i]))
	}

	return result.String(), nil
}

func isDigit(b byte) bool {
	return unicode.IsDigit(rune(b))
}

func isValidFirstCharacter(i int, b byte) error {
	if i == 0 && isDigit(b) {
		return ErrInvalidString
	}
	return nil
}

func isValidNumber(i int, s string) error {
	if i != len(s)-1 {
		if isDigit(s[i]) && isDigit(s[i+1]) {
			return ErrInvalidString
		}
	}
	return nil
}

func getCountCurr(i int, s string) (int, error) {
	if isDigit(s[i]) {
		count, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return 0, ErrInvalidString
		}
		if count == 0 {
			count++
		}
		return count, nil
	}
	return -1, nil
}

func getCountNext(i int, s string) (int, error) {
	if i != len(s)-1 {
		if isDigit(s[i+1]) {
			count, err := strconv.Atoi(string(s[i+1]))
			if err != nil {
				return 0, ErrInvalidString
			}
			return count, nil
		}
		return -1, nil
	}
	return -1, nil
}
