package main

import "github.com/ikawaha/kagome/tokenizer"

var t = tokenizer.New()

func Tokenize(text string) []string {
	tokens := t.Tokenize(text)
	var returnTokens []string
	for _, token := range tokens {
		// tokenizer.DUMMY: BOS: Begin Of Sentence, EOS: End Of Sentence.
		if token.Class != tokenizer.DUMMY {
			returnTokens = append(returnTokens, token.Surface)
		}
	}
	return returnTokens
}
