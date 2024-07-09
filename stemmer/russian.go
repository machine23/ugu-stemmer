package stemmer

import (
	"slices"
	"strings"
)

var ruStopWords []string = []string{
	"и", "в", "во", "не", "что", "он", "на", "я", "с",
	"со", "как", "а", "то", "все", "она", "так", "его",
	"но", "да", "ты", "к", "у", "же", "вы", "за", "бы",
	"по", "только", "ее", "мне", "было", "вот", "от",
	"меня", "еще", "нет", "о", "из", "ему", "теперь",
	"когда", "даже", "ну", "вдруг", "ли", "если", "уже",
	"или", "ни", "быть", "был", "него", "до", "вас",
	"нибудь", "опять", "уж", "вам", "ведь", "там", "потом",
	"себя", "ничего", "ей", "может", "они", "тут", "где",
	"есть", "надо", "ней", "для", "мы", "тебя", "их",
	"чем", "была", "сам", "чтоб", "без", "будто", "чего",
	"раз", "тоже", "себе", "под", "будет", "ж", "тогда",
	"кто", "этот", "того", "потому", "этого", "какой",
	"совсем", "ним", "здесь", "этом", "один", "почти",
	"мой", "тем", "чтобы", "нее", "сейчас", "были", "куда",
	"зачем", "всех", "никогда", "можно", "при", "наконец",
	"два", "об", "другой", "хоть", "после", "над", "больше",
	"тот", "через", "эти", "нас", "про", "всего", "них",
	"какая", "много", "разве", "три", "эту", "моя",
	"впрочем", "хорошо", "свою", "этой", "перед", "иногда",
	"лучше", "чуть", "том", "нельзя", "такой", "им", "более",
	"всегда", "конечно", "всю", "между",
}

type RussianStemmer struct {
	stopWords []string
}

// NewRussianStemmer creates a new RussianStemmer.
func NewRussianStemmer() *RussianStemmer {
	sww := ruStopWords
	slices.Sort(sww)
	return &RussianStemmer{
		stopWords: sww,
	}
}

// Stem returns the stem of the given word.
func (s *RussianStemmer) Stem(word string) string {
	if s.IsStopWord(word) {
		return word
	}

	word = cyrillicToRoman(word)
	_ = word

	panic("not implemented")
}

// IsStopWord returns true if the given word is a stop word.
func (s *RussianStemmer) IsStopWord(word string) bool {
	_, found := slices.BinarySearch(s.stopWords, strings.ToLower(word))
	return found
}

// cyrillicToRoman transliterates the given Russian word into the Latin alphabet.
func cyrillicToRoman(word string) string {
	transliterations := map[rune]string{
		'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d",
		'е': "e", 'ё': "e", 'ж': "zh", 'з': "z", 'и': "i",
		'й': "j", 'к': "k", 'л': "l", 'м': "m", 'н': "n",
		'о': "o", 'п': "p", 'р': "r", 'с': "s", 'т': "t",
		'у': "u", 'ф': "f", 'х': "kh", 'ц': "C", 'ч': "ch",
		'ш': "sh", 'щ': "shch", 'ъ': "''", 'ы': "y", 'ь': "'",
		'э': "E", 'ю': "U", 'я': "A",
	}

	var result strings.Builder
	for _, r := range strings.ToLower(word) {
		if translit, ok := transliterations[r]; ok {
			result.WriteString(translit)
		} else {
			result.WriteRune(r) // Keep the character as is if no transliteration is found
		}
	}
	return result.String()
}

// romanToCyrillic transliterates the given Latin word into the Russian alphabet.
func romanToCyrillic(word string) string {
	transliterations := map[string]string{
		"a": "а", "b": "б", "v": "в", "g": "г", "d": "д",
		"e": "е", "zh": "ж", "z": "з", "i": "и",
		"j": "й", "k": "к", "l": "л", "m": "м", "n": "н",
		"o": "о", "p": "п", "r": "р", "s": "с", "t": "т",
		"u": "у", "f": "ф", "kh": "х", "C": "ц", "ch": "ч",
		"sh": "ш", "shch": "щ", "''": "ъ", "y": "ы", "'": "ь",
		"E": "э", "U": "ю", "A": "я",
	}

	var result strings.Builder
	for i := 0; i < len(word); {
		matched := false
		// Check for multi-character transliterations first
		for k, v := range transliterations {
			if strings.HasPrefix(word[i:], k) {
				result.WriteString(v)
				i += len(k)
				matched = true
				break
			}
		}
		// If no multi-character transliteration was found, move by one character
		if !matched {
			if translit, ok := transliterations[string(word[i])]; ok {
				result.WriteString(translit)
			} else {
				result.WriteByte(word[i]) // Keep the character as is if no transliteration is found
			}
			i++
		}
	}
	return result.String()
}
