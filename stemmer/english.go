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

	enStep0Suffixes    = []string{"'s'", "'s", "'"}
	enStep1ASuffixes   = []string{"sses", "ied", "ies", "us", "ss", "s"}
	enStep1BSuffixes   = []string{"eedly", "ingly", "edly", "eed", "ing", "ed"}
	enDoubleConsonants = map[string]struct{}{
		"bb": {}, "dd": {}, "ff": {}, "gg": {}, "mm": {}, "nn": {}, "pp": {}, "rr": {}, "tt": {},
	}

	enStep2Suffixes = []string{
		"ization",
		"ational",
		"fulness",
		"ousness",
		"iveness",
		"tional",
		"biliti",
		"lessli",
		"entli",
		"ation",
		"alism",
		"aliti",
		"ousli",
		"iviti",
		"fulli",
		"enci",
		"anci",
		"abli",
		"izer",
		"ator",
		"alli",
		"bli",
		"ogi",
		"li",
	}

	enLiEnding = "cdeghkmnrt" // Letters that can precede 'li' suffix

	enStep3Suffixes = []string{
		"ational",
		"tional",
		"alize",
		"icate",
		"iciti",
		"ative",
		"ical",
		"ness",
		"ful",
	}

	enStep4Suffixes = []string{
		"ement",
		"ance",
		"ence",
		"able",
		"ible",
		"ment",
		"ant",
		"ent",
		"ism",
		"ate",
		"iti",
		"ous",
		"ive",
		"ize",
		"ion",
		"al",
		"er",
		"ic",
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

	if strings.HasPrefix(word, "y") {
		word = "Y" + word[1:]
	}

	word = s.replaceYAfterVowel(word)

	r1, r2 := "", ""

	if strings.HasPrefix(word, "gener") || strings.HasPrefix(word, "commun") || strings.HasPrefix(word, "arsen") {
		if strings.HasPrefix(word, "gener") || strings.HasPrefix(word, "arsen") {
			r1 = word[5:]
		} else {
			r1 = word[6:]
		}

		for i := 1; i < len(r1); i++ {
			if !s.isVowel(rune(r1[i])) && s.isVowel(rune(r1[i-1])) {
				r2 = r1[i+1:]
				break
			}
		}
	} else {
		r1, r2 = s.r1r2Standard(word)
	}

	word, r1, r2 = s.step0(word, r1, r2)
	word, r1, r2 = s.step1a(word, r1, r2)
	word, r1, r2 = s.step1b(word, r1, r2)
	word, r1, r2 = s.step1c(word, r1, r2)
	word, r1, r2 = s.step2(word, r1, r2)
	word, r1, r2 = s.step3(word, r1, r2)
	word, r1, r2 = s.step4(word, r1, r2)
	word = s.step5(word, r1, r2)

	word = strings.ReplaceAll(word, "Y", "y")

	return word
}

func (s EnglishStemmer) step0(word, r1, r2 string) (string, string, string) {
	for _, suffix := range enStep0Suffixes {
		if strings.HasSuffix(word, suffix) {
			suffixLen := len(suffix)
			wordLen := len(word)
			r1Len := len(r1)
			r2Len := len(r2)

			if wordLen >= suffixLen {
				word = word[:wordLen-suffixLen]
			}
			if r1Len >= suffixLen {
				r1 = r1[:r1Len-suffixLen]
			}
			if r2Len >= suffixLen {
				r2 = r2[:r2Len-suffixLen]
			}
			break
		}
	}

	return word, r1, r2
}

