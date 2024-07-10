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
	f("Подъезд", "pod\"ezd")
	f("держаться", "derxat'sA")
	f("привет", "privet")
	f("щёлкать", "Welkat'")
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
	f("pod\"ezd", "подъезд")
	f("derxat'sA", "держаться")
}

func TestRussianStemmer_regions(t *testing.T) {
	f := func(word, rv, r2 string) {
		t.Helper()
		s := NewRussianStemmer()
		nrv, nr2 := s.regions(word)
		require.Equal(t, rv, nrv)
		require.Equal(t, r2, nr2)
	}

	f("vesna", "sna", "")
	f("prygat'", "gat'", "'")
	f("knigu", "gu", "")
	f("igraet", "graet", "")
	f("solncem", "lncem", "")
	f("Hital", "tal", "")
	f("cvetka", "tka", "")
	f("dumaU", "maU", "")
	f("lUbov'U", "bov'U", "'U")
	f("hodili", "dili", "i")
	f("rabotaew'", "botaew'", "aew'")
	f("pesni", "sni", "")
	f("begut", "gut", "")
	f("vody", "dy", "")
	f("pisala", "sala", "a")
	f("uHiw'sA", "Hiw'sA", "'sA")
	f("osen'U", "sen'U", "'U")
	f("meHtaem", "Htaem", "")
	f("stolu", "lu", "")
	f("letali", "tali", "i")
	f("derevom", "revom", "om")
	f("gotovlU", "tovlU", "lU")
	f("kowek", "wek", "")
	f("spim", "m", "")
	f("oblaka", "blaka", "a")
	f("smotrel", "trel", "")
	f("list'ev", "st'ev", "")
	f("risuem", "suem", "")
	f("plavali", "vali", "i")
	f("gorode", "rode", "e")
	f("beseduU", "seduU", "uU")
	f("mawinami", "winami", "ami")
	f("Hitaew'", "taew'", "'")
	f("tvoril", "ril", "")
	f("edoj", "doj", "")
	f("pryxke", "xke", "")
	f("sidim", "dim", "")
	f("gorelo", "relo", "o")
	f("pomniw'", "mniw'", "'")
	f("lesom", "som", "")
	f("idut", "dut", "")
	f("govoril", "voril", "il")
	f("sHast'em", "st'em", "")
	f("begaew'", "gaew'", "'")
	f("znal", "l", "")
	f("uHilis'", "Hilis'", "is'")
	f("rybami", "bami", "i")
	f("vidim", "dim", "")
	f("leto", "to", "")
	f("samosoverwenstvovanie", "mosoverwenstvovanie", "overwenstvovanie")
	f("vysokointellektual'nyj", "sokointellektual'nyj", "ointellektual'nyj")
	f("obWepriznannyj", "bWepriznannyj", "riznannyj")
	f("nepredskazuemost'", "predskazuemost'", "skazuemost'")
	f("poluprovodnikovyj", "luprovodnikovyj", "rovodnikovyj")
	f("al'ternativnost'", "l'ternativnost'", "nativnost'")
	f("mnogoobeWaUWij", "goobeWaUWij", "eWaUWij")
	f("izobrazitel'noe", "zobrazitel'noe", "razitel'noe")
	f("aglomeracionnyj", "glomeracionnyj", "eracionnyj")
	f("vysokokvalificirovannyj", "sokokvalificirovannyj", "okvalificirovannyj")
	f("perpendikulArnyj", "rpendikulArnyj", "dikulArnyj")
	f("Emocional'nost'", "mocional'nost'", "ional'nost'")
	f("organizovannost'", "rganizovannost'", "izovannost'")
	f("predprinimatel'stvo", "dprinimatel'stvo", "imatel'stvo")
	f("konstruktivnost'", "nstruktivnost'", "tivnost'")
	f("gipotetiHeskij", "potetiHeskij", "etiHeskij")
	f("sel'skohozAjstvennyj", "l'skohozAjstvennyj", "ozAjstvennyj")
	f("mnogoznaHitel'nyj", "goznaHitel'nyj", "naHitel'nyj")
	f("giperinflAciA", "perinflAciA", "inflAciA")
	f("materializovat'sA", "terializovat'sA", "ializovat'sA")
	f("razvlekatel'nyj", "zvlekatel'nyj", "atel'nyj")
	f("programmirovanie", "grammirovanie", "mirovanie")
	f("mnogoslojnost'", "goslojnost'", "lojnost'")
	f("Eksperimental'nyj", "ksperimental'nyj", "imental'nyj")
	f("informacionnyj", "nformacionnyj", "macionnyj")
	f("EkologiHeski", "kologiHeski", "ogiHeski")
	f("protivopoloxnost'", "tivopoloxnost'", "opoloxnost'")
	f("strukturirovannyj", "kturirovannyj", "irovannyj")
	f("protivodejstvie", "tivodejstvie", "odejstvie")
	f("Elektrooborudovanie", "lektrooborudovanie", "trooborudovanie")
}
