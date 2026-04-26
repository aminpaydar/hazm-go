package hazm

import (
	"regexp"
	"strings"
	"unicode"
)

type NormalizerOptions struct {
	CorrectSpacing       bool
	RemoveDiacritics     bool
	RemoveSpecialsChars  bool
	DecreaseRepeatedChar bool
	PersianStyle         bool
	PersianNumbers       bool
	UnicodesReplacement  bool
	SeperateMi           bool
	WordsFile            string
	VerbsFile            string
}

func DefaultNormalizerOptions() NormalizerOptions {
	return NormalizerOptions{
		CorrectSpacing:       true,
		RemoveDiacritics:     true,
		RemoveSpecialsChars:  true,
		DecreaseRepeatedChar: true,
		PersianStyle:         true,
		PersianNumbers:       true,
		UnicodesReplacement:  true,
		SeperateMi:           true,
	}
}

type regexRule struct {
	re   *regexp.Regexp
	repl string
}

type Normalizer struct {
	options    NormalizerOptions
	tokenizer  *WordTokenizer
	words      map[string]WordInfo
	verbs      map[string]struct{}
	diacritics *regexp.Regexp
}

func NewNormalizer(opts NormalizerOptions) (*Normalizer, error) {
	if opts == (NormalizerOptions{}) {
		opts = DefaultNormalizerOptions()
	}
	tokenizer, err := NewWordTokenizer(TokenizerOptions{
		WordsFile:     opts.WordsFile,
		VerbsFile:     opts.VerbsFile,
		JoinVerbParts: false,
	})
	if err != nil {
		return nil, err
	}

	n := &Normalizer{
		options:    opts,
		tokenizer:  tokenizer,
		words:      tokenizer.Words(),
		verbs:      map[string]struct{}{},
		diacritics: regexp.MustCompile(`[ً-ْ]`),
	}

	if opts.SeperateMi {
		lemmatizer, err := NewLemmatizer(LemmatizerOptions{
			WordsFile:         opts.WordsFile,
			VerbsFile:         opts.VerbsFile,
			JoinedVerbParts:   false,
			GenerateVerbForms: true,
		})
		if err == nil {
			for k := range lemmatizer.verbs {
				n.verbs[k] = struct{}{}
			}
		}
	}

	return n, nil
}

func (n *Normalizer) Normalize(text string) string {
	text = translate(text, translationRunes)

	if n.options.PersianStyle {
		text = n.PersianStyle(text)
	}
	if n.options.PersianNumbers {
		text = n.PersianNumber(text)
	}
	if n.options.RemoveDiacritics {
		text = n.RemoveDiacritics(text)
	}
	if n.options.CorrectSpacing {
		text = n.CorrectSpacing(text)
	}
	if n.options.UnicodesReplacement {
		text = n.UnicodesReplacement(text)
	}
	if n.options.RemoveSpecialsChars {
		text = n.RemoveSpecialChars(text)
	}
	if n.options.DecreaseRepeatedChar {
		text = n.DecreaseRepeatedChars(text)
	}
	if n.options.SeperateMi {
		text = n.SeperateMi(text)
	}
	return text
}

func (n *Normalizer) CorrectSpacing(text string) string {
	text = strings.TrimSpace(text)
	text = regexp.MustCompile(` {2,}`).ReplaceAllString(text, " ")
	text = regexp.MustCompile(`‌{2,}`).ReplaceAllString(text, "‌")
	text = regexp.MustCompile(`[ـ\r]`).ReplaceAllString(text, "")

	lines := strings.Split(text, "\n")
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			out = append(out, line)
			continue
		}
		tokens := n.tokenizer.Tokenize(line)
		tokens = n.TokenSpacing(tokens)
		out = append(out, strings.Join(tokens, " "))
	}
	text = strings.Join(out, "\n")

	spaceBeforePunc := regexp.MustCompile(` ([\.:!،؛؟»\]\)\}])`)
	spaceAfterOpen := regexp.MustCompile(`([«\[\(\{]) `)
	text = spaceBeforePunc.ReplaceAllString(text, `$1`)
	text = spaceAfterOpen.ReplaceAllString(text, `$1`)
	text = regexp.MustCompile(`([^ ])([«\[\(\{])`).ReplaceAllString(text, `$1 $2`)
	text = regexp.MustCompile(`([۰-۹0-9])([آابپتثجچحخدذرزژسشصضطظعغفقکگلمنوهی])`).ReplaceAllString(text, `$1 $2`)
	text = regexp.MustCompile(`([آابپتثجچحخدذرزژسشصضطظعغفقکگلمنوهی])([۰-۹0-9])`).ReplaceAllString(text, `$1 $2`)
	text = regexp.MustCompile(`(^| )(ن?می) `).ReplaceAllString(text, `$1$2‌`)
	text = regexp.MustCompile(`(ه)(ها)`).ReplaceAllString(text, `$1‌$2`)
	return text
}