func (s EnglishStemmer) step1a(word, r1, r2 string) (string, string, string) {
	step1aVowelFound := false

	for _, suffix := range enStep1ASuffixes {
		if strings.HasSuffix(word, suffix) {
			switch suffix {
			case "sses":
				if len(word) >= 2 {
					word = word[:len(word)-2]
				}
				if len(r1) >= 2 {
					r1 = r1[:len(r1)-2]
				}
				if len(r2) >= 2 {
					r2 = r2[:len(r2)-2]
				}

			case "ied", "ies":
				if len(word[:len(word)-len(suffix)]) > 1 {
					if len(word) >= 2 {
						word = word[:len(word)-2]
					}
					if len(r1) >= 2 {
						r1 = r1[:len(r1)-2]
					}
					if len(r2) >= 2 {
						r2 = r2[:len(r2)-2]
					}
				} else {
					if len(word) >= 1 {
						word = word[:len(word)-1]
					}
					if len(r1) >= 1 {
						r1 = r1[:len(r1)-1]
					}
					if len(r2) >= 1 {
						r2 = r2[:len(r2)-1]
					}
				}

			case "s":
				if len(word) > 1 {
					for _, letter := range word[:len(word)-2] {
						if s.isVowel(letter) {
							step1aVowelFound = true
							break
						}
					}
				}
				if step1aVowelFound {
					if len(word) >= 1 {
						word = word[:len(word)-1]
					}
					if len(r1) >= 1 {
						r1 = r1[:len(r1)-1]
					}
					if len(r2) >= 1 {
						r2 = r2[:len(r2)-1]
					}
				}
			}
			break
		}
	}

	return word, r1, r2
}

func (s EnglishStemmer) step1b(word, r1, r2 string) (string, string, string) {
	step1bVowelFound := false

	for _, suffix := range enStep1BSuffixes {
		if strings.HasSuffix(word, suffix) {
			if suffix == "eed" || suffix == "eedly" {
				if strings.HasSuffix(r1, suffix) {
					word = suffixReplace(word, suffix, "ee")

					if len(r1) >= len(suffix) {
						r1 = suffixReplace(r1, suffix, "ee")
					} else {
						r1 = ""
					}

					if len(r2) >= len(suffix) {
						r2 = suffixReplace(r2, suffix, "ee")
					} else {
						r2 = ""
					}
				}
			} else {
				for _, letter := range word[:len(word)-len(suffix)] {
					if s.isVowel(letter) {
						step1bVowelFound = true
						break
					}
				}

				if step1bVowelFound {
					word = word[:len(word)-len(suffix)]
					if len(r1) >= len(suffix) {
						r1 = r1[:len(r1)-len(suffix)]
					} else {
						r1 = ""
					}
					if len(r2) >= len(suffix) {
						r2 = r2[:len(r2)-len(suffix)]
					} else {
						r2 = ""
					}

					if strings.HasSuffix(word, "at") || strings.HasSuffix(word, "bl") || strings.HasSuffix(word, "iz") {
						word += "e"
						r1 += "e"

						if len(word) > 5 || len(r1) >= 3 {
							r2 += "e"
						}
					} else if s.hasDoubleConstantSuffix(word) {
						if len(word) > 1 {
							word = word[:len(word)-1]
						}
						if len(r1) > 1 {
							r1 = r1[:len(r1)-1]
						}
						if len(r2) > 1 {
							r2 = r2[:len(r2)-1]
						}
					} else if (r1 == "" && len(word) >= 3 && !s.isVowel(rune(word[len(word)-1])) &&
						!strings.ContainsRune("wxY", rune(word[len(word)-1])) &&
						s.isVowel(rune(word[len(word)-2])) &&
						!s.isVowel(rune(word[len(word)-3]))) ||
						(r1 == "" && len(word) == 2 && s.isVowel(rune(word[0])) && !s.isVowel(rune(word[1]))) {
						word += "e"

						if len(r1) > 0 {
							r1 += "e"
						}

						if len(r2) > 0 {
							r2 += "e"
						}
					}
				}
			}
			break
		}
	}

	return word, r1, r2
}

func (s EnglishStemmer) step1c(word, r1, r2 string) (string, string, string) {
	if len(word) > 2 && (word[len(word)-1] == 'y' || word[len(word)-1] == 'Y') && !s.isVowel(rune(word[len(word)-2])) {
		word = word[:len(word)-1] + "i"
		if len(r1) >= 1 {
			r1 = r1[:len(r1)-1] + "i"
		} else {
			r1 = ""
		}
		if len(r2) >= 1 {
			r2 = r2[:len(r2)-1] + "i"
		} else {
			r2 = ""
		}
	}
	return word, r1, r2
}

