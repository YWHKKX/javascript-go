// Code generated from JavaScript/JavaScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // JavaScriptParser

import (
	"fmt"
	"strconv"
  	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type JavaScriptParser struct {
	*antlr.BaseParser
}

var JavaScriptParserParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func javascriptparserParserInit() {
  staticData := &JavaScriptParserParserStaticData
  staticData.LiteralNames = []string{
    "", "'function'", "'return'", "'var'", "'class'", "'this'", "'constructor'", 
    "'new'", "'null'", "'undefined'", "'if'", "'else'", "'while'", "'for'", 
    "'break'", "'continue'", "'case'", "'switch'", "'default'", "'true'", 
    "'false'", "'('", "')'", "'{'", "'}'", "'['", "']'", "", "", "", "':'", 
    "'.'", "','", "';'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", "'++'", 
    "'--'", "'+='", "'-='", "'*='", "'/='", "'%='", "'!'", "'=='", "'!='", 
    "'<'", "'>'", "'<='", "'>='", "'==='", "'!=='", "'&&'", "'||'",
  }
  staticData.SymbolicNames = []string{
    "", "FUNCTION", "RETURN", "VAR", "CLASS", "THIS", "CONS", "NEW", "NULL", 
    "UNDEFINED", "IF", "ELSE", "WHILE", "FOR", "BREAK", "CONTINUE", "CASE", 
    "SWITCH", "DEFAULT", "TRUE", "FALSE", "LPAREN", "RPAREN", "LBRACE", 
    "RBRACE", "LBRACKET", "RBRACKET", "NUMBER", "IDENTIFIER", "STRING", 
    "COL", "DOT", "COMMA", "SEMICOLON", "ASSIGN", "ADD", "SUB", "MUL", "DIV", 
    "MOD", "INC", "DEC", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN", 
    "MOD_ASSIGN", "NOT", "EQ", "NEQ", "LT", "GT", "LTE", "GTE", "AEQ", "NAEQ", 
    "AND", "OR", "WS", "COMMENT", "LINE_COMMENT",
  }
  staticData.RuleNames = []string{
    "num", "id", "str", "arr", "key", "obj", "bool", "this", "funcdef", 
    "paramlist", "param", "funcall", "exprlist", "cobj", "class", "cons", 
    "method", "program", "global", "expr", "stmg", "stm", "case", "default", 
    "block",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 60, 319, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2, 
	21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 1, 1, 
	1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 61, 8, 3, 10, 3, 12, 3, 
	64, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 
	5, 76, 8, 5, 10, 5, 12, 5, 79, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 
	1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 91, 8, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 
	9, 5, 9, 99, 8, 9, 10, 9, 12, 9, 102, 9, 9, 1, 10, 1, 10, 1, 10, 3, 10, 
	107, 8, 10, 1, 11, 1, 11, 1, 11, 3, 11, 112, 8, 11, 1, 11, 1, 11, 1, 12, 
	1, 12, 1, 12, 5, 12, 119, 8, 12, 10, 12, 12, 12, 122, 9, 12, 1, 13, 1, 
	13, 1, 13, 1, 13, 3, 13, 128, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 
	1, 14, 3, 14, 136, 8, 14, 1, 14, 5, 14, 139, 8, 14, 10, 14, 12, 14, 142, 
	9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 15, 149, 8, 15, 1, 15, 1, 
	15, 1, 15, 1, 16, 1, 16, 1, 16, 3, 16, 157, 8, 16, 1, 16, 1, 16, 1, 16, 
	1, 17, 5, 17, 163, 8, 17, 10, 17, 12, 17, 166, 9, 17, 1, 18, 1, 18, 1, 
	18, 3, 18, 171, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 
	1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 
	19, 190, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 
	1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 
	19, 5, 19, 211, 8, 19, 10, 19, 12, 19, 214, 9, 19, 1, 20, 1, 20, 3, 20, 
	218, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 
	21, 228, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 
	1, 21, 1, 21, 1, 21, 3, 21, 241, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 
	21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 252, 8, 21, 1, 21, 1, 21, 3, 21, 
	256, 8, 21, 1, 21, 1, 21, 3, 21, 260, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 
	1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 270, 8, 21, 10, 21, 12, 21, 273, 9, 
	21, 1, 21, 3, 21, 276, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 282, 8, 
	21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 289, 8, 22, 5, 22, 291, 8, 
	22, 10, 22, 12, 22, 294, 9, 22, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 300, 
	8, 23, 5, 23, 302, 8, 23, 10, 23, 12, 23, 305, 9, 23, 1, 24, 1, 24, 1, 
	24, 3, 24, 310, 8, 24, 5, 24, 312, 8, 24, 10, 24, 12, 24, 315, 9, 24, 1, 
	24, 1, 24, 1, 24, 0, 1, 38, 25, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 0, 7, 1, 0, 19, 
	20, 1, 0, 40, 41, 2, 0, 36, 36, 47, 47, 1, 0, 35, 38, 1, 0, 48, 55, 1, 
	0, 56, 57, 2, 0, 34, 34, 42, 45, 348, 0, 50, 1, 0, 0, 0, 2, 52, 1, 0, 0, 
	0, 4, 54, 1, 0, 0, 0, 6, 56, 1, 0, 0, 0, 8, 67, 1, 0, 0, 0, 10, 71, 1, 
	0, 0, 0, 12, 82, 1, 0, 0, 0, 14, 84, 1, 0, 0, 0, 16, 86, 1, 0, 0, 0, 18, 
	95, 1, 0, 0, 0, 20, 103, 1, 0, 0, 0, 22, 108, 1, 0, 0, 0, 24, 115, 1, 0, 
	0, 0, 26, 123, 1, 0, 0, 0, 28, 131, 1, 0, 0, 0, 30, 145, 1, 0, 0, 0, 32, 
	153, 1, 0, 0, 0, 34, 164, 1, 0, 0, 0, 36, 170, 1, 0, 0, 0, 38, 189, 1, 
	0, 0, 0, 40, 215, 1, 0, 0, 0, 42, 281, 1, 0, 0, 0, 44, 283, 1, 0, 0, 0, 
	46, 295, 1, 0, 0, 0, 48, 306, 1, 0, 0, 0, 50, 51, 5, 27, 0, 0, 51, 1, 1, 
	0, 0, 0, 52, 53, 5, 28, 0, 0, 53, 3, 1, 0, 0, 0, 54, 55, 5, 29, 0, 0, 55, 
	5, 1, 0, 0, 0, 56, 57, 5, 25, 0, 0, 57, 62, 3, 38, 19, 0, 58, 59, 5, 32, 
	0, 0, 59, 61, 3, 38, 19, 0, 60, 58, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 
	60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 65, 1, 0, 0, 0, 64, 62, 1, 0, 0, 
	0, 65, 66, 5, 26, 0, 0, 66, 7, 1, 0, 0, 0, 67, 68, 3, 2, 1, 0, 68, 69, 
	5, 30, 0, 0, 69, 70, 3, 38, 19, 0, 70, 9, 1, 0, 0, 0, 71, 72, 5, 23, 0, 
	0, 72, 77, 3, 8, 4, 0, 73, 74, 5, 32, 0, 0, 74, 76, 3, 8, 4, 0, 75, 73, 
	1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 
	78, 80, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 81, 5, 24, 0, 0, 81, 11, 1, 
	0, 0, 0, 82, 83, 7, 0, 0, 0, 83, 13, 1, 0, 0, 0, 84, 85, 5, 5, 0, 0, 85, 
	15, 1, 0, 0, 0, 86, 87, 5, 1, 0, 0, 87, 88, 5, 28, 0, 0, 88, 90, 5, 21, 
	0, 0, 89, 91, 3, 18, 9, 0, 90, 89, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 
	92, 1, 0, 0, 0, 92, 93, 5, 22, 0, 0, 93, 94, 3, 48, 24, 0, 94, 17, 1, 0, 
	0, 0, 95, 100, 3, 20, 10, 0, 96, 97, 5, 32, 0, 0, 97, 99, 3, 20, 10, 0, 
	98, 96, 1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 
	1, 0, 0, 0, 101, 19, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 103, 106, 5, 28, 
	0, 0, 104, 105, 5, 34, 0, 0, 105, 107, 3, 38, 19, 0, 106, 104, 1, 0, 0, 
	0, 106, 107, 1, 0, 0, 0, 107, 21, 1, 0, 0, 0, 108, 109, 5, 28, 0, 0, 109, 
	111, 5, 21, 0, 0, 110, 112, 3, 24, 12, 0, 111, 110, 1, 0, 0, 0, 111, 112, 
	1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 5, 22, 0, 0, 114, 23, 1, 0, 
	0, 0, 115, 120, 3, 38, 19, 0, 116, 117, 5, 32, 0, 0, 117, 119, 3, 38, 19, 
	0, 118, 116, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 
	121, 1, 0, 0, 0, 121, 25, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 124, 5, 
	7, 0, 0, 124, 125, 5, 28, 0, 0, 125, 127, 5, 21, 0, 0, 126, 128, 3, 24, 
	12, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 
	129, 130, 5, 22, 0, 0, 130, 27, 1, 0, 0, 0, 131, 132, 5, 4, 0, 0, 132, 
	133, 5, 28, 0, 0, 133, 135, 5, 23, 0, 0, 134, 136, 3, 30, 15, 0, 135, 134, 
	1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 140, 1, 0, 0, 0, 137, 139, 3, 32, 
	16, 0, 138, 137, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 
	140, 141, 1, 0, 0, 0, 141, 143, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 
	144, 5, 24, 0, 0, 144, 29, 1, 0, 0, 0, 145, 146, 5, 6, 0, 0, 146, 148, 
	5, 21, 0, 0, 147, 149, 3, 18, 9, 0, 148, 147, 1, 0, 0, 0, 148, 149, 1, 
	0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 5, 22, 0, 0, 151, 152, 3, 48, 
	24, 0, 152, 31, 1, 0, 0, 0, 153, 154, 5, 28, 0, 0, 154, 156, 5, 21, 0, 
	0, 155, 157, 3, 18, 9, 0, 156, 155, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 
	158, 1, 0, 0, 0, 158, 159, 5, 22, 0, 0, 159, 160, 3, 48, 24, 0, 160, 33, 
	1, 0, 0, 0, 161, 163, 3, 36, 18, 0, 162, 161, 1, 0, 0, 0, 163, 166, 1, 
	0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 35, 1, 0, 0, 
	0, 166, 164, 1, 0, 0, 0, 167, 171, 3, 16, 8, 0, 168, 171, 3, 40, 20, 0, 
	169, 171, 3, 28, 14, 0, 170, 167, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 
	169, 1, 0, 0, 0, 171, 37, 1, 0, 0, 0, 172, 173, 6, 19, -1, 0, 173, 174, 
	5, 21, 0, 0, 174, 175, 3, 38, 19, 0, 175, 176, 5, 22, 0, 0, 176, 190, 1, 
	0, 0, 0, 177, 178, 7, 1, 0, 0, 178, 190, 3, 38, 19, 11, 179, 180, 7, 2, 
	0, 0, 180, 190, 3, 38, 19, 9, 181, 190, 3, 14, 7, 0, 182, 190, 3, 26, 13, 
	0, 183, 190, 3, 22, 11, 0, 184, 190, 3, 0, 0, 0, 185, 190, 3, 2, 1, 0, 
	186, 190, 3, 4, 2, 0, 187, 190, 3, 6, 3, 0, 188, 190, 3, 10, 5, 0, 189, 
	172, 1, 0, 0, 0, 189, 177, 1, 0, 0, 0, 189, 179, 1, 0, 0, 0, 189, 181, 
	1, 0, 0, 0, 189, 182, 1, 0, 0, 0, 189, 183, 1, 0, 0, 0, 189, 184, 1, 0, 
	0, 0, 189, 185, 1, 0, 0, 0, 189, 186, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 
	189, 188, 1, 0, 0, 0, 190, 212, 1, 0, 0, 0, 191, 192, 10, 16, 0, 0, 192, 
	193, 5, 31, 0, 0, 193, 211, 3, 38, 19, 17, 194, 195, 10, 14, 0, 0, 195, 
	196, 7, 3, 0, 0, 196, 211, 3, 38, 19, 15, 197, 198, 10, 13, 0, 0, 198, 
	199, 7, 4, 0, 0, 199, 211, 3, 38, 19, 14, 200, 201, 10, 12, 0, 0, 201, 
	202, 7, 5, 0, 0, 202, 211, 3, 38, 19, 13, 203, 204, 10, 15, 0, 0, 204, 
	205, 5, 25, 0, 0, 205, 206, 3, 38, 19, 0, 206, 207, 5, 26, 0, 0, 207, 211, 
	1, 0, 0, 0, 208, 209, 10, 10, 0, 0, 209, 211, 7, 1, 0, 0, 210, 191, 1, 
	0, 0, 0, 210, 194, 1, 0, 0, 0, 210, 197, 1, 0, 0, 0, 210, 200, 1, 0, 0, 
	0, 210, 203, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212, 
	210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 39, 1, 0, 0, 0, 214, 212, 1, 
	0, 0, 0, 215, 217, 3, 42, 21, 0, 216, 218, 5, 33, 0, 0, 217, 216, 1, 0, 
	0, 0, 217, 218, 1, 0, 0, 0, 218, 41, 1, 0, 0, 0, 219, 220, 3, 38, 19, 0, 
	220, 221, 7, 6, 0, 0, 221, 222, 3, 38, 19, 0, 222, 282, 1, 0, 0, 0, 223, 
	224, 5, 3, 0, 0, 224, 227, 5, 28, 0, 0, 225, 226, 5, 34, 0, 0, 226, 228, 
	3, 38, 19, 0, 227, 225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 282, 1, 
	0, 0, 0, 229, 230, 5, 2, 0, 0, 230, 282, 3, 38, 19, 0, 231, 282, 5, 14, 
	0, 0, 232, 282, 5, 15, 0, 0, 233, 234, 5, 10, 0, 0, 234, 235, 5, 21, 0, 
	0, 235, 236, 3, 38, 19, 0, 236, 237, 5, 22, 0, 0, 237, 240, 3, 48, 24, 
	0, 238, 239, 5, 11, 0, 0, 239, 241, 3, 48, 24, 0, 240, 238, 1, 0, 0, 0, 
	240, 241, 1, 0, 0, 0, 241, 282, 1, 0, 0, 0, 242, 243, 5, 12, 0, 0, 243, 
	244, 5, 21, 0, 0, 244, 245, 3, 38, 19, 0, 245, 246, 5, 22, 0, 0, 246, 247, 
	3, 48, 24, 0, 247, 282, 1, 0, 0, 0, 248, 249, 5, 13, 0, 0, 249, 251, 5, 
	21, 0, 0, 250, 252, 3, 42, 21, 0, 251, 250, 1, 0, 0, 0, 251, 252, 1, 0, 
	0, 0, 252, 253, 1, 0, 0, 0, 253, 255, 5, 33, 0, 0, 254, 256, 3, 42, 21, 
	0, 255, 254, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 
	259, 5, 33, 0, 0, 258, 260, 3, 42, 21, 0, 259, 258, 1, 0, 0, 0, 259, 260, 
	1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 262, 5, 22, 0, 0, 262, 282, 3, 48, 
	24, 0, 263, 264, 5, 17, 0, 0, 264, 265, 5, 21, 0, 0, 265, 266, 3, 38, 19, 
	0, 266, 267, 5, 22, 0, 0, 267, 271, 5, 23, 0, 0, 268, 270, 3, 44, 22, 0, 
	269, 268, 1, 0, 0, 0, 270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 
	272, 1, 0, 0, 0, 272, 275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 276, 
	3, 46, 23, 0, 275, 274, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 1, 
	0, 0, 0, 277, 278, 5, 24, 0, 0, 278, 282, 1, 0, 0, 0, 279, 282, 3, 48, 
	24, 0, 280, 282, 3, 38, 19, 0, 281, 219, 1, 0, 0, 0, 281, 223, 1, 0, 0, 
	0, 281, 229, 1, 0, 0, 0, 281, 231, 1, 0, 0, 0, 281, 232, 1, 0, 0, 0, 281, 
	233, 1, 0, 0, 0, 281, 242, 1, 0, 0, 0, 281, 248, 1, 0, 0, 0, 281, 263, 
	1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 281, 280, 1, 0, 0, 0, 282, 43, 1, 0, 
	0, 0, 283, 284, 5, 16, 0, 0, 284, 285, 3, 38, 19, 0, 285, 292, 5, 30, 0, 
	0, 286, 288, 3, 42, 21, 0, 287, 289, 5, 33, 0, 0, 288, 287, 1, 0, 0, 0, 
	288, 289, 1, 0, 0, 0, 289, 291, 1, 0, 0, 0, 290, 286, 1, 0, 0, 0, 291, 
	294, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 45, 1, 
	0, 0, 0, 294, 292, 1, 0, 0, 0, 295, 296, 5, 18, 0, 0, 296, 303, 5, 30, 
	0, 0, 297, 299, 3, 42, 21, 0, 298, 300, 5, 33, 0, 0, 299, 298, 1, 0, 0, 
	0, 299, 300, 1, 0, 0, 0, 300, 302, 1, 0, 0, 0, 301, 297, 1, 0, 0, 0, 302, 
	305, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 47, 1, 
	0, 0, 0, 305, 303, 1, 0, 0, 0, 306, 313, 5, 23, 0, 0, 307, 309, 3, 42, 
	21, 0, 308, 310, 5, 33, 0, 0, 309, 308, 1, 0, 0, 0, 309, 310, 1, 0, 0, 
	0, 310, 312, 1, 0, 0, 0, 311, 307, 1, 0, 0, 0, 312, 315, 1, 0, 0, 0, 313, 
	311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 316, 1, 0, 0, 0, 315, 313, 
	1, 0, 0, 0, 316, 317, 5, 24, 0, 0, 317, 49, 1, 0, 0, 0, 32, 62, 77, 90, 
	100, 106, 111, 120, 127, 135, 140, 148, 156, 164, 170, 189, 210, 212, 217, 
	227, 240, 251, 255, 259, 271, 275, 281, 288, 292, 299, 303, 309, 313,
}
  deserializer := antlr.NewATNDeserializer(nil)
  staticData.atn = deserializer.Deserialize(staticData.serializedATN)
  atn := staticData.atn
  staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
  decisionToDFA := staticData.decisionToDFA
  for index, state := range atn.DecisionToState {
    decisionToDFA[index] = antlr.NewDFA(state, index)
  }
}

// JavaScriptParserInit initializes any static state used to implement JavaScriptParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJavaScriptParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JavaScriptParserInit() {
  staticData := &JavaScriptParserParserStaticData
  staticData.once.Do(javascriptparserParserInit)
}

// NewJavaScriptParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJavaScriptParser(input antlr.TokenStream) *JavaScriptParser {
	JavaScriptParserInit()
	this := new(JavaScriptParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &JavaScriptParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "JavaScriptParser.g4"

	return this
}


// JavaScriptParser tokens.
const (
	JavaScriptParserEOF = antlr.TokenEOF
	JavaScriptParserFUNCTION = 1
	JavaScriptParserRETURN = 2
	JavaScriptParserVAR = 3
	JavaScriptParserCLASS = 4
	JavaScriptParserTHIS = 5
	JavaScriptParserCONS = 6
	JavaScriptParserNEW = 7
	JavaScriptParserNULL = 8
	JavaScriptParserUNDEFINED = 9
	JavaScriptParserIF = 10
	JavaScriptParserELSE = 11
	JavaScriptParserWHILE = 12
	JavaScriptParserFOR = 13
	JavaScriptParserBREAK = 14
	JavaScriptParserCONTINUE = 15
	JavaScriptParserCASE = 16
	JavaScriptParserSWITCH = 17
	JavaScriptParserDEFAULT = 18
	JavaScriptParserTRUE = 19
	JavaScriptParserFALSE = 20
	JavaScriptParserLPAREN = 21
	JavaScriptParserRPAREN = 22
	JavaScriptParserLBRACE = 23
	JavaScriptParserRBRACE = 24
	JavaScriptParserLBRACKET = 25
	JavaScriptParserRBRACKET = 26
	JavaScriptParserNUMBER = 27
	JavaScriptParserIDENTIFIER = 28
	JavaScriptParserSTRING = 29
	JavaScriptParserCOL = 30
	JavaScriptParserDOT = 31
	JavaScriptParserCOMMA = 32
	JavaScriptParserSEMICOLON = 33
	JavaScriptParserASSIGN = 34
	JavaScriptParserADD = 35
	JavaScriptParserSUB = 36
	JavaScriptParserMUL = 37
	JavaScriptParserDIV = 38
	JavaScriptParserMOD = 39
	JavaScriptParserINC = 40
	JavaScriptParserDEC = 41
	JavaScriptParserADD_ASSIGN = 42
	JavaScriptParserSUB_ASSIGN = 43
	JavaScriptParserMUL_ASSIGN = 44
	JavaScriptParserDIV_ASSIGN = 45
	JavaScriptParserMOD_ASSIGN = 46
	JavaScriptParserNOT = 47
	JavaScriptParserEQ = 48
	JavaScriptParserNEQ = 49
	JavaScriptParserLT = 50
	JavaScriptParserGT = 51
	JavaScriptParserLTE = 52
	JavaScriptParserGTE = 53
	JavaScriptParserAEQ = 54
	JavaScriptParserNAEQ = 55
	JavaScriptParserAND = 56
	JavaScriptParserOR = 57
	JavaScriptParserWS = 58
	JavaScriptParserCOMMENT = 59
	JavaScriptParserLINE_COMMENT = 60
)

