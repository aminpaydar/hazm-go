package hazm

import (
	"sort"
	"strings"
	"unicode/utf8"
)

type Stemmer struct {
	ends []string
}

func NewStemmer() *Stemmer {
	ends := make([]string, 0, len(suffixes)+3)
	for s := range suffixes {
		ends = append(ends, s)
	}
	ends = append(ends, "ٔ", "\u200cا", "\u200c")
	sort.Slice(ends, func(i, j int) bool {
		return utf8.RuneCountInString(ends[i]) > utf8.RuneCountInString(ends[j])
	})
	return &Stemmer{ends: ends}
}

func (s *Stemmer) Stem(word string) string {
	for _, end := range s.ends {
		if strings.HasSuffix(word, end) {
			if utf8.RuneCountInString(end) == 1 && utf8.RuneCountInString(word)-1 < 3 {
				continue
			}
			word = strings.TrimSuffix(word, end)
			break
		}
	}

	if strings.HasSuffix(word, "ۀ") {
		word = strings.TrimSuffix(word, "ۀ") + "ه"
	}
	if strings.HasSuffix(word, "\u200c") {
		word = strings.TrimSuffix(word, "\u200c")
	}
	return word
}
