package listener

import (
	"fmt"

	Jsp "github.com/klang/js-go/JavaScript"
)

func (h *Handler) handExpr_cobj(new *Jsp.CobjContext) (BaseObject, int) {
	var args []Variant
	cname := fmt.Sprintf("%s", new.GetChild(1))
	cobjt := &CobjObject{Value: make(map[string]BaseObject), ClassName: cname}

	if class, ok := vl.getClass(cname); ok {
		for _, o := range new.GetChild(3).GetChildren() {
			if o.GetChild(0) != nil {
				value, typ := h.handExpr(o.(*Jsp.ExprContext))
				args = append(args, Variant{"unknown", value, typ})
			}
		}
		h.handMethod(cobjt, class, "constructor", args)
	}

	return cobjt, Jsp.JavaScriptParserRULE_cobj
}

func (h *Handler) handMethod(cobjt *CobjObject, c *Class, name string, args []Variant) (BaseObject, int) {
	currentT := current
	defer func() {
		vl.delVarAll()
		current = currentT
	}()
	current = name
	fun, _ := c.getFunc(name)

	if len(fun.args) != len(args) {
		panic("The parameters do not match")
	} else {
		for i, param := range fun.args {
			value, typ := args[i].value, args[i].typ
			vl.addVar(param, value, typ)
		}
	}

	block := fun.fctx
	for _, c := range block.GetChildren() {
		if h.isBreak {
			break
		}
		if h.isContinue {
			continue
		}
		if c.GetChild(0) != nil {
			value, typ := h.handStm(c.(*Jsp.StmContext))
			if this, ok := value.(*ThisObjetc); ok {
				cobjt.addValue(this.Value, this.Object)
			}
			if h.isReturn {
				return value, typ
			}
		}
	}

	return nil, -1
}