// JavaScriptParser rules.
const (
	JavaScriptParserRULE_num = 0
	JavaScriptParserRULE_id = 1
	JavaScriptParserRULE_str = 2
	JavaScriptParserRULE_arr = 3
	JavaScriptParserRULE_key = 4
	JavaScriptParserRULE_obj = 5
	JavaScriptParserRULE_bool = 6
	JavaScriptParserRULE_this = 7
	JavaScriptParserRULE_funcdef = 8
	JavaScriptParserRULE_paramlist = 9
	JavaScriptParserRULE_param = 10
	JavaScriptParserRULE_funcall = 11
	JavaScriptParserRULE_exprlist = 12
	JavaScriptParserRULE_cobj = 13
	JavaScriptParserRULE_class = 14
	JavaScriptParserRULE_cons = 15
	JavaScriptParserRULE_method = 16
	JavaScriptParserRULE_program = 17
	JavaScriptParserRULE_global = 18
	JavaScriptParserRULE_expr = 19
	JavaScriptParserRULE_stmg = 20
	JavaScriptParserRULE_stm = 21
	JavaScriptParserRULE_case = 22
	JavaScriptParserRULE_default = 23
	JavaScriptParserRULE_block = 24
)

// INumContext is an interface to support dynamic dispatch.
type INumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsNumContext differentiates from other interfaces.
	IsNumContext()
}

type NumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumContext() *NumContext {
	var p = new(NumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_num
	return p
}

func InitEmptyNumContext(p *NumContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_num
}

func (*NumContext) IsNumContext() {}

func NewNumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumContext {
	var p = new(NumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_num

	return p
}

func (s *NumContext) GetParser() antlr.Parser { return s.parser }

func (s *NumContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNUMBER, 0)
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterNum(s)
	}
}

func (s *NumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitNum(s)
	}
}




func (p *JavaScriptParser) Num() (localctx INumContext) {
	localctx = NewNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JavaScriptParserRULE_num)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(JavaScriptParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_id
	return p
}

func InitEmptyIdContext(p *IdContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_id
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIDENTIFIER, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitId(s)
	}
}




func (p *JavaScriptParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JavaScriptParserRULE_id)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(JavaScriptParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IStrContext is an interface to support dynamic dispatch.
type IStrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsStrContext differentiates from other interfaces.
	IsStrContext()
}

type StrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrContext() *StrContext {
	var p = new(StrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_str
	return p
}

func InitEmptyStrContext(p *StrContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_str
}

func (*StrContext) IsStrContext() {}

func NewStrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrContext {
	var p = new(StrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_str

	return p
}

func (s *StrContext) GetParser() antlr.Parser { return s.parser }

func (s *StrContext) STRING() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSTRING, 0)
}

func (s *StrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterStr(s)
	}
}

func (s *StrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitStr(s)
	}
}




func (p *JavaScriptParser) Str() (localctx IStrContext) {
	localctx = NewStrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JavaScriptParserRULE_str)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(JavaScriptParserSTRING)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IArrContext is an interface to support dynamic dispatch.
type IArrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	RBRACKET() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArrContext differentiates from other interfaces.
	IsArrContext()
}

type ArrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrContext() *ArrContext {
	var p = new(ArrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_arr
	return p
}

func InitEmptyArrContext(p *ArrContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_arr
}

func (*ArrContext) IsArrContext() {}

func NewArrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrContext {
	var p = new(ArrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_arr

	return p
}

func (s *ArrContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLBRACKET, 0)
}

func (s *ArrContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRBRACKET, 0)
}

func (s *ArrContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserCOMMA)
}

func (s *ArrContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCOMMA, i)
}

func (s *ArrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ArrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterArr(s)
	}
}

func (s *ArrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitArr(s)
	}
}




func (p *JavaScriptParser) Arr() (localctx IArrContext) {
	localctx = NewArrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JavaScriptParserRULE_arr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(JavaScriptParserLBRACKET)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(57)
		p.expr(0)
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == JavaScriptParserCOMMA {
		{
			p.SetState(58)
			p.Match(JavaScriptParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(59)
			p.expr(0)
		}


		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(65)
		p.Match(JavaScriptParserRBRACKET)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id() IIdContext
	COL() antlr.TerminalNode
	Expr() IExprContext

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_key
	return p
}

func InitEmptyKeyContext(p *KeyContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_key
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) Id() IIdContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *KeyContext) COL() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCOL, 0)
}

func (s *KeyContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitKey(s)
	}
}




func (p *JavaScriptParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JavaScriptParserRULE_key)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Id()
	}
	{
		p.SetState(68)
		p.Match(JavaScriptParserCOL)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(69)
		p.expr(0)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IObjContext is an interface to support dynamic dispatch.
type IObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	AllKey() []IKeyContext
	Key(i int) IKeyContext
	RBRACE() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsObjContext differentiates from other interfaces.
	IsObjContext()
}

type ObjContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjContext() *ObjContext {
	var p = new(ObjContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_obj
	return p
}

func InitEmptyObjContext(p *ObjContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_obj
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLBRACE, 0)
}

func (s *ObjContext) AllKey() []IKeyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKeyContext); ok {
			len++
		}
	}

	tst := make([]IKeyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKeyContext); ok {
			tst[i] = t.(IKeyContext)
			i++
		}
	}

	return tst
}

func (s *ObjContext) Key(i int) IKeyContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *ObjContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRBRACE, 0)
}

func (s *ObjContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserCOMMA)
}

func (s *ObjContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCOMMA, i)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterObj(s)
	}
}

func (s *ObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitObj(s)
	}
}




func (p *JavaScriptParser) Obj() (localctx IObjContext) {
	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JavaScriptParserRULE_obj)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(JavaScriptParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(72)
		p.Key()
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == JavaScriptParserCOMMA {
		{
			p.SetState(73)
			p.Match(JavaScriptParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Key()
		}


		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
		p.Match(JavaScriptParserRBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IBoolContext is an interface to support dynamic dispatch.
type IBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsBoolContext differentiates from other interfaces.
	IsBoolContext()
}

type BoolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolContext() *BoolContext {
	var p = new(BoolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_bool
	return p
}

func InitEmptyBoolContext(p *BoolContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_bool
}

func (*BoolContext) IsBoolContext() {}

func NewBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolContext {
	var p = new(BoolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_bool

	return p
}

func (s *BoolContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolContext) TRUE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserTRUE, 0)
}

func (s *BoolContext) FALSE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserFALSE, 0)
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitBool(s)
	}
}




func (p *JavaScriptParser) Bool_() (localctx IBoolContext) {
	localctx = NewBoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JavaScriptParserRULE_bool)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		_la = p.GetTokenStream().LA(1)

		if !(_la == JavaScriptParserTRUE || _la == JavaScriptParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IThisContext is an interface to support dynamic dispatch.
type IThisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	THIS() antlr.TerminalNode

	// IsThisContext differentiates from other interfaces.
	IsThisContext()
}

type ThisContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThisContext() *ThisContext {
	var p = new(ThisContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_this
	return p
}

func InitEmptyThisContext(p *ThisContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_this
}

func (*ThisContext) IsThisContext() {}

func NewThisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThisContext {
	var p = new(ThisContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_this

	return p
}

func (s *ThisContext) GetParser() antlr.Parser { return s.parser }

func (s *ThisContext) THIS() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserTHIS, 0)
}

