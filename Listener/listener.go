package listener

import (
	"fmt"
	"regexp"

	Jsp "github.com/klang/js-go/JavaScript"
)

var h *Handler
var renum *regexp.Regexp = regexp.MustCompile(`^\d+$`)
var resym *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

type Handler struct {
	isBreak    bool
	isContinue bool
	isReturn   bool
	nextId     int
}

func NewHandler() *Handler {
	return &Handler{
		isBreak:    false,
		isContinue: false,
		isReturn:   false,
	}
}

func (p *Handler) genId() string {
	id := fmt.Sprintf("%%t%d", p.nextId)
	p.nextId++
	return id
}

func (p *Handler) genLabelId(name string) string {
	id := fmt.Sprintf("%s.%d", name, p.nextId)
	p.nextId++
	return id
}

var current string

type MyListener struct {
	*Jsp.BaseJavaScriptParserListener
}

func (l *MyListener) EnterProgram(ctx *Jsp.ProgramContext) {
	vl = NewSymTable()
	h = NewHandler()
	vl.enterScope()
}

func (l *MyListener) EnterFuncdef(ctx *Jsp.FuncdefContext) {
	var args []string
	name := fmt.Sprintf("%v", ctx.GetChild(1))
	for _, v := range ctx.GetChild(3).GetChildren() {
		if v.GetChild(0) != nil {
			args = append(args, fmt.Sprintf("%v", v.GetChild(0)))
		}
	}
	vl.addFunc(name, args, ctx.GetChild(5).(*Jsp.BlockContext))
	vl.showFuncAll()
}

func (l *MyListener) EnterClass(ctx *Jsp.ClassContext) {
	name := fmt.Sprintf("%v", ctx.GetChild(1))
	class := NewClass(name)

	for _, c := range ctx.GetChildren() {
		if con, ok := c.(*Jsp.ConsContext); ok {
			class.initFunc(con)
		}
		if met, ok := c.(*Jsp.MethodContext); ok {
			class.addFunc(met)
		}
	}
	vl.addClass(class)
}

func (l *MyListener) EnterStmg(ctx *Jsp.StmgContext) {
	current = "globol"
	h.handStm(ctx.GetChild(0).(*Jsp.StmContext))
}
