package slasher

func toggleQuote(r rune, b, s, d *bool) {
	switch r {
	case '`':
		*b = !(*b)
	case '\'':
		*s = !(*s)
	case '"':
		*d = !(*d)
	}
}

func Slasher(cmd string) string {
	isBacktickOpened := false
	isSingleQuoteOpened := false
	isDoubleQuoteOpened := false
	isDash := false
	result := ""

	var prevRune rune

	for _, r := range cmd {
		toggleQuote(r, &isBacktickOpened, &isSingleQuoteOpened, &isDoubleQuoteOpened)
		if isBacktickOpened || isSingleQuoteOpened || isDoubleQuoteOpened {
			result += string(r)
			continue
		}

		if r == '-' {
			if !isDash && prevRune == ' ' {
				result += "\\\n  "
				isDash = true
			}
		} else {
			isDash = false
		}

		result += string(r)

		prevRune = r
	}

	return result
}
