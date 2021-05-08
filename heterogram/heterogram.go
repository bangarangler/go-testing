package heterogram

import (
	"strings"
	"unicode"
)

func IsHeterogram(s string) bool {
	str := strings.ToLower(s)
	exists := make([]rune, 0)

	for _, char := range str {
		if unicode.IsLetter(char) {
			if contains(exists, char) {
				return false
			} else {
				exists = append(exists, char)
			}
		}
	}
	return true
}

func contains(r []rune, char rune) bool {
	for _, v := range r {
		if v == char {
			return true
		}
	}
	return false
}
