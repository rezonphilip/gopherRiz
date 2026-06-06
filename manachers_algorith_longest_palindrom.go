package gopherriz

func longestPalindrome(s string) string {
	transformed := "#"
	for _, letter := range s {
		transformed += string(letter) + "#"
	}

	n := len(transformed)
	p := make([]int, n)
	c, r := 0, 0
	mirror := 0
	for i := 0; i < n; i++ {
		mirror = c*2 - i
		if i < r {
			p[i] = min(r-i, p[mirror])
		}

		for i+p[i]+1 < n && i-p[i]-1 >= 0 && transformed[i+p[i]+1] == transformed[i-p[i]-1] {
			p[i] += 1
		}

		if i+p[i] > r {
			c = i
			r = c + p[c]
		}
	}

	maxLength, centre := 0, 0
	for i, v := range p {
		if v > maxLength {
			maxLength = v
			centre = i
		}
	}

	start := (centre - maxLength) / 2
	return s[start : start+maxLength]

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
