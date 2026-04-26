package hazm

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

type TokenizerOptions struct {
	WordsFile         string
	VerbsFile         string
	JoinVerbParts     bool
	JoinAbbreviations bool
	SeparateEmoji     bool
	ReplaceLinks      bool
	ReplaceIDs        bool
	ReplaceEmails     bool
	ReplaceNumbers    bool
	ReplaceHashtags   bool
}

func DefaultTokenizerOptions() TokenizerOptions {
	return TokenizerOptions{
		JoinVerbParts: true,
	}
}

type WordTokenizer struct {
	options       TokenizerOptions
	words         map[string]WordInfo
	verbs         []string
	bons          map[string]struct{}
	verbE         map[string]struct{}
	beforeVerbs   map[string]struct{}
	afterVerbs    map[string]struct{}
	abbreviations []string

	splitPattern       *regexp.Regexp
	emojiPattern       *regexp.Regexp
	idPattern          *regexp.Regexp
	linkPattern        *regexp.Regexp
	emailPattern       *regexp.Regexp
	numberIntPattern   *regexp.Regexp
	numberFloatPattern *regexp.Regexp
	hashtagPattern     *regexp.Regexp
}

func NewWordTokenizer(opts TokenizerOptions) (*WordTokenizer, error) {
	if opts == (TokenizerOptions{}) {
		opts = DefaultTokenizerOptions()
	}
	if opts.WordsFile == "" {
		opts.WordsFile = DefaultWordsPath()
	}
	if opts.VerbsFile == "" {
		opts.VerbsFile = DefaultVerbsPath()
	}

	words, err := LoadWords(opts.WordsFile)
	if err != nil {
		return nil, err
	}

	t := &WordTokenizer{
		options: opts,
		words:   words,

		bons:  map[string]struct{}{},
		verbE: map[string]struct{}{},
		beforeVerbs: map[string]struct{}{
			"خواهم": {}, "خواهی": {}, "خواهد": {}, "خواهیم": {}, "خواهید": {}, "خواهند": {},
			"نخواهم": {}, "نخواهی": {}, "نخواهد": {}, "نخواهیم": {}, "نخواهید": {}, "نخواهند": {},
		},
		afterVerbs: map[string]struct{}{
			"ام": {}, "ای": {}, "است": {}, "ایم": {}, "اید": {}, "اند": {}, "بودم": {}, "بودی": {}, "بود": {}, "بودیم": {}, "بودید": {}, "بودند": {},
			"باشم": {}, "باشی": {}, "باشد": {}, "باشیم": {}, "باشید": {}, "باشند": {}, "شده_ام": {}, "شده_ای": {}, "شده_است": {},
			"شده_ایم": {}, "شده_اید": {}, "شده_اند": {}, "شده_بودم": {}, "شده_بودی": {}, "شده_بود": {}, "شده_بودیم": {},
			"شده_بودید": {}, "شده_بودند": {}, "شده_باشم": {}, "شده_باشی": {}, "شده_باشد": {}, "شده_باشیم": {}, "شده_باشید": {}, "شده_باشند": {},
			"نشده_ام": {}, "نشده_ای": {}, "نشده_است": {}, "نشده_ایم": {}, "نشده_اید": {}, "نشده_اند": {}, "نشده_بودم": {}, "نشده_بودی": {}, "نشده_بود": {},
			"نشده_بودیم": {}, "نشده_بودید": {}, "نشده_بودند": {}, "نشده_باشم": {}, "نشده_باشی": {}, "نشده_باشد": {}, "نشده_باشیم": {}, "نشده_باشید": {}, "نشده_باشند": {},
			"شوم": {}, "شوی": {}, "شود": {}, "شویم": {}, "شوید": {}, "شوند": {}, "شدم": {}, "شدی": {}, "شد": {}, "شدیم": {}, "شدید": {}, "شدند": {},
			"نشوم": {}, "نشوی": {}, "نشود": {}, "نشویم": {}, "نشوید": {}, "نشوند": {}, "نشدم": {}, "نشدی": {}, "نشد": {}, "نشدیم": {}, "نشدید": {}, "نشدند": {},
			"می‌شوم": {}, "می‌شوی": {}, "می‌شود": {}, "می‌شویم": {}, "می‌شوید": {}, "می‌شوند": {}, "می‌شدم": {}, "می‌شدی": {}, "می‌شد": {}, "می‌شدیم": {}, "می‌شدید": {}, "می‌شدند": {},
			"نمی‌شوم": {}, "نمی‌شوی": {}, "نمی‌شود": {}, "نمی‌شویم": {}, "نمی‌شوید": {}, "نمی‌شوند": {}, "نمی‌شدم": {}, "نمی‌شدی": {}, "نمی‌شد": {}, "نمی‌شدیم": {}, "نمی‌شدید": {}, "نمی‌شدند": {},
			"خواهم_شد": {}, "خواهی_شد": {}, "خواهد_شد": {}, "خواهیم_شد": {}, "خواهید_شد": {}, "خواهند_شد": {},
			"نخواهم_شد": {}, "نخواهی_شد": {}, "نخواهد_شد": {}, "نخواهیم_شد": {}, "نخواهید_شد": {}, "نخواهند_شد": {},
		},

		splitPattern:       regexp.MustCompile(`([؟!?]+|[\d.:]+|[:.،؛»\])}"«\[({/\\])`),
		emojiPattern:       regexp.MustCompile(`[\x{1f300}-\x{1f5ff}\x{1f600}-\x{1f64f}]`),
		idPattern:          regexp.MustCompile(`(@[\w_]+)`),
		linkPattern:        regexp.MustCompile(`((https?|ftp)://)?(([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,})[-\w@:%_.+/~#?=&]*`),
		emailPattern:       regexp.MustCompile(`[a-zA-Z0-9._+-]+@([a-zA-Z0-9-]+\.)+[A-Za-z]{2,}`),
		numberIntPattern:   regexp.MustCompile(`\b[\d۰-۹]+\b`),
		numberFloatPattern: regexp.MustCompile(`\b[\d۰-۹,٬]+[.٫٬][\d۰-۹]+\b`),
		hashtagPattern:     regexp.MustCompile(`#(\S+)`),
	}

	if opts.JoinVerbParts {
		verbs, err := LoadVerbs(opts.VerbsFile)
		if err != nil {
			return nil, err
		}
		for i := len(verbs) - 1; i >= 0; i-- {
			verb := verbs[i]
			t.verbs = append(t.verbs, verb)
			bon := strings.SplitN(verb, "#", 2)[0]
			t.bons[bon] = struct{}{}
		}
		for bon := range t.bons {
			t.verbE[bon+"ه"] = struct{}{}
			t.verbE["ن"+bon+"ه"] = struct{}{}
		}
	}

	if opts.JoinAbbreviations {
		lines, err := LoadLines(DefaultAbbreviationsPath())
		if err != nil {
			return nil, err
		}
		t.abbreviations = lines
	}

	return t, nil
}

