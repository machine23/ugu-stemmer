package stemmer

import (
	"slices"
	"strings"
)

var (
	ruStopWords []string = []string{
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

	perfectiveSuffixes = []string{
		"ivwis'",
		"yvwis'",
		"vwis'",
		"ivwi",
		"yvwi",
		"vwi",
		"iv",
		"yv",
		"v",
	}

	adjectivalSuffixes = []string{
		"uUWUU",
		"uUWAA",
		"uUWimi",
		"uUWymi",
		"uUWego",
		"uUWogo",
		"uUWemu",
		"uUWomu",
		"uUWih",
		"uUWyh",
		"uUWuU",
		"uUWaia",
		"uUWoU",
		"uUWeU",
		"UWUU",
		"UWAA",
		"uUWee",
		"uUWie",
		"uUWye",
		"uUWoe",
		"uUWej",
		"uUWij",
		"uUWyj",
		"uUWoj",
		"uUWem",
		"uUWim",
		"uUWym",
		"uUWom",
		"UWimi",
		"UWymi",
		"UWego",
		"UWogo",
		"UWemu",
		"UWomu",
		"UWih",
		"UWyh",
		"UWuU",
		"UWaA",
		"UWoU",
		"UWeU",
		"UWee",
		"UWie",
		"UWye",
		"UWoe",
		"UWej",
		"UWij",
		"UWyj",
		"UWoj",
		"UWem",
		"UWim",
		"UWym",
		"UWom",
		"WUU",
		"WAA",
		"ivwUU",
		"ivwAA",
		"yvwUU",
		"yvwAA",
		"Wimi",
		"Wymi",
		"Wego",
		"Wogo",
		"Wemu",
		"Womu",
		"Wih",
		"Wyh",
		"WuU",
		"WaA",
		"WoU",
		"WeU",
		"ivwimi",
		"ivwymi",
		"ivwego",
		"ivwogo",
		"ivwemu",
		"ivwomu",
		"ivwih",
		"ivwyh",
		"ivwuU",
		"ivwaA",
		"ivwoU",
		"ivweU",
		"yvwimi",
		"yvwymi",
		"yvwego",
		"yvwogo",
		"yvwemu",
		"yvwomu",
		"yvwih",
		"yvwyh",
		"yvwuU",
		"yvwaA",
		"yvwoU",
		"yvweU",
		"vwUU",
		"vwAA",
		"Wee",
		"Wie",
		"Wye",
		"Woe",
		"Wej",
		"Wij",
		"Wyj",
		"Woj",
		"Wem",
		"Wim",
		"Wym",
		"Wom",
		"ivwee",
		"ivwie",
		"ivwye",
		"ivwoe",
		"ivwej",
		"ivwij",
		"ivwyj",
		"ivwoj",
		"ivwem",
		"ivwim",
		"ivwym",
		"ivwom",
		"yvwee",
		"yvwie",
		"yvwye",
		"yvwoe",
		"yvwej",
		"yvwij",
		"yvwyj",
		"yvwoj",
		"yvwem",
		"yvwim",
		"yvwym",
		"yvwom",
		"vwimi",
		"vwymi",
		"vwego",
		"vwogo",
		"vwemu",
		"vwomu",
		"vwih",
		"vwyh",
		"vwuU",
		"vwaA",
		"vwoU",
		"vweU",
		"emUU",
		"emAA",
		"nnUU",
		"nnAA",
		"vwee",
		"vwie",
		"vwye",
		"vwoe",
		"vwej",
		"vwij",
		"vwyj",
		"vwoj",
		"vwem",
		"vwim",
		"vwym",
		"vwom",
		"emimi",
		"emymi",
		"emego",
		"emogo",
		"ememu",
		"emomu",
		"emih",
		"emyh",
		"emuU",
		"emaA",
		"emoU",
		"emeU",
		"nnimi",
		"nnymi",
		"nnego",
		"nnogo",
		"nnemu",
		"nnomu",
		"nnih",
		"nnyh",
		"nnuU",
		"nnaA",
		"nnoU",
		"nneU",
		"emee",
		"emie",
		"emye",
		"emoe",
		"emej",
		"emij",
		"emyj",
		"emoj",
		"emem",
		"emim",
		"emym",
		"emom",
		"nnee",
		"nnie",
		"nnye",
		"nnoe",
		"nnej",
		"nnij",
		"nnyj",
		"nnoj",
		"nnem",
		"nnim",
		"nnym",
		"nnom",
		"UU",
		"AA",
		"imi",
		"ymi",
		"ego",
		"ogo",
		"emu",
		"omu",
		"ih",
		"yh",
		"uU",
		"aA",
		"oU",
		"eU",
		"ee",
		"ie",
		"ye",
		"oe",
		"ej",
		"ij",
		"yj",
		"oj",
		"em",
		"im",
		"ym",
		"om",
	}

	adjectivalSuffixes2 = []string{
		"UWUU",
		"UWAA",
		"UWuU",
		"UWaA",
		"UWoU",
		"UWeU",
		"UWimi",
		"UWymi",
		"UWego",
		"UWogo",
		"UWemu",
		"UWomu",
		"UWih",
		"UWyh",
		"WUU",
		"WAA",
		"UWee",
		"UWie",
		"UWye",
		"UWoe",
		"UWej",
		"UWij",
		"UWyj",
		"UWoj",
		"UWem",
		"UWim",
		"UWym",
		"UWom",
		"vwUU",
		"vwAA",
		"WuU",
		"WaA",
		"WoU",
		"WeU",
		"emUU",
		"emAA",
		"nnUU",
		"nnAA",
		"Wimi",
		"Wymi",
		"Wego",
		"Wogo",
		"Wemu",
		"Womu",
		"Wih",
		"Wyh",
		"vwuU",
		"vwaA",
		"vwoU",
		"vweU",
		"Wee",
		"Wie",
		"Wye",
		"Woe",
		"Wej",
		"Wij",
		"Wyj",
		"Woj",
		"Wem",
		"Wim",
		"Wym",
		"Wom",
		"vwimi",
		"vwymi",
		"vwego",
		"vwogo",
		"vwemu",
		"vwomu",
		"vwih",
		"vwyh",
		"emuU",
		"emaA",
		"emoU",
		"emeU",
		"nnuU",
		"nnaA",
		"nnoU",
		"nneU",
		"vwee",
		"vwie",
		"vwye",
		"vwoe",
		"vwej",
		"vwij",
		"vwyj",
		"vwoj",
		"vwem",
		"vwim",
		"vwym",
		"vwom",
		"emimi",
		"emymi",
		"emego",
		"emogo",
		"ememu",
		"emomu",
		"emih",
		"emyh",
		"nnimi",
		"nnymi",
		"nnego",
		"nnogo",
		"nnemu",
		"nnomu",
		"nnih",
		"nnyh",
		"emee",
		"emie",
		"emye",
		"emoe",
		"emej",
		"emij",
		"emyj",
		"emoj",
		"emem",
		"emim",
		"emym",
		"emom",
		"nnee",
		"nnie",
		"nnye",
		"nnoe",
		"nnej",
		"nnij",
		"nnyj",
		"nnoj",
		"nnem",
		"nnim",
		"nnym",
		"nnom",
	}

	reflexiveSuffixes = []string{
		"sA",
		"s'",
	}

	verbSuffixes = []string{
		"ew'",
		"ejte",
		"ujte",
		"uUt",
		"iw'",
		"ete",
		"jte",
		"Ut",
		"nno",
		"ila",
		"yla",
		"ena",
		"ite",
		"ili",
		"yli",
		"ilo",
		"ylo",
		"eno",
		"At",
		"uet",
		"eny",
		"it'",
		"yt'",
		"uU",
		"la",
		"na",
		"li",
		"em",
		"lo",
		"no",
		"et",
		"ny",
		"t'",
		"ej",
		"uj",
		"il",
		"yl",
		"im",
		"ym",
		"en",
		"it",
		"yt",
		"U",
		"j",
		"l",
		"n",
	}

	verbSuffixes2 []string = []string{
		"la",
		"na",
		"ete",
		"jte",
		"li",
		"j",
		"l",
		"em",
		"n",
		"lo",
		"no",
		"et",
		"Ut",
		"ny",
		"t'",
		"ew'",
		"nno",
	}

	nounSuffixes = []string{
		"iAmi",
		"iAh",
		"Ami",
		"iAm",
		"Ah",
		"ami",
		"iej",
		"Am",
		"iem",
		"ah",
		"iU",
		"'U",
		"iA",
		"'A",
		"ev",
		"ov",
		"ie",
		"'e",
		"ei",
		"ii",
		"ej",
		"oj",
		"ij",
		"em",
		"am",
		"om",
		"U",
		"A",
		"a",
		"e",
		"i",
		"j",
		"o",
		"u",
		"y",
		"'",
	}

	superlativeSuffixes  = []string{"ejwe", "ejw"}
	derivationalSuffixes = []string{"ost'", "ost"}
)

type RussianStemmer struct {
	stopWords []string
	// superlativeSuffixes  []string
	// derivationalSuffixes []string
	// nounSuffixes         []string
	// verbSuffixes         []string
	verbSuffixes2 []string
	// adjectivalSuffixes   []string
	adjectivalSuffixes2 []string
	// reflexiveSuffixes    []string
	// perfectiveSuffixes   []string
}

// NewRussianStemmer creates a new RussianStemmer.
func NewRussianStemmer() *RussianStemmer {
	sww := ruStopWords
	slices.Sort(sww)

	adjs2 := adjectivalSuffixes2
	slices.Sort(adjs2)

	verb2 := verbSuffixes2
	slices.Sort(verb2)
	return &RussianStemmer{
		stopWords:           sww,
		verbSuffixes2:       verb2,
		adjectivalSuffixes2: adjs2,
	}
}

// Stem returns the stem of the given word.
func (s RussianStemmer) Stem(word string) string {
	word = strings.ToLower(word)
	if s.isStopWord(word) {
		return word
	}

	word = cyrillicToRoman(word)
	rv, r2 := s.regions(word)

	word, rv, r2 = s.step1(word, rv, r2)
	word, r2 = s.step2(word, rv, r2)
	word = s.step3(word, r2)
	word = s.step4(word)

	return romanToCyrillic(word)
}

func (s RussianStemmer) step1(word, rv, r2 string) (string, string, string) {
	for _, suffix := range perfectiveSuffixes {
		if strings.HasSuffix(rv, suffix) {
			suffixLen := len(suffix)
			if suffix == "v" || suffix == "vwi" || suffix == "vwis'" {
				if len(rv) > suffixLen && (rv[len(rv)-suffixLen-1:len(rv)-suffixLen] == "A" || rv[len(rv)-suffixLen-1:len(rv)-suffixLen] == "a") {
					return s.trimSuffix(word, rv, r2, suffix)
				}
			} else {
				return s.trimSuffix(word, rv, r2, suffix)
			}
		}
	}

	for _, suffix := range reflexiveSuffixes {
		if strings.HasSuffix(rv, suffix) {
			word, rv, r2 = s.trimSuffix(word, rv, r2, suffix)
			break
		}
	}

	for _, suffix := range adjectivalSuffixes {
		if strings.HasSuffix(rv, suffix) {
			suffixLen := len(suffix)
			_, found := slices.BinarySearch(s.adjectivalSuffixes2, suffix)
			if found {
				if len(rv) >= suffixLen+1 && (strings.HasSuffix(rv[:len(rv)-suffixLen], "A") || strings.HasSuffix(rv[:len(rv)-suffixLen], "a")) {
					return s.trimSuffix(word, rv, r2, suffix)
				}
			} else {
				return s.trimSuffix(word, rv, r2, suffix)
			}
		}
	}

	for _, suffix := range verbSuffixes {
		if strings.HasSuffix(rv, suffix) {
			suffixLen := len(suffix)
			if _, found := slices.BinarySearch(s.verbSuffixes2, suffix); found {
				if len(rv) >= suffixLen+1 && (strings.HasSuffix(rv[:len(rv)-suffixLen], "A") || strings.HasSuffix(rv[:len(rv)-suffixLen], "a")) {
					return s.trimSuffix(word, rv, r2, suffix)
				}
			} else {
				return s.trimSuffix(word, rv, r2, suffix)
			}
		}
	}

	for _, suffix := range nounSuffixes {
		if strings.HasSuffix(rv, suffix) {
			return s.trimSuffix(word, rv, r2, suffix)
		}
	}

	return word, rv, r2
}

func (s RussianStemmer) trimSuffix(word, rv, r2, suffix string) (string, string, string) {
	suffixLen := len(suffix)
	if len(word) >= suffixLen {
		word = word[:len(word)-suffixLen]
	}
	if len(r2) >= suffixLen {
		r2 = r2[:len(r2)-suffixLen]
	}
	if len(rv) >= suffixLen {
		rv = rv[:len(rv)-suffixLen]
	}
	return word, rv, r2
}

func (s RussianStemmer) step2(word, rv, r2 string) (string, string) {
	if strings.HasSuffix(rv, "i") {
		word = word[:len(word)-1]
		r2 = r2[:len(r2)-1]
	}
	return word, r2
}

func (s RussianStemmer) step3(word, r2 string) string {
	for _, suffix := range derivationalSuffixes {
		if strings.HasSuffix(r2, suffix) {
			return word[:len(word)-len(suffix)]
		}
	}
	return word
}

func (s RussianStemmer) step4(word string) string {
	var superlativeRemoved bool

	if strings.HasSuffix(word, "nn") {
		return word[:len(word)-1]
	}

	for _, suffix := range superlativeSuffixes {
		if strings.HasSuffix(word, suffix) {
			word = word[:len(word)-len(suffix)]
			superlativeRemoved = true
			break
		}
	}

	if strings.HasSuffix(word, "nn") {
		word = word[:len(word)-1]
	}

	if !superlativeRemoved {
		word = strings.TrimSuffix(word, "'")
	}

	return word
}

// isStopWord returns true if the given word is a stop word.
func (s RussianStemmer) isStopWord(word string) bool {
	_, found := slices.BinarySearch(s.stopWords, word)
	return found
}

var russianVowelSet = map[rune]bool{
	'A': true,
	'U': true,
	'E': true,
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'y': true,
}

func (s RussianStemmer) regions(word string) (string, string) {
	r1Start, r2Start, rvStart := len(word), len(word), len(word)

	// Find RV
	for i, r := range word {
		if _, isVowel := russianVowelSet[r]; isVowel {
			rvStart = i + 1
			break
		}
	}

	// Find R1 and R2 in a single pass
	for i := rvStart; i < len(word)-1; i++ {
		if _, isVowelCurr := russianVowelSet[rune(word[i])]; !isVowelCurr {
			if _, isVowelPrev := russianVowelSet[rune(word[i-1])]; isVowelPrev {
				if r1Start == len(word) { // First time finding R1
					r1Start = i + 1
				} else if i > r1Start { // Finding R2
					r2Start = i + 1
					break
				}
			}
		}
	}

	// Slicing strings based on calculated positions
	rv := ""
	r2 := ""
	if rvStart < len(word) {
		rv = word[rvStart:]
	}
	if r2Start < len(word) {
		r2 = word[r2Start:]
	}

	return rv, r2
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
