package gopherriz

import "strings"

// Given a string s, find the length of the longest substring
// without duplicate characters.

// Our solution

func lengthOfLongestSubstring(s string) int {
	max := 0
	unique := ""
	var currentLetter = ""
	for _, letter := range s {
		currentLetter = string(letter)
		if !strings.Contains(unique, currentLetter) {
			unique = unique + currentLetter
			subSLen := len(unique)
			if max < subSLen {
				max = subSLen
			}
		} else {

			unique = unique[strings.Index(unique, currentLetter)+1:] + currentLetter
		}
	}

	return max
}

// another way, more efficient way to handle the same problem.
// using constant space, with same time complexity, but much less actual code complexity

func lengthOfLongestSubstringOne(s string) int {
	Map := make(map[byte]bool)
	Left, Max := 0, 0

	for Right := 0; Right < len(s); Right++ {
		for Map[s[Right]] {
			delete(Map, s[Left])
			Left++
		}

		Map[s[Right]] = true

		temp := Right - Left + 1
		if temp > Max {
			Max = temp
		}
	}

	return Max
}
