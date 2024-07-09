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

	panic("not implemented")
}

// IsStopWord returns true if the given word is a stop word.
func (s *RussianStemmer) IsStopWord(word string) bool {
	_, found := slices.BinarySearch(s.stopWords, strings.ToLower(word))
	return found
}
