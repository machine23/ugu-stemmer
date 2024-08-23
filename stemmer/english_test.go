package stemmer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnglishStemmer_normalizeApostrophes(t *testing.T) {
	s := NewEnglishStemmer()

	f := func(input, expected string) {
		actual := s.normalizeApostrophes(input)
		require.Equal(t, expected, actual)
	}

	f("example", "example")

	f("'example", "example")
	f("‘example", "example")
	f("’example", "example")
	f("‛example", "example")

	f("example's", "example's")
	f("example‘s", "example's")
	f("example’s", "example's")
	f("example‛s", "example's")
	f("'example's", "example's")
	f("‘example‘s", "example's")
	f("’example’s", "example's")
	f("‛example‛s", "example's")
}