func (s EnglishStemmer) step2(word, r1, r2 string) (string, string, string) {
	for _, suffix := range enStep2Suffixes {
		if strings.HasSuffix(word, suffix) {
			if strings.HasSuffix(r1, suffix) {
				switch suffix {
				case "tional":
					word = word[:len(word)-2]
					r1 = r1[:len(r1)-2]
					r2 = r2[:len(r2)-2]

				case "enci", "anci", "abli":
					word = word[:len(word)-1] + "e"

					if len(r1) >= 1 {
						r1 = r1[:len(r1)-1] + "e"
					} else {
						r1 = ""
					}

					if len(r2) >= 1 {
						r2 = r2[:len(r2)-1] + "e"
					} else {
						r2 = ""
					}

				case "entli":
					word = word[:len(word)-2]
					r1 = r1[:len(r1)-2]
					r2 = r2[:len(r2)-2]

				case "izer", "ization":
					word = suffixReplace(word, suffix, "ize")

					if len(r1) >= len(suffix) {
						r1 = suffixReplace(r1, suffix, "ize")
					} else {
						r1 = ""
					}

					if len(r2) >= len(suffix) {
						r2 = suffixReplace(r2, suffix, "ize")
					} else {
						r2 = ""
					}

				case "ational", "ation", "ator":
					word = suffixReplace(word, suffix, "ate")

					if len(r1) >= len(suffix) {
						r1 = suffixReplace(r1, suffix, "ate")
					} else {
						r1 = ""
					}

					if len(r2) >= len(suffix) {
						r2 = suffixReplace(r2, suffix, "ate")
					} else {
						r2 = "e"
					}

				case "alism", "aliti", "alli":
					word = suffixReplace(word, suffix, "al")

					if len(r1) >= len(suffix) {
						r1 = suffixReplace(r1, suffix, "al")
					} else {
						r1 = ""
					}

					if len(r2) >= len(suffix) {
						r2 = suffixReplace(r2, suffix, "al")
					} else {
						r2 = ""
					}

				case "fulness":
					word = word[:len(word)-4]
					r1 = r1[:len(r1)-4]
					r2 = r2[:len(r2)-4]

				case "ousli", "ousness":
					word = suffixReplace(word, suffix, "ous")

					if len(r1) >= len(suffix) {
						r1 = suffixReplace(r1, suffix, "ous")
					} else {
						r1 = ""
					}

					if len(r2) >= len(suffix) {
						r2 = suffixReplace(r2, suffix, "ous")
					} else {
						r2 = ""
					}

				case "iveness", "iviti":
					word = suffixReplace(word, suffix, "ive")

					if len(r1) >= len(suffix) {
						r1 = suffixReplace(r1, suffix, "ive")
					} else {
						r1 = ""
					}

					if len(r2) >= len(suffix) {
						r2 = suffixReplace(r2, suffix, "ive")
					} else {
						r2 = "e"
					}

				case "biliti", "bli":
					word = suffixReplace(word, suffix, "ble")

					if len(r1) >= len(suffix) {
						r1 = suffixReplace(r1, suffix, "ble")
					} else {
						r1 = ""
					}

					if len(r2) >= len(suffix) {
						r2 = suffixReplace(r2, suffix, "ble")
					} else {
						r2 = ""
					}

				case "ogi":
					if len(word) > 3 && word[len(word)-4] == 'l' {
						word = word[:len(word)-1]
						r1 = r1[:len(r1)-1]
						r2 = r2[:len(r2)-1]
					}

				case "fulli", "lessli":
					word = word[:len(word)-2]
					r1 = r1[:len(r1)-2]
					r2 = r2[:len(r2)-2]

				case "li":
					if len(word) > 2 && strings.ContainsRune(enLiEnding, rune(word[len(word)-3])) {
						word = word[:len(word)-2]
						r1 = r1[:len(r1)-2]
						r2 = r2[:len(r2)-2]
					}
				}
			}
			break
		}
	}
	return word, r1, r2
}

