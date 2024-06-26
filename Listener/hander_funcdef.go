package listener

import (
	"fmt"

	Jsp "github.com/klang/js-go/JavaScript"
)

func (h *Handler) handFuncdef(name string, args []Variant) (BaseObject, int) {
	currentT := current
	defer func() {
		vl.delVarAll()
		current = currentT
	}()
	current = name
	params, ctx := vl.getFunc(name)

	if ctx == nil {
		value, typ := vl.callInFunc(name, args)
		return value, typ
	}

	if len(params) != len(args) {
		panic("The parameters do not match")
	} else {
		for i, param := range params {
			value, typ := args[i].value, args[i].typ
			vl.addVar(param, value, typ)
		}
	}
	block := ctx
	revalue, retyp := h.handBlock(block)

	return revalue, retyp
}

func (h *Handler) handBlock(blo *Jsp.BlockContext) (BaseObject, int) {
	defer vl.leaveScope(vl.scope)
	vl.enterScope()

	for _, c := range blo.GetChildren() {
		if h.isBreak {
			break
		}
		if h.isContinue {
			continue
		}
		if c.GetChild(0) != nil {
			value, typ := h.handStm(c.(*Jsp.StmContext))
			if h.isReturn {
				return value, typ
			}
		}
	}

	return nil, -1
}

func (h *Handler) handExpr_funcall(call *Jsp.FuncallContext) (BaseObject, int) {
	var args []Variant = []Variant{}

	name := fmt.Sprintf("%v", call.GetChild(0))

	for _, c := range call.GetChild(2).GetChildren() {
		if c.GetChild(0) != nil {
			value, typ := h.handExpr(c.(*Jsp.ExprContext))
			args = append(args, Variant{"unknown", value, typ})
		}
	}

	return h.handFuncdef(name, args)
}