func (s *ThisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ThisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterThis(s)
	}
}

func (s *ThisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitThis(s)
	}
}




func (p *JavaScriptParser) This() (localctx IThisContext) {
	localctx = NewThisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JavaScriptParserRULE_this)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(JavaScriptParserTHIS)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IFuncdefContext is an interface to support dynamic dispatch.
type IFuncdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	Paramlist() IParamlistContext

	// IsFuncdefContext differentiates from other interfaces.
	IsFuncdefContext()
}

type FuncdefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncdefContext() *FuncdefContext {
	var p = new(FuncdefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_funcdef
	return p
}

func InitEmptyFuncdefContext(p *FuncdefContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_funcdef
}

func (*FuncdefContext) IsFuncdefContext() {}

func NewFuncdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncdefContext {
	var p = new(FuncdefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_funcdef

	return p
}

func (s *FuncdefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncdefContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserFUNCTION, 0)
}

func (s *FuncdefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIDENTIFIER, 0)
}

func (s *FuncdefContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLPAREN, 0)
}

func (s *FuncdefContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRPAREN, 0)
}

func (s *FuncdefContext) Block() IBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncdefContext) Paramlist() IParamlistContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamlistContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamlistContext)
}

func (s *FuncdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FuncdefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterFuncdef(s)
	}
}

func (s *FuncdefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitFuncdef(s)
	}
}




func (p *JavaScriptParser) Funcdef() (localctx IFuncdefContext) {
	localctx = NewFuncdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JavaScriptParserRULE_funcdef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(JavaScriptParserFUNCTION)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Match(JavaScriptParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(JavaScriptParserLPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == JavaScriptParserIDENTIFIER {
		{
			p.SetState(89)
			p.Paramlist()
		}

	}
	{
		p.SetState(92)
		p.Match(JavaScriptParserRPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Block()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IParamlistContext is an interface to support dynamic dispatch.
type IParamlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParamlistContext differentiates from other interfaces.
	IsParamlistContext()
}

type ParamlistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamlistContext() *ParamlistContext {
	var p = new(ParamlistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_paramlist
	return p
}

func InitEmptyParamlistContext(p *ParamlistContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_paramlist
}

func (*ParamlistContext) IsParamlistContext() {}

func NewParamlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamlistContext {
	var p = new(ParamlistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_paramlist

	return p
}

func (s *ParamlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamlistContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamlistContext) Param(i int) IParamContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamlistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserCOMMA)
}

func (s *ParamlistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCOMMA, i)
}

func (s *ParamlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParamlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterParamlist(s)
	}
}

func (s *ParamlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitParamlist(s)
	}
}




func (p *JavaScriptParser) Paramlist() (localctx IParamlistContext) {
	localctx = NewParamlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JavaScriptParserRULE_paramlist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Param()
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == JavaScriptParserCOMMA {
		{
			p.SetState(96)
			p.Match(JavaScriptParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(97)
			p.Param()
		}


		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIDENTIFIER, 0)
}

func (s *ParamContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserASSIGN, 0)
}

func (s *ParamContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitParam(s)
	}
}




func (p *JavaScriptParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JavaScriptParserRULE_param)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(JavaScriptParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == JavaScriptParserASSIGN {
		{
			p.SetState(104)
			p.Match(JavaScriptParserASSIGN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(105)
			p.expr(0)
		}

	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IFuncallContext is an interface to support dynamic dispatch.
type IFuncallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Exprlist() IExprlistContext

	// IsFuncallContext differentiates from other interfaces.
	IsFuncallContext()
}

type FuncallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncallContext() *FuncallContext {
	var p = new(FuncallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_funcall
	return p
}

func InitEmptyFuncallContext(p *FuncallContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_funcall
}

func (*FuncallContext) IsFuncallContext() {}

func NewFuncallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncallContext {
	var p = new(FuncallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_funcall

	return p
}

func (s *FuncallContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIDENTIFIER, 0)
}

func (s *FuncallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLPAREN, 0)
}

func (s *FuncallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRPAREN, 0)
}

func (s *FuncallContext) Exprlist() IExprlistContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprlistContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprlistContext)
}

func (s *FuncallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FuncallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterFuncall(s)
	}
}

func (s *FuncallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitFuncall(s)
	}
}




func (p *JavaScriptParser) Funcall() (localctx IFuncallContext) {
	localctx = NewFuncallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JavaScriptParserRULE_funcall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(JavaScriptParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(JavaScriptParserLPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 144105726279840) != 0) {
		{
			p.SetState(110)
			p.Exprlist()
		}

	}
	{
		p.SetState(113)
		p.Match(JavaScriptParserRPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IExprlistContext is an interface to support dynamic dispatch.
type IExprlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExprlistContext differentiates from other interfaces.
	IsExprlistContext()
}

type ExprlistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprlistContext() *ExprlistContext {
	var p = new(ExprlistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_exprlist
	return p
}

func InitEmptyExprlistContext(p *ExprlistContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_exprlist
}

func (*ExprlistContext) IsExprlistContext() {}

func NewExprlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprlistContext {
	var p = new(ExprlistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_exprlist

	return p
}

func (s *ExprlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprlistContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprlistContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprlistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserCOMMA)
}

func (s *ExprlistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCOMMA, i)
}

func (s *ExprlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterExprlist(s)
	}
}

func (s *ExprlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitExprlist(s)
	}
}




func (p *JavaScriptParser) Exprlist() (localctx IExprlistContext) {
	localctx = NewExprlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JavaScriptParserRULE_exprlist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.expr(0)
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == JavaScriptParserCOMMA {
		{
			p.SetState(116)
			p.Match(JavaScriptParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(117)
			p.expr(0)
		}


		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ICobjContext is an interface to support dynamic dispatch.
type ICobjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NEW() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Exprlist() IExprlistContext

	// IsCobjContext differentiates from other interfaces.
	IsCobjContext()
}

type CobjContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCobjContext() *CobjContext {
	var p = new(CobjContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_cobj
	return p
}

func InitEmptyCobjContext(p *CobjContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_cobj
}

func (*CobjContext) IsCobjContext() {}

func NewCobjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CobjContext {
	var p = new(CobjContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_cobj

	return p
}

func (s *CobjContext) GetParser() antlr.Parser { return s.parser }

func (s *CobjContext) NEW() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNEW, 0)
}

func (s *CobjContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIDENTIFIER, 0)
}

func (s *CobjContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLPAREN, 0)
}

func (s *CobjContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRPAREN, 0)
}

func (s *CobjContext) Exprlist() IExprlistContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprlistContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprlistContext)
}

func (s *CobjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CobjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CobjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterCobj(s)
	}
}

func (s *CobjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitCobj(s)
	}
}




func (p *JavaScriptParser) Cobj() (localctx ICobjContext) {
	localctx = NewCobjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JavaScriptParserRULE_cobj)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(JavaScriptParserNEW)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(JavaScriptParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(JavaScriptParserLPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 144105726279840) != 0) {
		{
			p.SetState(126)
			p.Exprlist()
		}

	}
	{
		p.SetState(129)
		p.Match(JavaScriptParserRPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IClassContext is an interface to support dynamic dispatch.
type IClassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLASS() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	Cons() IConsContext
	AllMethod() []IMethodContext
	Method(i int) IMethodContext

	// IsClassContext differentiates from other interfaces.
	IsClassContext()
}

type ClassContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassContext() *ClassContext {
	var p = new(ClassContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_class
	return p
}

func InitEmptyClassContext(p *ClassContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_class
}

func (*ClassContext) IsClassContext() {}

func NewClassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassContext {
	var p = new(ClassContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_class

	return p
}

func (s *ClassContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassContext) CLASS() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCLASS, 0)
}

func (s *ClassContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIDENTIFIER, 0)
}

