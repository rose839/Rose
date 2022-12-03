package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/rose839/Rose/grammar"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3 /* hello */")

	// Create the Lexer
	lexer := grammar.NewRoseLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
