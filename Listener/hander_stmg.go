package listener

import (
	"fmt"
	"strconv"

	Jsp "github.com/klang/js-go/JavaScript"
)

func (h *Handler) handStm(stm *Jsp.StmContext) (BaseObject, int) {
	var str_1 string
	var value BaseObject
	var typ int

	if stm == nil {
		return nil, 0
	}

	str_1 = fmt.Sprintf("%v", stm.GetChild(0))
	if str_1 == "return" {
		h.isReturn = true
		value, typ = h.handExpr(stm.GetChild(1).(*Jsp.ExprContext))
		return value, typ
	}
	if str_1 == "var" {
		name := fmt.Sprintf("%v", stm.GetChild(1))
		if stm.GetChildCount() == 4 {
			value, typ = h.handExpr(stm.GetChild(3).(*Jsp.ExprContext))
		} else {
			value = &UndefObject{}
		}
		vl.addVar(name, value, typ)
		vl.showVarAll()
	}
	if str_1 == "if" {
		vaule, typ := h.handExpr(stm.GetChild(2).(*Jsp.ExprContext))
		if ok := h.checkExpr(vaule, typ); ok {
			h.handBlock(stm.GetChild(4).(*Jsp.BlockContext))
		} else if stm.GetChildCount() == 7 {
			h.handBlock(stm.GetChild(6).(*Jsp.BlockContext))
		}
	}
	if str_1 == "while" {
		for {
			if h.isBreak {
				h.isBreak = false
				break
			}
			if h.isContinue {
				h.isContinue = false
				continue
			}
			vaule, typ := h.handExpr(stm.GetChild(2).(*Jsp.ExprContext))
			if ok := h.checkExpr(vaule, typ); !ok {
				break
			}
			h.handBlock(stm.GetChild(4).(*Jsp.BlockContext))
		}
	}
	if str_1 == "for" {
		var smts [3]*Jsp.StmContext
		var index int = 0
		for _, c := range stm.GetChildren() {
			_, ok := c.(*Jsp.StmContext)
			if tmp := fmt.Sprintf("%v", c); tmp == ";" {
				index++
			}
			if c.GetChild(0) != nil && ok {
				smts[index] = c.(*Jsp.StmContext)
			}
		}
		if smts[0] != nil {
			h.handStm(smts[0])
		}
		for {
			if h.isBreak {
				h.isBreak = false
				break
			}
			if h.isContinue {
				h.isContinue = false
				continue
			}
			if smts[1] != nil {
				vaule, typ := h.handStm(smts[1])
				if ok := h.checkExpr(vaule, typ); !ok {
					break
				}
			}
			h.handBlock(stm.GetChild(stm.GetChildCount() - 1).(*Jsp.BlockContext))
			if smts[2] != nil {
				h.handStm(smts[2])
			}
		}
	}
	if str_1 == "switch" {
		var key bool = false
		vaule, _ := h.handExpr(stm.GetChild(2).(*Jsp.ExprContext))
		for i := 0; i < stm.GetChildCount()-7; i++ {
			num, _ := h.handExpr(stm.GetChild(i + 5).GetChild(1).(*Jsp.ExprContext))
			if num.getString() == vaule.getString() || key {
				key = true
				if h.isBreak {
					h.isBreak = false
					break
				}
				for y := 0; y < stm.GetChild(i+5).GetChildCount()-2; y++ {
					if stm.GetChild(i+5).GetChild(y+2).GetChild(0) != nil {
						h.handStm(stm.GetChild(i + 5).GetChild(y + 2).(*Jsp.StmContext))
					}
				}
			}
		}
		if !key {
			for y := 0; y < stm.GetChild(stm.GetChildCount()-2).GetChildCount(); y++ {
				if stm.GetChild(stm.GetChildCount()-2).GetChild(y).GetChild(0) != nil {
					h.handStm(stm.GetChild(stm.GetChildCount() - 2).GetChild(y).(*Jsp.StmContext))
				}
			}
		}
	}
	if str_1 == "break" {
		h.isBreak = true
		return nil, 0
	}
	if str_1 == "continue" {
		h.isContinue = true
		return nil, 0
	}

	if block, ok := stm.GetChild(0).(*Jsp.BlockContext); ok {
		h.handBlock(block)
	}

	if exp1, ok := stm.GetChild(0).(*Jsp.ExprContext); ok {
		value, typ = h.handExpr(exp1)
		if stm.GetChildCount() == 3 {
			if exp2, ok := stm.GetChild(2).(*Jsp.ExprContext); ok {
				var name string
				value2, typ2 := h.handExpr(exp2)
				if exp1.GetChild(0).GetChild(0).GetChild(0) == nil {
					name = fmt.Sprintf("%v", exp1.GetChild(0).GetChild(0))
					vl.setVar(name, value2, typ2)
				} else {
					name = fmt.Sprintf("%v", exp1.GetChild(0).GetChild(0).GetChild(0))
					name += "@"
					name += fmt.Sprintf("%v", exp1.GetChild(2).GetChild(0).GetChild(0))
					vl.setVar(name, value2, Jsp.JavaScriptParserRULE_obj)
				}
			}
		}
	}

	return value, typ
}

