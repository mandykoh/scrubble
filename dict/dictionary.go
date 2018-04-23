package dict

// Dictionary represents a function which can validate whether a word is allowed
// or not.
type Dictionary func(word string) (valid bool)
