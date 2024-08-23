package stemmer

import "strings"

var (
	enStopWords = map[string]struct{}{
		"a":          {},
		"about":      {},
		"above":      {},
		"after":      {},
		"again":      {},
		"against":    {},
		"all":        {},
		"am":         {},
		"an":         {},
		"and":        {},
		"any":        {},
		"are":        {},
		"as":         {},
		"at":         {},
		"be":         {},
		"because":    {},
		"been":       {},
		"before":     {},
		"being":      {},
		"below":      {},
		"between":    {},
		"both":       {},
		"but":        {},
		"by":         {},
		"can":        {},
		"did":        {},
		"do":         {},
		"does":       {},
		"doing":      {},
		"don":        {},
		"down":       {},
		"during":     {},
		"each":       {},
		"few":        {},
		"for":        {},
		"from":       {},
		"further":    {},
		"had":        {},
		"has":        {},
		"have":       {},
		"having":     {},
		"he":         {},
		"her":        {},
		"here":       {},
		"hers":       {},
		"herself":    {},
		"him":        {},
		"himself":    {},
		"his":        {},
		"how":        {},
		"i":          {},
		"if":         {},
		"in":         {},
		"into":       {},
		"is":         {},
		"it":         {},
		"its":        {},
		"itself":     {},
		"just":       {},
		"me":         {},
		"more":       {},
		"most":       {},
		"my":         {},
		"myself":     {},
		"no":         {},
		"nor":        {},
		"not":        {},
		"now":        {},
		"of":         {},
		"off":        {},
		"on":         {},
		"once":       {},
		"only":       {},
		"or":         {},
		"other":      {},
		"our":        {},
		"ours":       {},
		"ourselves":  {},
		"out":        {},
		"over":       {},
		"own":        {},
		"s":          {},
		"same":       {},
		"she":        {},
		"should":     {},
		"so":         {},
		"some":       {},
		"such":       {},
		"t":          {},
		"than":       {},
		"that":       {},
		"the":        {},
		"their":      {},
		"theirs":     {},
		"them":       {},
		"themselves": {},
		"then":       {},
		"there":      {},
		"these":      {},
		"they":       {},
		"this":       {},
		"those":      {},
		"through":    {},
		"to":         {},
		"too":        {},
		"under":      {},
		"until":      {},
		"up":         {},
		"very":       {},
		"was":        {},
		"we":         {},
		"were":       {},
		"what":       {},
		"when":       {},
		"where":      {},
		"which":      {},
		"while":      {},
		"who":        {},
		"whom":       {},
		"why":        {},
		"will":       {},
		"with":       {},
		"you":        {},
		"your":       {},
		"yours":      {},
		"yourself":   {},
		"yourselves": {},
	}

	enSpecialWords = map[string]string{
		"skis":       "ski",
		"skies":      "sky",
		"dying":      "die",
		"lying":      "lie",
		"tying":      "tie",
		"idly":       "idl",
		"gently":     "gentl",
		"ugly":       "ugli",
		"early":      "earli",
		"only":       "onli",
		"singly":     "singl",
		"sky":        "sky",
		"news":       "news",
		"howe":       "howe",
		"atlas":      "atlas",
		"cosmos":     "cosmos",
		"bias":       "bias",
		"andes":      "andes",
		"inning":     "inning",
		"innings":    "inning",
		"outing":     "outing",
		"outings":    "outing",
		"canning":    "canning",
		"cannings":   "canning",
		"herring":    "herring",
		"herrings":   "herring",
		"earring":    "earring",
		"earrings":   "earring",
		"proceed":    "proceed",
		"proceeds":   "proceed",
		"proceeded":  "proceed",
		"proceeding": "proceed",
		"exceed":     "exceed",
		"exceeds":    "exceed",
		"exceeded":   "exceed",
		"exceeding":  "exceed",
		"succeed":    "succeed",
		"succeeds":   "succeed",
		"succeeded":  "succeed",
		"succeeding": "succeed",
	}
)

type EnglishStemmer struct{}

func NewEnglishStemmer() *EnglishStemmer {
	return &EnglishStemmer{}
}

func (s *EnglishStemmer) Stem(word string) string {
	word = strings.ToLower(word)
	if s.isStopWord(word) || len(word) < 3 {
		return word
	}

	if special, ok := enSpecialWords[word]; ok {
		return special
	}

	word = s.normalizeApostrophes(word)

	return word
}

func (s EnglishStemmer) isStopWord(word string) bool {
	_, ok := enStopWords[word]
	return ok
}

func (s EnglishStemmer) normalizeApostrophes(word string) string {
	var builder strings.Builder
	builder.Grow(len(word))

	for _, r := range word {
		switch r {
		case '\u2019', '\u2018', '\u201B':
			builder.WriteRune('\x27')
		default:
			builder.WriteRune(r)
		}
	}

	nw := builder.String()

	if strings.HasPrefix(nw, "\x27") {
		return nw[1:]
	}

	return nw
}
