// Code generated from JavaScript/JavaScriptParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // JavaScriptParser

import "github.com/antlr4-go/antlr/v4"

// BaseJavaScriptParserListener is a complete listener for a parse tree produced by JavaScriptParser.
type BaseJavaScriptParserListener struct{}

var _ JavaScriptParserListener = &BaseJavaScriptParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJavaScriptParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJavaScriptParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJavaScriptParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJavaScriptParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNum is called when production num is entered.
func (s *BaseJavaScriptParserListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production num is exited.
func (s *BaseJavaScriptParserListener) ExitNum(ctx *NumContext) {}

// EnterId is called when production id is entered.
func (s *BaseJavaScriptParserListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseJavaScriptParserListener) ExitId(ctx *IdContext) {}

// EnterStr is called when production str is entered.
func (s *BaseJavaScriptParserListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BaseJavaScriptParserListener) ExitStr(ctx *StrContext) {}

// EnterArr is called when production arr is entered.
func (s *BaseJavaScriptParserListener) EnterArr(ctx *ArrContext) {}

// ExitArr is called when production arr is exited.
func (s *BaseJavaScriptParserListener) ExitArr(ctx *ArrContext) {}

// EnterKey is called when production key is entered.
func (s *BaseJavaScriptParserListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseJavaScriptParserListener) ExitKey(ctx *KeyContext) {}

// EnterObj is called when production obj is entered.
func (s *BaseJavaScriptParserListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseJavaScriptParserListener) ExitObj(ctx *ObjContext) {}

// EnterBool is called when production bool is entered.
func (s *BaseJavaScriptParserListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production bool is exited.
func (s *BaseJavaScriptParserListener) ExitBool(ctx *BoolContext) {}

// EnterThis is called when production this is entered.
func (s *BaseJavaScriptParserListener) EnterThis(ctx *ThisContext) {}

// ExitThis is called when production this is exited.
func (s *BaseJavaScriptParserListener) ExitThis(ctx *ThisContext) {}

// EnterFuncdef is called when production funcdef is entered.
func (s *BaseJavaScriptParserListener) EnterFuncdef(ctx *FuncdefContext) {}

// ExitFuncdef is called when production funcdef is exited.
func (s *BaseJavaScriptParserListener) ExitFuncdef(ctx *FuncdefContext) {}

// EnterParamlist is called when production paramlist is entered.
func (s *BaseJavaScriptParserListener) EnterParamlist(ctx *ParamlistContext) {}

// ExitParamlist is called when production paramlist is exited.
func (s *BaseJavaScriptParserListener) ExitParamlist(ctx *ParamlistContext) {}

// EnterParam is called when production param is entered.
func (s *BaseJavaScriptParserListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseJavaScriptParserListener) ExitParam(ctx *ParamContext) {}

// EnterFuncall is called when production funcall is entered.
func (s *BaseJavaScriptParserListener) EnterFuncall(ctx *FuncallContext) {}

// ExitFuncall is called when production funcall is exited.
func (s *BaseJavaScriptParserListener) ExitFuncall(ctx *FuncallContext) {}

// EnterExprlist is called when production exprlist is entered.
func (s *BaseJavaScriptParserListener) EnterExprlist(ctx *ExprlistContext) {}

// ExitExprlist is called when production exprlist is exited.
func (s *BaseJavaScriptParserListener) ExitExprlist(ctx *ExprlistContext) {}

// EnterCobj is called when production cobj is entered.
func (s *BaseJavaScriptParserListener) EnterCobj(ctx *CobjContext) {}

// ExitCobj is called when production cobj is exited.
func (s *BaseJavaScriptParserListener) ExitCobj(ctx *CobjContext) {}

// EnterClass is called when production class is entered.
func (s *BaseJavaScriptParserListener) EnterClass(ctx *ClassContext) {}

// ExitClass is called when production class is exited.
func (s *BaseJavaScriptParserListener) ExitClass(ctx *ClassContext) {}

// EnterCons is called when production cons is entered.
func (s *BaseJavaScriptParserListener) EnterCons(ctx *ConsContext) {}

// ExitCons is called when production cons is exited.
func (s *BaseJavaScriptParserListener) ExitCons(ctx *ConsContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseJavaScriptParserListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseJavaScriptParserListener) ExitMethod(ctx *MethodContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseJavaScriptParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseJavaScriptParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterGlobal is called when production global is entered.
func (s *BaseJavaScriptParserListener) EnterGlobal(ctx *GlobalContext) {}

// ExitGlobal is called when production global is exited.
func (s *BaseJavaScriptParserListener) ExitGlobal(ctx *GlobalContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseJavaScriptParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseJavaScriptParserListener) ExitExpr(ctx *ExprContext) {}

// EnterStmg is called when production stmg is entered.
func (s *BaseJavaScriptParserListener) EnterStmg(ctx *StmgContext) {}

// ExitStmg is called when production stmg is exited.
func (s *BaseJavaScriptParserListener) ExitStmg(ctx *StmgContext) {}

// EnterStm is called when production stm is entered.
func (s *BaseJavaScriptParserListener) EnterStm(ctx *StmContext) {}

// ExitStm is called when production stm is exited.
func (s *BaseJavaScriptParserListener) ExitStm(ctx *StmContext) {}

// EnterCase is called when production case is entered.
func (s *BaseJavaScriptParserListener) EnterCase(ctx *CaseContext) {}

// ExitCase is called when production case is exited.
func (s *BaseJavaScriptParserListener) ExitCase(ctx *CaseContext) {}

// EnterDefault is called when production default is entered.
func (s *BaseJavaScriptParserListener) EnterDefault(ctx *DefaultContext) {}

// ExitDefault is called when production default is exited.
func (s *BaseJavaScriptParserListener) ExitDefault(ctx *DefaultContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseJavaScriptParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseJavaScriptParserListener) ExitBlock(ctx *BlockContext) {}
