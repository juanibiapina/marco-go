package scanner

import (
	"unicode"
)

var lexIdentifier matcher = or(r('-'), r('='), unicode.IsLetter)