func (s *ClassContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLBRACE, 0)
}

func (s *ClassContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRBRACE, 0)
}

func (s *ClassContext) Cons() IConsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConsContext)
}

func (s *ClassContext) AllMethod() []IMethodContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMethodContext); ok {
			len++
		}
	}

	tst := make([]IMethodContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMethodContext); ok {
			tst[i] = t.(IMethodContext)
			i++
		}
	}

	return tst
}

func (s *ClassContext) Method(i int) IMethodContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *ClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ClassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterClass(s)
	}
}

func (s *ClassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitClass(s)
	}
}




func (p *JavaScriptParser) Class() (localctx IClassContext) {
	localctx = NewClassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, JavaScriptParserRULE_class)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(JavaScriptParserCLASS)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Match(JavaScriptParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(JavaScriptParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == JavaScriptParserCONS {
		{
			p.SetState(134)
			p.Cons()
		}

	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == JavaScriptParserIDENTIFIER {
		{
			p.SetState(137)
			p.Method()
		}


		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
		p.Match(JavaScriptParserRBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IConsContext is an interface to support dynamic dispatch.
type IConsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONS() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	Paramlist() IParamlistContext

	// IsConsContext differentiates from other interfaces.
	IsConsContext()
}

type ConsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConsContext() *ConsContext {
	var p = new(ConsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_cons
	return p
}

func InitEmptyConsContext(p *ConsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_cons
}

func (*ConsContext) IsConsContext() {}

func NewConsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConsContext {
	var p = new(ConsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_cons

	return p
}

func (s *ConsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConsContext) CONS() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCONS, 0)
}

func (s *ConsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLPAREN, 0)
}

func (s *ConsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRPAREN, 0)
}

func (s *ConsContext) Block() IBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ConsContext) Paramlist() IParamlistContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamlistContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamlistContext)
}

func (s *ConsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterCons(s)
	}
}

func (s *ConsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitCons(s)
	}
}




func (p *JavaScriptParser) Cons() (localctx IConsContext) {
	localctx = NewConsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, JavaScriptParserRULE_cons)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(JavaScriptParserCONS)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Match(JavaScriptParserLPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == JavaScriptParserIDENTIFIER {
		{
			p.SetState(147)
			p.Paramlist()
		}

	}
	{
		p.SetState(150)
		p.Match(JavaScriptParserRPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Block()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	Paramlist() IParamlistContext

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_method
	return p
}

func InitEmptyMethodContext(p *MethodContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_method
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIDENTIFIER, 0)
}

func (s *MethodContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLPAREN, 0)
}

func (s *MethodContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRPAREN, 0)
}

func (s *MethodContext) Block() IBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *MethodContext) Paramlist() IParamlistContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamlistContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamlistContext)
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitMethod(s)
	}
}




func (p *JavaScriptParser) Method() (localctx IMethodContext) {
	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, JavaScriptParserRULE_method)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(JavaScriptParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Match(JavaScriptParserLPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == JavaScriptParserIDENTIFIER {
		{
			p.SetState(155)
			p.Paramlist()
		}

	}
	{
		p.SetState(158)
		p.Match(JavaScriptParserRPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Block()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllGlobal() []IGlobalContext
	Global(i int) IGlobalContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllGlobal() []IGlobalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGlobalContext); ok {
			len++
		}
	}

	tst := make([]IGlobalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGlobalContext); ok {
			tst[i] = t.(IGlobalContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Global(i int) IGlobalContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobalContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitProgram(s)
	}
}




func (p *JavaScriptParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, JavaScriptParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 144105726473406) != 0) {
		{
			p.SetState(161)
			p.Global()
		}


		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IGlobalContext is an interface to support dynamic dispatch.
type IGlobalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Funcdef() IFuncdefContext
	Stmg() IStmgContext
	Class() IClassContext

	// IsGlobalContext differentiates from other interfaces.
	IsGlobalContext()
}

type GlobalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalContext() *GlobalContext {
	var p = new(GlobalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_global
	return p
}

func InitEmptyGlobalContext(p *GlobalContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_global
}

func (*GlobalContext) IsGlobalContext() {}

func NewGlobalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalContext {
	var p = new(GlobalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_global

	return p
}

func (s *GlobalContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalContext) Funcdef() IFuncdefContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncdefContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncdefContext)
}

func (s *GlobalContext) Stmg() IStmgContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmgContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmgContext)
}

func (s *GlobalContext) Class() IClassContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassContext)
}

func (s *GlobalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *GlobalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterGlobal(s)
	}
}

func (s *GlobalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitGlobal(s)
	}
}




func (p *JavaScriptParser) Global() (localctx IGlobalContext) {
	localctx = NewGlobalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, JavaScriptParserRULE_global)
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JavaScriptParserFUNCTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Funcdef()
		}


	case JavaScriptParserRETURN, JavaScriptParserVAR, JavaScriptParserTHIS, JavaScriptParserNEW, JavaScriptParserIF, JavaScriptParserWHILE, JavaScriptParserFOR, JavaScriptParserBREAK, JavaScriptParserCONTINUE, JavaScriptParserSWITCH, JavaScriptParserLPAREN, JavaScriptParserLBRACE, JavaScriptParserLBRACKET, JavaScriptParserNUMBER, JavaScriptParserIDENTIFIER, JavaScriptParserSTRING, JavaScriptParserSUB, JavaScriptParserINC, JavaScriptParserDEC, JavaScriptParserNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(168)
			p.Stmg()
		}


	case JavaScriptParserCLASS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(169)
			p.Class()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	RPAREN() antlr.TerminalNode
	INC() antlr.TerminalNode
	DEC() antlr.TerminalNode
	NOT() antlr.TerminalNode
	SUB() antlr.TerminalNode
	This() IThisContext
	Cobj() ICobjContext
	Funcall() IFuncallContext
	Num() INumContext
	Id() IIdContext
	Str() IStrContext
	Arr() IArrContext
	Obj() IObjContext
	DOT() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	ADD() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	EQ() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GTE() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	AEQ() antlr.TerminalNode
	NAEQ() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLPAREN, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRPAREN, 0)
}

func (s *ExprContext) INC() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserINC, 0)
}

func (s *ExprContext) DEC() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDEC, 0)
}

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNOT, 0)
}

func (s *ExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSUB, 0)
}

func (s *ExprContext) This() IThisContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThisContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThisContext)
}

func (s *ExprContext) Cobj() ICobjContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICobjContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICobjContext)
}

func (s *ExprContext) Funcall() IFuncallContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncallContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncallContext)
}

func (s *ExprContext) Num() INumContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *ExprContext) Id() IIdContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ExprContext) Str() IStrContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *ExprContext) Arr() IArrContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrContext)
}

