package hazm

import (
	"strings"
	"testing"
)

func TestTokenizerBasic(t *testing.T) {
	tokenizer, err := NewWordTokenizer(DefaultTokenizerOptions())
	if err != nil {
		t.Fatalf("failed to build tokenizer: %v", err)
	}

	got := tokenizer.Tokenize("این جمله (خیلی) پیچیده نیست!!!")
	want := []string{"این", "جمله", "(", "خیلی", ")", "پیچیده", "نیست", "!!!"}
	if strings.Join(got, "|") != strings.Join(want, "|") {
		t.Fatalf("unexpected tokens: got=%v want=%v", got, want)
	}
}

func TestTokenizerJoinedVerb(t *testing.T) {
	tokenizer, err := NewWordTokenizer(DefaultTokenizerOptions())
	if err != nil {
		t.Fatalf("failed to build tokenizer: %v", err)
	}

	got := tokenizer.Tokenize("گفته شده است")
	want := []string{"گفته_شده_است"}
	if strings.Join(got, "|") != strings.Join(want, "|") {
		t.Fatalf("unexpected tokens: got=%v want=%v", got, want)
	}
}

func TestNormalizerNormalize(t *testing.T) {
	normalizer, err := NewNormalizer(DefaultNormalizerOptions())
	if err != nil {
		t.Fatalf("failed to build normalizer: %v", err)
	}

	got := normalizer.Normalize(`"سلام   دنیا"`)
	if got != "«سلام دنیا»" {
		t.Fatalf("unexpected normalized text: %q", got)
	}
}

func TestLemmatizer(t *testing.T) {
	lemmatizer, err := NewLemmatizer(DefaultLemmatizerOptions())
	if err != nil {
		t.Fatalf("failed to build lemmatizer: %v", err)
	}

	cases := map[string]string{
		"کتاب‌ها":       "کتاب",
		"می‌روم":        "رفت#رو",
		"گفته_شده_است":  "گفت#گو",
		"نچشیده_است":    "چشید#چش",
	}

	for word, expected := range cases {
		if got := lemmatizer.Lemmatize(word, ""); got != expected {
			t.Fatalf("unexpected lemma for %q: got=%q want=%q", word, got, expected)
		}
	}
}