func (n *Normalizer) RemoveDiacritics(text string) string {
	return n.diacritics.ReplaceAllString(text, "")
}

func (n *Normalizer) RemoveSpecialChars(text string) string {
	return strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Mn, r) {
			return -1
		}
		return r
	}, text)
}

func (n *Normalizer) DecreaseRepeatedChars(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		if _, ok := n.words[w]; ok {
			continue
		}

		noRepeat := reduceRuns(w, 1)
		twoRepeat := reduceRuns(w, 2)
		_, noIn := n.words[noRepeat]
		_, twoIn := n.words[twoRepeat]
		switch {
		case noIn && !twoIn:
			words[i] = noRepeat
		default:
			words[i] = twoRepeat
		}
	}
	return strings.Join(words, " ")
}

func (n *Normalizer) PersianStyle(text string) string {
	text = regexp.MustCompile(`"([^\n"]+)"`).ReplaceAllString(text, "«$1»")
	text = regexp.MustCompile(`([\d+])\.([\d+])`).ReplaceAllString(text, "$1٫$2")
	text = regexp.MustCompile(` ?\.\.\.`).ReplaceAllString(text, " …")
	return text
}

func (n *Normalizer) PersianNumber(text string) string {
	return translate(text, numberRunes)
}

func (n *Normalizer) UnicodesReplacement(text string) string {
	replacements := [][2]string{
		{"﷽", "بسم الله الرحمن الرحیم"},
		{"﷼", "ریال"},
		{"ﷲ", "الله"},
		{"ﷴ", "محمد"},
		{"ﻻ", "لا"},
		{"ﻼ", "لا"},
	}
	for _, r := range replacements {
		text = strings.ReplaceAll(text, r[0], r[1])
	}
	return text
}

func (n *Normalizer) SeperateMi(text string) string {
	pattern := regexp.MustCompile(`\bن?می[آابپتثجچحخدذرزژسشصضطظعغفقکگلمنوهی]+`)
	return pattern.ReplaceAllStringFunc(text, func(m string) string {
		candidate := regexp.MustCompile(`^(ن?می)`).ReplaceAllString(m, `${1}‌`)
		if _, ok := n.verbs[candidate]; ok {
			return candidate
		}
		return m
	})
}

func (n *Normalizer) TokenSpacing(tokens []string) []string {
	result := make([]string, 0, len(tokens))
	for i, token := range tokens {
		joined := false
		if len(result) > 0 {
			pair := result[len(result)-1] + "‌" + token
			if info, ok := n.words[pair]; ok && info.Frequency > 0 {
				joined = true
				if i < len(tokens)-1 {
					if _, ok := n.verbs[token+"_"+tokens[i+1]]; ok {
						joined = false
					}
				}
			} else if _, suffix := suffixes[token]; suffix {
				if _, ok := n.words[result[len(result)-1]]; ok {
					joined = true
				}
			}
			if joined {
				result[len(result)-1] = pair
				continue
			}
		}
		result = append(result, token)
	}
	return result
}

func translate(text string, table map[rune]rune) string {
	var b strings.Builder
	b.Grow(len(text))
	for _, r := range text {
		if v, ok := table[r]; ok {
			b.WriteRune(v)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func reduceRuns(text string, maxRun int) string {
	if maxRun < 1 {
		maxRun = 1
	}
	runes := []rune(text)
	if len(runes) < 3 {
		return text
	}

	var b strings.Builder
	last := rune(0)
	count := 0
	for _, r := range runes {
		if r == last {
			count++
		} else {
			last = r
			count = 1
		}
		if count <= maxRun {
			b.WriteRune(r)
		}
	}
	return b.String()
}
