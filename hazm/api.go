package hazm

func WordTokenize(text string) ([]string, error) {
	tokenizer, err := NewWordTokenizer(DefaultTokenizerOptions())
	if err != nil {
		return nil, err
	}
	return tokenizer.Tokenize(text), nil
}
