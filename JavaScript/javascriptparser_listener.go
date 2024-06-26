// Code generated from JavaScript/JavaScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // JavaScriptParser

import "github.com/antlr4-go/antlr/v4"


// JavaScriptParserListener is a complete listener for a parse tree produced by JavaScriptParser.
type JavaScriptParserListener interface {
	antlr.ParseTreeListener

	// EnterNum is called when entering the num production.
	EnterNum(c *NumContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterStr is called when entering the str production.
	EnterStr(c *StrContext)

	// EnterArr is called when entering the arr production.
	EnterArr(c *ArrContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterBool is called when entering the bool production.
	EnterBool(c *BoolContext)

	// EnterThis is called when entering the this production.
	EnterThis(c *ThisContext)

	// EnterFuncdef is called when entering the funcdef production.
	EnterFuncdef(c *FuncdefContext)

	// EnterParamlist is called when entering the paramlist production.
	EnterParamlist(c *ParamlistContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterFuncall is called when entering the funcall production.
	EnterFuncall(c *FuncallContext)

	// EnterExprlist is called when entering the exprlist production.
	EnterExprlist(c *ExprlistContext)

	// EnterCobj is called when entering the cobj production.
	EnterCobj(c *CobjContext)

	// EnterClass is called when entering the class production.
	EnterClass(c *ClassContext)

	// EnterCons is called when entering the cons production.
	EnterCons(c *ConsContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterGlobal is called when entering the global production.
	EnterGlobal(c *GlobalContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterStmg is called when entering the stmg production.
	EnterStmg(c *StmgContext)

	// EnterStm is called when entering the stm production.
	EnterStm(c *StmContext)

	// EnterCase is called when entering the case production.
	EnterCase(c *CaseContext)

	// EnterDefault is called when entering the default production.
	EnterDefault(c *DefaultContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// ExitNum is called when exiting the num production.
	ExitNum(c *NumContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitStr is called when exiting the str production.
	ExitStr(c *StrContext)

	// ExitArr is called when exiting the arr production.
	ExitArr(c *ArrContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitBool is called when exiting the bool production.
	ExitBool(c *BoolContext)

	// ExitThis is called when exiting the this production.
	ExitThis(c *ThisContext)

	// ExitFuncdef is called when exiting the funcdef production.
	ExitFuncdef(c *FuncdefContext)

	// ExitParamlist is called when exiting the paramlist production.
	ExitParamlist(c *ParamlistContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitFuncall is called when exiting the funcall production.
	ExitFuncall(c *FuncallContext)

	// ExitExprlist is called when exiting the exprlist production.
	ExitExprlist(c *ExprlistContext)

	// ExitCobj is called when exiting the cobj production.
	ExitCobj(c *CobjContext)

	// ExitClass is called when exiting the class production.
	ExitClass(c *ClassContext)

	// ExitCons is called when exiting the cons production.
	ExitCons(c *ConsContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitGlobal is called when exiting the global production.
	ExitGlobal(c *GlobalContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitStmg is called when exiting the stmg production.
	ExitStmg(c *StmgContext)

	// ExitStm is called when exiting the stm production.
	ExitStm(c *StmContext)

	// ExitCase is called when exiting the case production.
	ExitCase(c *CaseContext)

	// ExitDefault is called when exiting the default production.
	ExitDefault(c *DefaultContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)
}
