package scanner

import (
	"unicode"
)

func lexIdentifier(r rune) bool {
	if unicode.IsLetter(r) {
		return true
	}

	if r == '-' {
		return true
	}

	return false
}
