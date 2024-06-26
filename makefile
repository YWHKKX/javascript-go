all:
	antlr4 -Dlanguage=Go JavaScript/JavaScriptLexer.g4 JavaScript/JavaScriptParser.g4
	go build main.go