package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(sourceString string) (string, error) {
	resultStr := strings.Builder{}

	previousChar := rune(0)
	isPreviousSlash := false
	isPreviousRealDigit := true

	for _, currentChar := range sourceString {
		if previousChar == '\\' && currentChar == '\\' && !isPreviousSlash {
			isPreviousSlash = true
			resultStr.WriteRune(previousChar)
		} else if unicode.IsDigit(currentChar) && previousChar == '\\' && !isPreviousSlash {
			resultStr.WriteRune(currentChar)
			isPreviousRealDigit = false
		} else {
			if unicode.IsDigit(currentChar) {
				if previousChar == 0 {
					return "", ErrInvalidString
				}

				digit := currentChar - '0'

				if isPreviousSlash || !isPreviousRealDigit {
					digit--
				}

				if unicode.IsDigit(previousChar) && isPreviousRealDigit {
					return "", ErrInvalidString
				}

				isPreviousRealDigit = true

				resultStr.WriteString(strings.Repeat(string(previousChar), int(digit)))
			} else if previousChar != 0 && !unicode.IsDigit(previousChar) && !isPreviousSlash {
				resultStr.WriteRune(previousChar)
			}

			isPreviousSlash = false
		}

		previousChar = currentChar
	}

	if previousChar != 0 && !unicode.IsDigit(previousChar) {
		resultStr.WriteRune(previousChar)
	}

	return resultStr.String(), nil
}