func (h *Handler) handExpr(exp *Jsp.ExprContext) (value BaseObject, typ int) {
	if exp.GetChildCount() == 4 { // expr '[' expr ']'
		index, _ := h.handExpr(exp.GetChild(2).(*Jsp.ExprContext))
		exp, typ := h.handExpr(exp.GetChild(0).(*Jsp.ExprContext))

		if arr, ok := exp.(*ArrayObject); ok {
			i, _ := strconv.Atoi(index.getString())
			return arr.getValue(i), typ
		} else {
			panic("Array syntax error")
		}
	}

	if exp.GetChildCount() == 3 {
		if exp1, ok := exp.GetChild(1).(*Jsp.ExprContext); ok { // '(' expr ')'
			return h.handExpr(exp1)
		}

		op := fmt.Sprintf("%v", exp.GetChild(1))
		exp1, typ1 := h.handExpr(exp.GetChild(0).(*Jsp.ExprContext))
		exp2, typ2 := h.handExpr(exp.GetChild(2).(*Jsp.ExprContext))

		if op == "." { // expr '.' expr
			name1 := fmt.Sprintf("%v", exp.GetChild(0).GetChild(0).GetChild(0))
			name2 := fmt.Sprintf("%v", exp.GetChild(2).GetChild(0).GetChild(0))
			if obj, ok := exp1.(*ObjObject); ok { // object
				obj.setValue(name1, obj.Value[name2])
				str := fmt.Sprintf("%v", obj.Value[name2].getString())
				if str[0] == '"' {
					return obj.Value[name2], Jsp.JavaScriptLexerSTRING
				}
				if renum.MatchString(str) {
					return obj.Value[name2], Jsp.JavaScriptLexerNUMBER
				}
				return obj.Value[name2], -1
			} else if obj, ok := exp1.(*CobjObject); ok { // class object
				if obj.Value[name2] != nil {
					obj.setValue(name1, obj.Value[name2])
					str := fmt.Sprintf("%v", obj.Value[name2].getString())
					if str[0] == '"' {
						return obj.Value[name2], Jsp.JavaScriptLexerSTRING
					}
					if renum.MatchString(str) {
						return obj.Value[name2], Jsp.JavaScriptLexerNUMBER
					}
					return obj.Value[name2], -1
				} else {
					var args []Variant
					cname := obj.ClassName

					if call, ok := exp.GetChild(2).GetChild(0).(*Jsp.FuncallContext); ok {
						for _, c := range call.GetChild(2).GetChildren() {
							if c.GetChild(0) != nil {
								value, typ := h.handExpr(c.(*Jsp.ExprContext))
								args = append(args, *NewVariant("unknown", value, typ))
							}
						}
					}
					if class, ok := vl.getClass(cname); ok {
						value, typ := h.handMethod(obj, class, name2, args)
						return value, typ
					}
					return nil, -1
				}
			} else if this, ok := exp1.(*ThisObjetc); ok { // this
				this.setValue(name2)
				this.setObject(exp2)
				return this, Jsp.JavaScriptParserRULE_this
			} else {
				panic("Object syntax error")
			}
		}

		if op == "+" {
			if typ1 == Jsp.JavaScriptLexerNUMBER && typ2 == Jsp.JavaScriptLexerNUMBER {
				num1, _ := strconv.Atoi(exp1.getString())
				num2, _ := strconv.Atoi(exp2.getString())
				return &NumberObject{Value: num1 + num2}, Jsp.JavaScriptLexerNUMBER
			} else if typ1 == Jsp.JavaScriptLexerSTRING || typ2 == Jsp.JavaScriptLexerSTRING {
				str1 := exp1.getString()
				str2 := exp2.getString()
				if str1[0] == '"' {
					str1 = str1[1 : len(str1)-1]
				}
				if str2[0] == '"' {
					str2 = str2[1 : len(str2)-1]
				}

				return &StringObject{Value: "\"" + str1 + str2 + "\""}, Jsp.JavaScriptLexerSTRING
			} else {
				return nil, -1
			}
		}
		if op == "-" {
			if typ1 == Jsp.JavaScriptLexerNUMBER && typ2 == Jsp.JavaScriptLexerNUMBER {
				num1, _ := strconv.Atoi(exp1.getString())
				num2, _ := strconv.Atoi(exp2.getString())
				return &NumberObject{Value: num1 - num2}, Jsp.JavaScriptLexerNUMBER
			} else {
				return nil, -1
			}
		}
		if op == "*" {
			if typ1 == Jsp.JavaScriptLexerNUMBER && typ2 == Jsp.JavaScriptLexerNUMBER {
				num1, _ := strconv.Atoi(exp1.getString())
				num2, _ := strconv.Atoi(exp2.getString())
				return &NumberObject{Value: num1 * num2}, Jsp.JavaScriptLexerNUMBER
			} else {
				return nil, -1
			}
		}
		if op == "/" {
			if typ1 == Jsp.JavaScriptLexerNUMBER && typ2 == Jsp.JavaScriptLexerNUMBER {
				num1, _ := strconv.Atoi(exp1.getString())
				num2, _ := strconv.Atoi(exp2.getString())
				return &NumberObject{Value: num1 / num2}, Jsp.JavaScriptLexerNUMBER
			} else {
				return nil, -1
			}
		}
		if op == "%" {
			if typ1 == Jsp.JavaScriptLexerNUMBER && typ2 == Jsp.JavaScriptLexerNUMBER {
				num1, _ := strconv.Atoi(exp1.getString())
				num2, _ := strconv.Atoi(exp2.getString())
				return &NumberObject{Value: num1 % num2}, Jsp.JavaScriptLexerNUMBER
			} else {
				return nil, -1
			}
		}
		if op == ">" {
			if typ1 == Jsp.JavaScriptLexerNUMBER && typ2 == Jsp.JavaScriptLexerNUMBER {
				num1, _ := strconv.Atoi(exp1.getString())
				num2, _ := strconv.Atoi(exp2.getString())
				return &BooleanObject{Value: num1 > num2}, Jsp.JavaScriptLexerNUMBER
			} else {
				return nil, -1
			}
		}
		if op == ">=" {
			if typ1 == Jsp.JavaScriptLexerNUMBER && typ2 == Jsp.JavaScriptLexerNUMBER {
				num1, _ := strconv.Atoi(exp1.getString())
				num2, _ := strconv.Atoi(exp2.getString())
				return &BooleanObject{Value: num1 >= num2}, Jsp.JavaScriptLexerNUMBER
			} else {
				return nil, -1
			}
		}
		if op == "<" {
			if typ1 == Jsp.JavaScriptLexerNUMBER && typ2 == Jsp.JavaScriptLexerNUMBER {
				num1, _ := strconv.Atoi(exp1.getString())
				num2, _ := strconv.Atoi(exp2.getString())
				return &BooleanObject{Value: num1 < num2}, Jsp.JavaScriptLexerNUMBER
			} else {
				return nil, -1
			}
		}
		if op == "<=" {
			if typ1 == Jsp.JavaScriptLexerNUMBER && typ2 == Jsp.JavaScriptLexerNUMBER {
				num1, _ := strconv.Atoi(exp1.getString())
				num2, _ := strconv.Atoi(exp2.getString())
				return &BooleanObject{Value: num1 <= num2}, Jsp.JavaScriptLexerNUMBER
			} else {
				return nil, -1
			}
		}
		if op == "==" {
			if typ1 == Jsp.JavaScriptLexerNUMBER && typ2 == Jsp.JavaScriptLexerNUMBER {
				num1, _ := strconv.Atoi(exp1.getString())
				num2, _ := strconv.Atoi(exp2.getString())
				return &BooleanObject{Value: num1 == num2}, Jsp.JavaScriptLexerNUMBER
			} else {
				return nil, -1
			}
		}
		if op == "!=" {
			if typ1 == Jsp.JavaScriptLexerNUMBER && typ2 == Jsp.JavaScriptLexerNUMBER {
				num1, _ := strconv.Atoi(exp1.getString())
				num2, _ := strconv.Atoi(exp2.getString())
				return &BooleanObject{Value: num1 != num2}, Jsp.JavaScriptLexerNUMBER
			} else {
				return nil, -1
			}
		}
		if op == "===" {
			if exp1.getString() == exp2.getString() && typ1 == typ2 {
				return &BooleanObject{Value: true}, Jsp.JavaScriptLexerNUMBER
			} else {
				return &BooleanObject{Value: false}, Jsp.JavaScriptLexerNUMBER
			}
		}
		if op == "!==" {
			if exp1.getString() != exp2.getString() || typ1 != typ2 {
				return &BooleanObject{Value: true}, Jsp.JavaScriptLexerNUMBER
			} else {
				return &BooleanObject{Value: false}, Jsp.JavaScriptLexerNUMBER
			}
		}
	}

	if exp.GetChildCount() == 2 {
		var exp1 BaseObject
		var typ1 int
		var op string

		if exp.GetChild(0).GetChildCount() == 0 {
			op = fmt.Sprintf("%v", exp.GetChild(0))
			exp1, typ1 = h.handExpr(exp.GetChild(1).(*Jsp.ExprContext))
		} else {
			op = fmt.Sprintf("%v", exp.GetChild(1))
			exp1, typ1 = h.handExpr(exp.GetChild(0).(*Jsp.ExprContext))
		}

		if op == "-" {
			if typ1 == Jsp.JavaScriptLexerNUMBER {
				num1, _ := strconv.Atoi(exp1.getString())
				return &NumberObject{Value: -num1}, Jsp.JavaScriptLexerNUMBER
			} else {
				return nil, -1
			}
		}
		if op == "++" {
			if typ1 == Jsp.JavaScriptLexerNUMBER {
				num1, _ := strconv.Atoi(exp1.getString())
				num := &NumberObject{Value: num1 + 1}
				name := fmt.Sprintf("%v", exp.GetChild(0).GetChild(0).GetChild(0))
				vl.setVar(name, num, Jsp.JavaScriptLexerNUMBER)
				return num, Jsp.JavaScriptLexerNUMBER
			} else {
				return nil, -1
			}
		}
		if op == "--" {
			if typ1 == Jsp.JavaScriptLexerNUMBER {
				num1, _ := strconv.Atoi(exp1.getString())
				num := &NumberObject{Value: num1 - 1}
				name := fmt.Sprintf("%v", exp.GetChild(0).GetChild(0).GetChild(0))
				vl.setVar(name, num, Jsp.JavaScriptLexerNUMBER)
				return num, Jsp.JavaScriptLexerNUMBER
			} else {
				return nil, -1
			}
		}
	}

	if exp.GetChildCount() == 1 { // this | cobj | funcall | num | id | str | arr | obj
		if call, ok := exp.GetChild(0).(*Jsp.FuncallContext); ok { // funcall
			return h.handExpr_funcall(call)
		}
		if _, ok := exp.GetChild(0).(*Jsp.ThisContext); ok { // this
			return &ThisObjetc{}, Jsp.JavaScriptParserRULE_this
		}

		if cobj, ok := exp.GetChild(0).(*Jsp.CobjContext); ok { // class object
			return h.handExpr_cobj(cobj)
		} else if arr, ok := exp.GetChild(0).(*Jsp.ArrContext); ok { // array
			arrt := &ArrayObject{Value: make([]BaseObject, 0)}
			for _, a := range arr.GetChildren() {
				if a.GetChild(0) != nil {
					num, _ := strconv.Atoi(fmt.Sprintf("%v", a.GetChild(0).GetChild(0)))
					arrt.setValue(&NumberObject{Value: num})
				}
			}
			return arrt, Jsp.JavaScriptParserRULE_arr
		} else if obj, ok := exp.GetChild(0).(*Jsp.ObjContext); ok { // object
			objt := &ObjObject{Value: make(map[string]BaseObject)}
			for _, o := range obj.GetChildren() {
				if o.GetChild(0) != nil {
					name := fmt.Sprintf("%v", o.GetChild(0).GetChild(0))
					value, _ := h.handExpr(o.GetChild(2).(*Jsp.ExprContext))
					objt.addValue(name, value)
				}
			}
			return objt, Jsp.JavaScriptParserRULE_obj
		} else {
			str := fmt.Sprintf("%v", exp.GetChild(0).GetChild(0))
			if str[0] == '"' { // string
				return &StringObject{Value: str}, Jsp.JavaScriptLexerSTRING
			}
			if renum.MatchString(str) { // number
				num, _ := strconv.Atoi(str)
				return &NumberObject{Value: num}, Jsp.JavaScriptLexerNUMBER
			}
			if resym.MatchString(str) { // var
				value, typ := vl.getVarByName(str)
				return value, typ
			}
		}
	}

	return nil, -1
}

func (h *Handler) checkExpr(vaule BaseObject, typ int) bool {
	if str, ok := vaule.(*StringObject); ok {
		if typ == Jsp.JavaScriptLexerSTRING && str.Value != "" {
			return true
		}
	}
	if num, ok := vaule.(*NumberObject); ok {
		if typ == Jsp.JavaScriptLexerNUMBER && num.Value != 0 {
			return true
		}
	}
	if bol, ok := vaule.(*BooleanObject); ok {
		return bol.Value
	}

	return false
}
