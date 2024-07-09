package ugustemmer

import "github.com/machine23/ugu-stemmer/stemmer"

type Stemmer interface {
	Stem(word string) string
}

type SnowballStemmer struct {
	stemmer Stemmer
	lang    string
}

// NewSnowballStemmer creates a new SnowballStemmer for the given language.
// The language must be a valid ISO 639-1 code.
// If the language is not supported, the function will return nil.
// Supported languages are:
//   - "en" (English) - not implemented
//   - "es" (Spanish) - not implemented
//   - "fr" (French) - not implemented
//   - "it" (Italian) - not implemented
//   - "pt" (Portuguese) - not implemented
//   - "ru" (Russian) - not implemented
//   - "de" (German) - not implemented
//   - "nl" (Dutch) - not implemented
//   - "sv" (Swedish) - not implemented
//   - "no" (Norwegian) - not implemented
//   - "da" (Danish) - not implemented
//   - "fi" (Finnish) - not implemented
func NewSnowballStemmer(lang string) *SnowballStemmer {
	stemmers := map[string]Stemmer{
		"ru": stemmer.NewRussianStemmer(),
	}
	stemmer, ok := stemmers[lang]
	if !ok {
		return nil
	}
	return &SnowballStemmer{
		stemmer: stemmer,
		lang:    lang,
	}
}

// Stem returns the stem of the given word.
// If the language is not supported, the function will return the word unchanged.
func (s *SnowballStemmer) Stem(word string) string {
	return s.stemmer.Stem(word)
}
