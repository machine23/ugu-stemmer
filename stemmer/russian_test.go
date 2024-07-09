package stemmer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRussianStemmer(t *testing.T) {
	s := NewRussianStemmer()
	require.NotNil(t, s)
}

func TestRussianStemmer_IsStopWord(t *testing.T) {
	s := NewRussianStemmer()
	require.True(t, s.IsStopWord("и"))
	require.True(t, s.IsStopWord("И"))
	require.False(t, s.IsStopWord("яблоко"))
}

func TestRussianStemmer_Stem(t *testing.T) {
	t.Run("stop word", func(t *testing.T) {
		s := NewRussianStemmer()
		require.Equal(t, "и", s.Stem("и"))
		require.Equal(t, "И", s.Stem("И"))
	})
}

func TestCyrillicToRoman(t *testing.T) {
	f := func(cyrillic, roman string) {
		t.Helper()
		require.Equal(t, roman, cyrillicToRoman(cyrillic))
	}

	f("", "")
	f("и", "i")
	f("И", "i")
	f("быстрый", "bystryj")
	f("яблоко", "Abloko")
	f("Яблоко", "Abloko")
	f("привет", "privet")
	f("foo", "foo")
	f("s',", "s',")
}

func Test_romanToCyrillic(t *testing.T) {
	f := func(roman, cyrillic string) {
		t.Helper()
		require.Equal(t, cyrillic, romanToCyrillic(roman))
	}

	f("", "")
	f("i", "и")
	f("bystryj", "быстрый")
	f("Abloko", "яблоко")
	f("privet", "привет")
}
