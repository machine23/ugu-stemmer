package stemmer

type RussianStemmer struct{}

// NewRussianStemmer creates a new RussianStemmer.
func NewRussianStemmer() *RussianStemmer {
	return &RussianStemmer{}
}

// Stem returns the stem of the given word.
func (s *RussianStemmer) Stem(word string) string {
	panic("not implemented")
}
