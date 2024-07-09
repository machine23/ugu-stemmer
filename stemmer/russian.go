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

var cyrillicToLatinMap = map[rune]rune{
	'а': 'a', 'б': 'b', 'в': 'v', 'г': 'g', 'д': 'd',
	'е': 'e', 'ё': 'e', 'ж': 'x', 'з': 'z', 'и': 'i',
	'й': 'j', 'к': 'k', 'л': 'l', 'м': 'm', 'н': 'n',
	'о': 'o', 'п': 'p', 'р': 'r', 'с': 's', 'т': 't',
	'у': 'u', 'ф': 'f', 'х': 'h', 'ц': 'c', 'ч': 'H',
	'ш': 'w', 'щ': 'W', 'ъ': '"', 'ы': 'y', 'ь': '\'',
	'э': 'E', 'ю': 'U', 'я': 'A',

	'А': 'a', 'Б': 'b', 'В': 'v', 'Г': 'g', 'Д': 'd',
	'Е': 'e', 'Ё': 'e', 'Ж': 'x', 'З': 'z', 'И': 'i',
	'Й': 'j', 'К': 'k', 'Л': 'l', 'М': 'm', 'Н': 'n',
	'О': 'o', 'П': 'p', 'Р': 'r', 'С': 's', 'Т': 't',
	'У': 'u', 'Ф': 'f', 'Х': 'h', 'Ц': 'c', 'Ч': 'H',
	'Ш': 'w', 'Щ': 'W', 'Ъ': '"', 'Ы': 'y', 'Ь': '\'',
	'Э': 'E', 'Ю': 'U', 'Я': 'A',
}

// cyrillicToRoman transliterates the given Russian word into the Latin alphabet.
func cyrillicToRoman(word string) string {
	// Preallocate builder capacity based on input length to minimize allocations.
	var result strings.Builder
	result.Grow(len(word))
	for _, r := range word {
		if translit, ok := cyrillicToLatinMap[r]; ok {
			result.WriteRune(translit)
		} else {
			result.WriteRune(r) // Keep the character as is if no transliteration is found
		}
	}
	return result.String()
}

// Creating the reverse map from the transliterations map
var latinToCyrillicMap = map[rune]rune{
	'a': 'а', 'b': 'б', 'v': 'в', 'g': 'г', 'd': 'д',
	'e': 'е', 'x': 'ж', 'z': 'з', 'i': 'и', 'j': 'й',
	'k': 'к', 'l': 'л', 'm': 'м', 'n': 'н', 'o': 'о',
	'p': 'п', 'r': 'р', 's': 'с', 't': 'т', 'u': 'у',
	'f': 'ф', 'h': 'х', 'c': 'ц', 'H': 'ч', 'w': 'ш',
	'W': 'щ', '"': 'ъ', 'y': 'ы', '\'': 'ь', 'E': 'э',
	'U': 'ю', 'A': 'я',
}

// romanToCyrillic transliterates the given Latin word into the Russian alphabet.
func romanToCyrillic(word string) string {
	var result strings.Builder
	result.Grow(len(word)) // Preallocate builder capacity to minimize allocations

	for _, char := range word {
		if translit, ok := latinToCyrillicMap[char]; ok {
			result.WriteRune(translit)
		} else {
			result.WriteRune(char) // Keep the character as is if no transliteration is found
		}
	}
	return result.String()
}
