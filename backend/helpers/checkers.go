package helpers

import "unicode"

func IsNumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
