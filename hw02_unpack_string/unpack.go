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
		if i == 0 && isDigit(s[i]) {
			return "", ErrInvalidString
		}
		if i != len(s)-1 {
			if isDigit(s[i]) && isDigit(s[i+1]) {
				return "", ErrInvalidString
			}
		}
		if string(s[i]) == "\\" {
			if isDigit(s[i+1]) || string(s[i+1]) == "\\" {
				result.WriteRune(rune(s[i+1]))
				isEscaped = true
				continue
			}
			return "", ErrInvalidString
		}
		if isDigit(s[i]) {
			count, err := strconv.Atoi(string(s[i]))
			if err != nil {
				return "", ErrInvalidString
			}
			if count == 0 {
				count++
			}
			result.WriteString(strings.Repeat(string(s[i-1]), count-1))
			continue
		}
		if i != len(s)-1 {
			if isDigit(s[i+1]) {
				count, _ := strconv.Atoi(string(s[i+1]))
				if count == 0 {
					continue
				}
			}
		}
		result.WriteRune(rune(s[i]))
	}

	return result.String(), nil
}

func isDigit(b byte) bool {
	return unicode.IsDigit(rune(b))
}
