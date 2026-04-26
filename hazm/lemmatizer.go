package hazm

import "strings"

type LemmatizerOptions struct {
	WordsFile         string
	VerbsFile         string
	JoinedVerbParts   bool
	GenerateVerbForms bool
}

func DefaultLemmatizerOptions() LemmatizerOptions {
	return LemmatizerOptions{
		JoinedVerbParts:   true,
		GenerateVerbForms: true,
	}
}

type Lemmatizer struct {
	words   map[string]WordInfo
	verbs   map[string]string
	stemmer *Stemmer
}

func NewLemmatizer(opts LemmatizerOptions) (*Lemmatizer, error) {
	if opts == (LemmatizerOptions{}) {
		opts = DefaultLemmatizerOptions()
	}
	tokenizer, err := NewWordTokenizer(TokenizerOptions{
		WordsFile:     opts.WordsFile,
		VerbsFile:     opts.VerbsFile,
		JoinVerbParts: true,
	})
	if err != nil {
		return nil, err
	}

	l := &Lemmatizer{
		words:   tokenizer.Words(),
		verbs:   map[string]string{"است": "#است"},
		stemmer: NewStemmer(),
	}

	var conj Conjugation
	if opts.GenerateVerbForms {
		for _, verb := range tokenizer.Verbs() {
			for _, tense := range conj.GetAll(verb) {
				l.verbs[tense] = verb
			}
		}
	}

	if opts.JoinedVerbParts {
		for _, verb := range tokenizer.Verbs() {
			bon := strings.SplitN(verb, "#", 2)[0]
			for after := range tokenizer.afterVerbs {
				l.verbs[bon+"ه_"+after] = verb
				l.verbs["ن"+bon+"ه_"+after] = verb
			}
			for before := range tokenizer.beforeVerbs {
				l.verbs[before+"_"+bon] = verb
			}
		}
	}

	return l, nil
}

func (l *Lemmatizer) Lemmatize(word, pos string) string {
	if pos == "" {
		if _, ok := l.words[word]; ok {
			return word
		}
	}

	if pos == "" || pos == "VERB" {
		if v, ok := l.verbs[word]; ok {
			return v
		}
	}

	if strings.HasPrefix(pos, "ADJ") && strings.HasSuffix(word, "ی") {
		return word
	}
	if pos == "PRON" {
		return word
	}
	if _, ok := l.words[word]; ok {
		return word
	}

	stem := l.stemmer.Stem(word)
	if _, ok := l.words[stem]; ok {
		return stem
	}
	return word
}