func (s EnglishStemmer) step3(word, r1, r2 string) (string, string, string) {
	for _, suffix := range enStep3Suffixes {
		if strings.HasSuffix(word, suffix) {
			if strings.HasSuffix(r1, suffix) {
				switch suffix {
				case "tional":
					word = word[:len(word)-2]
					r1 = r1[:len(r1)-2]
					r2 = r2[:len(r2)-2]

				case "ational":
					word = suffixReplace(word, suffix, "ate")
					if len(r1) >= len(suffix) {
						r1 = suffixReplace(r1, suffix, "ate")
					} else {
						r1 = ""
					}
					if len(r2) >= len(suffix) {
						r2 = suffixReplace(r2, suffix, "ate")
					} else {
						r2 = ""
					}

				case "alize":
					word = word[:len(word)-3]
					r1 = r1[:len(r1)-3]
					r2 = r2[:len(r2)-3]

				case "icate", "iciti", "ical":
					word = suffixReplace(word, suffix, "ic")
					if len(r1) >= len(suffix) {
						r1 = suffixReplace(r1, suffix, "ic")
					} else {
						r1 = ""
					}
					if len(r2) >= len(suffix) {
						r2 = suffixReplace(r2, suffix, "ic")
					} else {
						r2 = ""
					}

				case "ful", "ness":
					word = word[:len(word)-len(suffix)]
					r1 = r1[:len(r1)-len(suffix)]
					r2 = r2[:len(r2)-len(suffix)]

				case "ative":
					if strings.HasSuffix(r2, suffix) {
						word = word[:len(word)-5]
						r1 = r1[:len(r1)-5]
						r2 = r2[:len(r2)-5]
					}
				}
			}
			break
		}
	}
	return word, r1, r2
}

func (s EnglishStemmer) step4(word, r1, r2 string) (string, string, string) {
	for _, suffix := range enStep4Suffixes {
		if strings.HasSuffix(word, suffix) {
			if strings.HasSuffix(r2, suffix) {
				if suffix == "ion" {
					if len(word) > 3 && (word[len(word)-4] == 's' || word[len(word)-4] == 't') {
						word = word[:len(word)-3]
						r1 = r1[:len(r1)-3]
						r2 = r2[:len(r2)-3]
					}
				} else {
					word = word[:len(word)-len(suffix)]
					r1 = r1[:len(r1)-len(suffix)]
					r2 = r2[:len(r2)-len(suffix)]
				}
			}
			break
		}
	}
	return word, r1, r2
}

func (s EnglishStemmer) step5(word, r1, r2 string) string {
	if strings.HasSuffix(r2, "l") && len(word) > 1 && word[len(word)-2] == 'l' {
		word = word[:len(word)-1]
	} else if strings.HasSuffix(r2, "e") {
		word = word[:len(word)-1]
	} else if strings.HasSuffix(r1, "e") {
		if len(word) >= 4 && (s.isVowel(rune(word[len(word)-2])) ||
			strings.ContainsRune("wxY", rune(word[len(word)-2])) ||
			!s.isVowel(rune(word[len(word)-3])) ||
			s.isVowel(rune(word[len(word)-4]))) {
			word = word[:len(word)-1]
		}
	}
	return word
}

func (s EnglishStemmer) hasDoubleConstantSuffix(word string) bool {
	if len(word) < 2 {
		return false
	}

	// Check the last two characters of the word
	suffix := word[len(word)-2:]

	_, exists := enDoubleConsonants[suffix]
	return exists
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

func (s EnglishStemmer) isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

func (s EnglishStemmer) replaceYAfterVowel(word string) string {
	var builder strings.Builder
	builder.Grow(len(word))

	for i, r := range word {
		if r == 'y' && i > 0 && s.isVowel(rune(word[i-1])) {
			builder.WriteRune('Y')
		} else {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}

func (s EnglishStemmer) r1r2Standard(word string) (string, string) {
	var r1Start, r2Start int
	wordRunes := []rune(word)
	wordLen := len(wordRunes)

	// Initialize R1 and R2 to the end of the word
	r1Start, r2Start = wordLen, wordLen

	// Find R1 and R2 in a single pass
	for i := 1; i < wordLen; i++ {
		if r1Start == wordLen && !s.isVowel(wordRunes[i]) && s.isVowel(wordRunes[i-1]) {
			r1Start = i + 1
		} else if r1Start < wordLen && r2Start == wordLen && !s.isVowel(wordRunes[i]) && s.isVowel(wordRunes[i-1]) {
			r2Start = i + 1
			break // No need to continue after finding R2
		}
	}

	r1 := ""
	if r1Start < wordLen {
		r1 = string(wordRunes[r1Start:])
	}

	r2 := ""
	if r2Start < wordLen {
		r2 = string(wordRunes[r2Start:])
	}

	return r1, r2
}

func suffixReplace(word, suffix, replacement string) string {
	suffixLen := len(suffix)
	wordLen := len(word)

	if wordLen < suffixLen || word[wordLen-suffixLen:] != suffix {
		return word
	}

	return word[:wordLen-suffixLen] + replacement
}
