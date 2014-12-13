package scanner

import (
	"unicode"
)

var lexIdentifier matcher = or(unicode.IsLetter, r('+'), r('-'), r('*'), r('/'), r('='), r('?'))
