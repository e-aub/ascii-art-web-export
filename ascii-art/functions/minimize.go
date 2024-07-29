package functions

import (
	"sort"
	"strings"
)

func Minimize(str string) []rune {
	str = strings.ReplaceAll(str, "\n", "")
	var result []rune
	for _, letter := range str {
		if !strings.Contains(string(result), string(letter)) {
			result = append(result, letter)
		}
	}
	return sortRunes(result)
}

func sortRunes(runes []rune) []rune {
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return runes
}

func FilterText(s string) string {
	var result string
	for _, letter := range s {
		if !((letter >= ' ' && letter <= '~') || letter == '\n') {
			continue
		}
		result += string(letter)

	}
	return result
}