func (s *ExprContext) Obj() IObjContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *ExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDOT, 0)
}

func (s *ExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMUL, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDIV, 0)
}

func (s *ExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserADD, 0)
}

func (s *ExprContext) LT() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLT, 0)
}

func (s *ExprContext) GT() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserGT, 0)
}

func (s *ExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserEQ, 0)
}

func (s *ExprContext) LTE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLTE, 0)
}

func (s *ExprContext) GTE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserGTE, 0)
}

func (s *ExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNEQ, 0)
}

func (s *ExprContext) AEQ() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAEQ, 0)
}

func (s *ExprContext) NAEQ() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNAEQ, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOR, 0)
}

func (s *ExprContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLBRACKET, 0)
}

func (s *ExprContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRBRACKET, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitExpr(s)
	}
}





func (p *JavaScriptParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *JavaScriptParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, JavaScriptParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(173)
			p.Match(JavaScriptParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(174)
			p.expr(0)
		}
		{
			p.SetState(175)
			p.Match(JavaScriptParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 2:
		{
			p.SetState(177)
			_la = p.GetTokenStream().LA(1)

			if !(_la == JavaScriptParserINC || _la == JavaScriptParserDEC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(178)
			p.expr(11)
		}


	case 3:
		{
			p.SetState(179)
			_la = p.GetTokenStream().LA(1)

			if !(_la == JavaScriptParserSUB || _la == JavaScriptParserNOT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(180)
			p.expr(9)
		}


	case 4:
		{
			p.SetState(181)
			p.This()
		}


	case 5:
		{
			p.SetState(182)
			p.Cobj()
		}


	case 6:
		{
			p.SetState(183)
			p.Funcall()
		}


	case 7:
		{
			p.SetState(184)
			p.Num()
		}


	case 8:
		{
			p.SetState(185)
			p.Id()
		}


	case 9:
		{
			p.SetState(186)
			p.Str()
		}


	case 10:
		{
			p.SetState(187)
			p.Arr()
		}


	case 11:
		{
			p.SetState(188)
			p.Obj()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(210)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_expr)
				p.SetState(191)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(192)
					p.Match(JavaScriptParserDOT)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(193)
					p.expr(17)
				}


			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_expr)
				p.SetState(194)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(195)
					_la = p.GetTokenStream().LA(1)

					if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 515396075520) != 0)) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(196)
					p.expr(15)
				}


			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_expr)
				p.SetState(197)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(198)
					_la = p.GetTokenStream().LA(1)

					if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 71776119061217280) != 0)) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(199)
					p.expr(14)
				}


			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_expr)
				p.SetState(200)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(201)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JavaScriptParserAND || _la == JavaScriptParserOR) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(202)
					p.expr(13)
				}


			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_expr)
				p.SetState(203)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(204)
					p.Match(JavaScriptParserLBRACKET)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(205)
					p.expr(0)
				}
				{
					p.SetState(206)
					p.Match(JavaScriptParserRBRACKET)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_expr)
				p.SetState(208)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(209)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JavaScriptParserINC || _la == JavaScriptParserDEC) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}



	errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IStmgContext is an interface to support dynamic dispatch.
type IStmgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Stm() IStmContext
	SEMICOLON() antlr.TerminalNode

	// IsStmgContext differentiates from other interfaces.
	IsStmgContext()
}

type StmgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmgContext() *StmgContext {
	var p = new(StmgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_stmg
	return p
}

func InitEmptyStmgContext(p *StmgContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_stmg
}

func (*StmgContext) IsStmgContext() {}

func NewStmgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmgContext {
	var p = new(StmgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_stmg

	return p
}

func (s *StmgContext) GetParser() antlr.Parser { return s.parser }

func (s *StmgContext) Stm() IStmContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmContext)
}

func (s *StmgContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSEMICOLON, 0)
}

func (s *StmgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StmgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterStmg(s)
	}
}

func (s *StmgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitStmg(s)
	}
}




func (p *JavaScriptParser) Stmg() (localctx IStmgContext) {
	localctx = NewStmgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, JavaScriptParserRULE_stmg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Stm()
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == JavaScriptParserSEMICOLON {
		{
			p.SetState(216)
			p.Match(JavaScriptParserSEMICOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IStmContext is an interface to support dynamic dispatch.
type IStmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	ASSIGN() antlr.TerminalNode
	ADD_ASSIGN() antlr.TerminalNode
	SUB_ASSIGN() antlr.TerminalNode
	MUL_ASSIGN() antlr.TerminalNode
	DIV_ASSIGN() antlr.TerminalNode
	VAR() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	RETURN() antlr.TerminalNode
	BREAK() antlr.TerminalNode
	CONTINUE() antlr.TerminalNode
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode
	WHILE() antlr.TerminalNode
	FOR() antlr.TerminalNode
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	AllStm() []IStmContext
	Stm(i int) IStmContext
	SWITCH() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllCase_() []ICaseContext
	Case_(i int) ICaseContext
	Default_() IDefaultContext

	// IsStmContext differentiates from other interfaces.
	IsStmContext()
}

type StmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmContext() *StmContext {
	var p = new(StmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_stm
	return p
}

func InitEmptyStmContext(p *StmContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_stm
}

func (*StmContext) IsStmContext() {}

func NewStmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmContext {
	var p = new(StmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_stm

	return p
}

func (s *StmContext) GetParser() antlr.Parser { return s.parser }

func (s *StmContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *StmContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserASSIGN, 0)
}

func (s *StmContext) ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserADD_ASSIGN, 0)
}

func (s *StmContext) SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSUB_ASSIGN, 0)
}

func (s *StmContext) MUL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMUL_ASSIGN, 0)
}

func (s *StmContext) DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDIV_ASSIGN, 0)
}

func (s *StmContext) VAR() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserVAR, 0)
}

func (s *StmContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIDENTIFIER, 0)
}

func (s *StmContext) RETURN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRETURN, 0)
}

func (s *StmContext) BREAK() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBREAK, 0)
}

func (s *StmContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCONTINUE, 0)
}

func (s *StmContext) IF() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIF, 0)
}

func (s *StmContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLPAREN, 0)
}

func (s *StmContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRPAREN, 0)
}

func (s *StmContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *StmContext) Block(i int) IBlockContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StmContext) ELSE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserELSE, 0)
}

func (s *StmContext) WHILE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserWHILE, 0)
}

func (s *StmContext) FOR() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserFOR, 0)
}

func (s *StmContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserSEMICOLON)
}

func (s *StmContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSEMICOLON, i)
}

func (s *StmContext) AllStm() []IStmContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmContext); ok {
			len++
		}
	}

	tst := make([]IStmContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmContext); ok {
			tst[i] = t.(IStmContext)
			i++
		}
	}

	return tst
}

func (s *StmContext) Stm(i int) IStmContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmContext)
}

func (s *StmContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSWITCH, 0)
}

func (s *StmContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLBRACE, 0)
}

func (s *StmContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRBRACE, 0)
}

func (s *StmContext) AllCase_() []ICaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseContext); ok {
			len++
		}
	}

	tst := make([]ICaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseContext); ok {
			tst[i] = t.(ICaseContext)
			i++
		}
	}

	return tst
}

func (s *StmContext) Case_(i int) ICaseContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseContext)
}

