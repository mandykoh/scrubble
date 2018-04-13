package scrubble

import "strings"

// DefaultEnglishDictionary validates words against a default English word list.
//
// This dictionary is based on the public domain ENABLE word list collated by
// Alan Beale.
func DefaultEnglishDictionary(word string) (isValid bool) {
	return defaultEnglishDictionaryWords[strings.ToLower(word)]
}
