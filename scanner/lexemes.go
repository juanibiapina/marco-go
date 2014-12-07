package scanner

import (
	"unicode"
)

var lexIdentifier matcher = or(r('-'), unicode.IsLetter)
