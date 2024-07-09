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