func (t *WordTokenizer) Words() map[string]WordInfo {
	return t.words
}

func (t *WordTokenizer) Verbs() []string {
	return t.verbs
}

func (t *WordTokenizer) Tokenize(text string) []string {
	abbrMap := map[string]string{}
	if t.options.JoinAbbreviations {
		rnd := 313
		for strings.Contains(text, fmt.Sprintf("%d", rnd)) {
			rnd++
		}
		for i, abbr := range t.abbreviations {
			placeholder := fmt.Sprintf("%d%d", rnd, i)
			text = strings.ReplaceAll(text, " "+abbr+" ", " "+placeholder+" ")
			abbrMap[placeholder] = abbr
		}
	}

	if t.options.SeparateEmoji {
		text = t.emojiPattern.ReplaceAllString(text, "$0 ")
	}
	if t.options.ReplaceEmails {
		text = t.emailPattern.ReplaceAllString(text, " EMAIL ")
	}
	if t.options.ReplaceLinks {
		text = t.linkPattern.ReplaceAllString(text, " LINK ")
	}
	if t.options.ReplaceIDs {
		text = t.idPattern.ReplaceAllString(text, " ID ")
	}
	if t.options.ReplaceHashtags {
		text = t.hashtagPattern.ReplaceAllStringFunc(text, func(m string) string {
			tag := strings.TrimPrefix(m, "#")
			return "TAG " + strings.ReplaceAll(tag, "_", " ")
		})
	}
	if t.options.ReplaceNumbers {
		text = t.numberFloatPattern.ReplaceAllString(text, " NUMF ")
		text = t.numberIntPattern.ReplaceAllStringFunc(text, func(m string) string {
			return fmt.Sprintf(" NUM%d ", utf8.RuneCountInString(m))
		})
	}

	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")
	text = t.splitPattern.ReplaceAllString(text, " $1 ")

	tokens := []string{}
	for _, token := range strings.Split(text, " ") {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}
		if original, ok := abbrMap[token]; ok {
			token = original
		}
		tokens = append(tokens, token)
	}

	if t.options.JoinVerbParts {
		return t.JoinVerbParts(tokens)
	}

	return tokens
}

func (t *WordTokenizer) JoinVerbParts(tokens []string) []string {
	if len(tokens) <= 1 {
		return tokens
	}

	result := []string{""}
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		last := result[len(result)-1]

		_, inBefore := t.beforeVerbs[token]
		_, lastInAfter := t.afterVerbs[last]
		_, inVerbE := t.verbE[token]
		if inBefore || (lastInAfter && inVerbE) {
			result[len(result)-1] = token + "_" + last
		} else {
			result = append(result, token)
		}
	}

	out := make([]string, 0, len(result)-1)
	for i := len(result) - 1; i >= 1; i-- {
		out = append(out, result[i])
	}
	return out
}
