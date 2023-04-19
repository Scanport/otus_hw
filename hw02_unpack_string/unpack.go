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
		count, err := getCountCurr(i, r)
		if err != nil {
			return "", err
		}
		if count > 0 {
			result.WriteString(strings.Repeat(string(r[i-1]), count-1))
			continue
		}
		count, err = getCountNext(i, r)
		if err != nil {
			return "", err
		}
		if count == 0 {
			continue
		}
		result.WriteRune(r[i])
	}

	return result.String(), nil
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
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

func getCountCurr(i int, r []rune) (int, error) {
	if isDigit(r[i]) {
		count, err := strconv.Atoi(string(r[i]))
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

func getCountNext(i int, r []rune) (int, error) {
	if i != len(r)-1 {
		if isDigit(r[i+1]) {
			count, err := strconv.Atoi(string(r[i+1]))
			if err != nil {
				return 0, ErrInvalidString
			}
			return count, nil
		}
		return -1, nil
	}
	return -1, nil
}
