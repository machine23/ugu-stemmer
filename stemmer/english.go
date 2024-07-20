package stemmer

type EnglishStemmer struct{}

func NewEnglishStemmer() *EnglishStemmer {
	return &EnglishStemmer{}
}

func (s *EnglishStemmer) Stem(word string) string {
	return word
}
