# hazm-go

Go rewrite of core Hazm text processing components:

- tokenizer (`WordTokenizer`)
- normalizer (`Normalizer`)
- lemmatizer (`Lemmatizer`)

This module reads lexical resources from Hazm's existing data files. By default it looks for:

- `data/*.dat` in the current directory, or
- `../hazm/data/*.dat` (works from this repository root layout).

## Run tests

```bash
go test ./...
```

## Quick usage

```go
tokenizer, _ := hazm.NewWordTokenizer(hazm.DefaultTokenizerOptions())
tokens := tokenizer.Tokenize("گفته شده است")

normalizer, _ := hazm.NewNormalizer(hazm.DefaultNormalizerOptions())
normalized := normalizer.Normalize(`"سلام   دنیا"`)

lemmatizer, _ := hazm.NewLemmatizer(hazm.DefaultLemmatizerOptions())
lemma := lemmatizer.Lemmatize("می‌روم", "")
```
