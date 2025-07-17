package sanitize

import (
	"golang.org/x/text/unicode/norm"
	"html"
	"net/url"
	"regexp"
	"strings"
)

var searchRegex = regexp.MustCompile(`[^a-zA-Z0-9@._\- ]+`)

func SearchQuery(query string) string {
	return searchRegex.ReplaceAllString(query, "")
}

// TrimInput removes leading and trailing whitespace.
func TrimInput(input string) string {
	return strings.TrimSpace(input)
}

// NormalizeUTF8 normalizes UTF-8 strings to NFC form.
func NormalizeUTF8(input string) string {
	return norm.NFC.String(input)
}

// RemoveControlChars removes non-printable/control characters from input.
func RemoveControlChars(input string) string {
	var b strings.Builder
	for _, r := range input {
		if r >= 32 && r != 127 {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// EscapeForHTML escapes special characters for safe HTML output.
func EscapeForHTML(input string) string {
	return html.EscapeString(input)
}

// URLEncode encodes the input string for safe URL usage.
func URLEncode(input string) string {
	return url.QueryEscape(input)
}

func SanitizeInput(input string) string {
	s := TrimInput(input)
	s = NormalizeUTF8(s)
	s = RemoveControlChars(s)
	s = EscapeForHTML(s)
	return s
}

func SanitizeEmail(input string) string {
	s := TrimInput(input)
	s = NormalizeUTF8(s)
	s = RemoveControlChars(s)
	s = EscapeForHTML(s)
	return s
}
