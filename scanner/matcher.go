package scanner

type matcher func(rune) bool

// matches exact rune
func r(r rune) matcher {
	return func(c rune) bool {
		return c == r
	}
}

func or(ms ...matcher) matcher {
	return func(r rune) bool {
		for _, m := range ms {
			if m(r) {
				return true
			}
		}
		return false
	}
}
