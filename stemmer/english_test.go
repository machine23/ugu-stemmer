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

func TestEnglishStemmer_replaceYAfterVowel(t *testing.T) {
	s := NewEnglishStemmer()

	f := func(input, expected string) {
		actual := s.replaceYAfterVowel(input)
		require.Equal(t, expected, actual)
	}

	// Test cases where 'y' is after a vowel
	f("playing", "plaYing")
	f("staying", "staYing")
	f("buying", "buYing")

	// Test cases where 'y' is not after a vowel
	f("yellow", "yellow")
	f("rhythm", "rhythm")

	// Test cases without 'y'
	f("example", "example")
	f("test", "test")

	// Edge cases
	f("", "")
	f("y", "y")
	f("a", "a")
	f("ay", "aY")
}
