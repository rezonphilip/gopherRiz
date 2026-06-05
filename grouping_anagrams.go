package gopherriz

import "sort"

// Given a list of words, group together
// those that are anagrams of each other ("eat", "tea", "ate" belong together).

func groupAnagrams(words []string) [][]string {
	groups := make(map[string][]string)

	for _, word := range words {
		str := []byte(word)
		sort.Slice(str, func(i, j int) bool { return str[i] < str[j] })
		groups[string(str)] = append(groups[string(str)], word)
	}

	result := make([][]string, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result

}
