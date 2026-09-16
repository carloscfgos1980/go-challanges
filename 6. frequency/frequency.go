package frequency

import (
	"regexp"
	"strings"
)

// CountWordFrequency takes a string containing multiple words and returns
// a map where each key is a word and the value is the number of times that
// word appears in the string. The comparison is case-insensitive.
//
// Words are defined as sequences of letters and digits.
// All words are converted to lowercase before counting.
// All punctuation, spaces, and other non-alphanumeric characters are ignored.
//
// For example:
// Input: "The quick brown fox jumps over the lazy dog."
// Output: map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1}
func CountWordFrequency(text string) map[string]int {
	wordCount := make(map[string]int)
	lowerCase := strings.ToLower(text)
	re := regexp.MustCompile(`[a-z0-9']+`)
	words := re.FindAllString(lowerCase, -1)

	for _, word := range words {
		// Keep contractions as a single word by stripping apostrophes.
		word = strings.ReplaceAll(word, "'", "")
		if word == "" {
			continue
		}

		if _, ok := wordCount[word]; !ok {
			wordCount[word] = 1
		} else {
			wordCount[word]++
		}
	}
	return wordCount
}
