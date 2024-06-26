package main

import (
	"fmt"
	"os"

	Jsp "github.com/klang/js-go/JavaScript"
	Lis "github.com/klang/js-go/Listener"

	"github.com/antlr4-go/antlr/v4"
)

func readSource(filename string) (string, error) {
	d, err := os.ReadFile(filename)
	return string(d), err
}

func Parser(input string) {
	charStream := antlr.NewInputStream(input)
	lexer := Jsp.NewJavaScriptLexer(charStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := Jsp.NewJavaScriptParser(stream)
	tree := parser.Program()
	treeString := tree.ToStringTree(nil, parser)
	fmt.Printf("%v\n\n", treeString)

	listener := &Lis.MyListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}

func main() {
	code, err := readSource("./hello.js")
	fmt.Printf("%v\n\n", code)
	if err != nil {
		return
	}

	Parser(code)
}
