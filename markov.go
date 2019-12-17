package main

import (
	"github.com/mb-14/gomarkov"
	"strings"
)

func NewChain() *gomarkov.Chain {
	chain := gomarkov.NewChain(2)
	return chain
}

func GenerateTweetText(chain *gomarkov.Chain) string {
	order := chain.Order
	tokens := make([]string, 0)
	for i := 0; i < order; i++ {
		tokens = append(tokens, gomarkov.StartToken)
	}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - order):])
		//fmt.Println(next)
		tokens = append(tokens, next)
	}
	return strings.Join(tokens[order:len(tokens)-1], "")
}
