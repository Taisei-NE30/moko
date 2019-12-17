package main

import (
	"github.com/mb-14/gomarkov"
	"strings"
)

func NewChain(texts []string) *gomarkov.Chain {
	chain := gomarkov.NewChain(3)

	for _, text := range texts {
		chain.Add(Tokenize(text))
	}
	return chain
}

func generateTweetText(chain *gomarkov.Chain) string {
	order := chain.Order
	tokens := make([]string, 0)
	for i := 0; i < order; i++ {
		tokens = append(tokens, gomarkov.StartToken)
	}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[len(tokens)-order:])
		tokens = append(tokens, next)
	}
	return strings.Join(tokens[order:len(tokens)-1], "")
}
