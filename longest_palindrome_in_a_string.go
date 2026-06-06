package gopherriz

func longestPalindromeAlternative(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLength := 0, 1
	expand := func(l, i int) {
		for l >= 0 && i < len(s) && s[l] == s[i] {
			if i-l+1 > maxLength {
				start = l
				maxLength = i - l + 1
			}
			l = l - 1
			i = i + 1
		}
	}

	for i, _ := range s {
		expand(i, i)
		expand(i, i+1)
	}

	return s[start : start+maxLength]
}
