package validation

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

var ErrValidation = errors.New("validation error")
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
var yearFormatRegex = regexp.MustCompile(`^\d{4}_(1|2)$`)

func isAlphanumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func isEmail(s string) bool {
	return emailRegex.MatchString(s)
}

func IsValidClassFormat(input string) bool {
	return yearFormatRegex.MatchString(input)
}

func ValidatePage(input int) error {
	if input <= 0 && input >= 100 {
		return fmt.Errorf("%w: the page value must be between 0 and 100, received value: "+strconv.Itoa(input), ErrValidation)
	}

	return nil
}

func ValidateLimit(input int) error {
	if input < 5 && input > 50 {
		return fmt.Errorf("%w: the limit value must be between 5 and 50, received value: "+strconv.Itoa(input), ErrValidation)
	}
	return nil
}
