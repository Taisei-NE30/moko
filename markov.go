package main

import (
	"github.com/mb-14/gomarkov"
)

func NewChain(texts []string) *gomarkov.Chain {
	chain := gomarkov.NewChain(3)

	for _, text := range texts {
		chain.Add(Tokenize(text))
	}
	return chain
}