func (s *StmContext) Default_() IDefaultContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultContext)
}

func (s *StmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterStm(s)
	}
}

func (s *StmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitStm(s)
	}
}




func (p *JavaScriptParser) Stm() (localctx IStmContext) {
	localctx = NewStmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, JavaScriptParserRULE_stm)
	var _la int

	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)
			p.expr(0)
		}
		{
			p.SetState(220)
			_la = p.GetTokenStream().LA(1)

			if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 65987877535744) != 0)) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(221)
			p.expr(0)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.Match(JavaScriptParserVAR)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(JavaScriptParserIDENTIFIER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == JavaScriptParserASSIGN {
			{
				p.SetState(225)
				p.Match(JavaScriptParserASSIGN)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(226)
				p.expr(0)
			}

		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(229)
			p.Match(JavaScriptParserRETURN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(230)
			p.expr(0)
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(231)
			p.Match(JavaScriptParserBREAK)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(232)
			p.Match(JavaScriptParserCONTINUE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(233)
			p.Match(JavaScriptParserIF)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Match(JavaScriptParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(235)
			p.expr(0)
		}
		{
			p.SetState(236)
			p.Match(JavaScriptParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Block()
		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == JavaScriptParserELSE {
			{
				p.SetState(238)
				p.Match(JavaScriptParserELSE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(239)
				p.Block()
			}

		}


	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(242)
			p.Match(JavaScriptParserWHILE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Match(JavaScriptParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(244)
			p.expr(0)
		}
		{
			p.SetState(245)
			p.Match(JavaScriptParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Block()
		}


	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(248)
			p.Match(JavaScriptParserFOR)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(249)
			p.Match(JavaScriptParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 144105726473388) != 0) {
			{
				p.SetState(250)
				p.Stm()
			}

		}
		{
			p.SetState(253)
			p.Match(JavaScriptParserSEMICOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 144105726473388) != 0) {
			{
				p.SetState(254)
				p.Stm()
			}

		}
		{
			p.SetState(257)
			p.Match(JavaScriptParserSEMICOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 144105726473388) != 0) {
			{
				p.SetState(258)
				p.Stm()
			}

		}
		{
			p.SetState(261)
			p.Match(JavaScriptParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(262)
			p.Block()
		}


	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(263)
			p.Match(JavaScriptParserSWITCH)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(264)
			p.Match(JavaScriptParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(265)
			p.expr(0)
		}
		{
			p.SetState(266)
			p.Match(JavaScriptParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Match(JavaScriptParserLBRACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == JavaScriptParserCASE {
			{
				p.SetState(268)
				p.Case_()
			}


			p.SetState(273)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == JavaScriptParserDEFAULT {
			{
				p.SetState(274)
				p.Default_()
			}

		}
		{
			p.SetState(277)
			p.Match(JavaScriptParserRBRACE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(279)
			p.Block()
		}


	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(280)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ICaseContext is an interface to support dynamic dispatch.
type ICaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expr() IExprContext
	COL() antlr.TerminalNode
	AllStm() []IStmContext
	Stm(i int) IStmContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsCaseContext differentiates from other interfaces.
	IsCaseContext()
}

type CaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseContext() *CaseContext {
	var p = new(CaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_case
	return p
}

func InitEmptyCaseContext(p *CaseContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_case
}

func (*CaseContext) IsCaseContext() {}

func NewCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseContext {
	var p = new(CaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_case

	return p
}

func (s *CaseContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCASE, 0)
}

func (s *CaseContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CaseContext) COL() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCOL, 0)
}

func (s *CaseContext) AllStm() []IStmContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmContext); ok {
			len++
		}
	}

	tst := make([]IStmContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmContext); ok {
			tst[i] = t.(IStmContext)
			i++
		}
	}

	return tst
}

func (s *CaseContext) Stm(i int) IStmContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmContext)
}

func (s *CaseContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserSEMICOLON)
}

func (s *CaseContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSEMICOLON, i)
}

func (s *CaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterCase(s)
	}
}

func (s *CaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitCase(s)
	}
}




func (p *JavaScriptParser) Case_() (localctx ICaseContext) {
	localctx = NewCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, JavaScriptParserRULE_case)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Match(JavaScriptParserCASE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(284)
		p.expr(0)
	}
	{
		p.SetState(285)
		p.Match(JavaScriptParserCOL)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 144105726473388) != 0) {
		{
			p.SetState(286)
			p.Stm()
		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == JavaScriptParserSEMICOLON {
			{
				p.SetState(287)
				p.Match(JavaScriptParserSEMICOLON)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}


		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IDefaultContext is an interface to support dynamic dispatch.
type IDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	COL() antlr.TerminalNode
	AllStm() []IStmContext
	Stm(i int) IStmContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsDefaultContext differentiates from other interfaces.
	IsDefaultContext()
}

type DefaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultContext() *DefaultContext {
	var p = new(DefaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_default
	return p
}

func InitEmptyDefaultContext(p *DefaultContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_default
}

func (*DefaultContext) IsDefaultContext() {}

func NewDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultContext {
	var p = new(DefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_default

	return p
}

func (s *DefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDEFAULT, 0)
}

func (s *DefaultContext) COL() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCOL, 0)
}

func (s *DefaultContext) AllStm() []IStmContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmContext); ok {
			len++
		}
	}

	tst := make([]IStmContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmContext); ok {
			tst[i] = t.(IStmContext)
			i++
		}
	}

	return tst
}

func (s *DefaultContext) Stm(i int) IStmContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmContext)
}

func (s *DefaultContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserSEMICOLON)
}

func (s *DefaultContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSEMICOLON, i)
}

func (s *DefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterDefault(s)
	}
}

func (s *DefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitDefault(s)
	}
}




func (p *JavaScriptParser) Default_() (localctx IDefaultContext) {
	localctx = NewDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, JavaScriptParserRULE_default)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Match(JavaScriptParserDEFAULT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(296)
		p.Match(JavaScriptParserCOL)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 144105726473388) != 0) {
		{
			p.SetState(297)
			p.Stm()
		}
		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == JavaScriptParserSEMICOLON {
			{
				p.SetState(298)
				p.Match(JavaScriptParserSEMICOLON)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}


		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStm() []IStmContext
	Stm(i int) IStmContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRBRACE, 0)
}

func (s *BlockContext) AllStm() []IStmContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmContext); ok {
			len++
		}
	}

	tst := make([]IStmContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmContext); ok {
			tst[i] = t.(IStmContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stm(i int) IStmContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmContext)
}

func (s *BlockContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserSEMICOLON)
}

func (s *BlockContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSEMICOLON, i)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitBlock(s)
	}
}




func (p *JavaScriptParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, JavaScriptParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(JavaScriptParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 144105726473388) != 0) {
		{
			p.SetState(307)
			p.Stm()
		}
		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == JavaScriptParserSEMICOLON {
			{
				p.SetState(308)
				p.Match(JavaScriptParserSEMICOLON)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}


		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(316)
		p.Match(JavaScriptParserRBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


func (p *JavaScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 19:
			var t *ExprContext = nil
			if localctx != nil { t = localctx.(*ExprContext) }
			return p.Expr_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JavaScriptParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 16)

	case 1:
			return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
			return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
			return p.Precpred(p.GetParserRuleContext(), 15)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

