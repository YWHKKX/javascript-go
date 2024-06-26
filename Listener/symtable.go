package listener

import (
	"fmt"
	"strings"

	Jsp "github.com/klang/js-go/JavaScript"
)

var vl *SymTable

type SymTable struct {
	Fuctions   []Fuction
	Classes    []*Class
	InFuctions inFunction
	scope      *Scope
}

type Variant struct {
	name  string
	value BaseObject
	typ   int
}

type Fuction struct {
	name string
	args []string
	vars []string
	fctx *Jsp.BlockContext
}

type inFunction struct {
	log func(string, bool) int
	len func(string) int
}

type Class struct {
	name  string
	vars  []string
	funcs map[string]*Fuction
}

func (p *SymTable) enterScope() {
	p.scope = NewScope(p.scope)
}

func (p *SymTable) leaveScope(scope *Scope) {
	p.scope = scope
}

func NewSymTable() *SymTable {
	sym := &SymTable{
		Fuctions:   []Fuction{},
		InFuctions: inFunction{},
		scope:      NewScope(nil),
	}
	sym.initBuildInFunc()

	return sym
}

func NewVariant(name string, value BaseObject, typ int) *Variant {
	return &Variant{name, value, typ}
}

func (vl *SymTable) showVarAll() {
	fmt.Println("|----------SymTable--------------------------------------|")
	fmt.Println("|name \t\t\t type \t\t value \t\t |")
	fmt.Println("|--------------------------------------------------------|")

	s := vl.scope
	for s != nil {
		for _, v := range s.Objects {
			fmt.Printf("|%-24v %-14v %-15v\n", v.name, v.typ, v.value.getString())
		}
		s = s.Outer
	}

	fmt.Printf("|--------------------------------------------------------|\n\n")
}

func (vl *SymTable) addVar(name string, value BaseObject, typ int) {
	vl.scope.Insert(NewVariant(current+"."+name, value, typ))
	vl.setFuncByName(current, name)
}

func (vl *SymTable) setVar(name string, value BaseObject, typ int) {
	substrs := strings.Split(name, "@")
	if _, obj := vl.scope.Lookup(current + "." + substrs[0]); obj != nil {
		if typ == Jsp.JavaScriptParserRULE_obj {
			if o, ok := obj.value.(*ObjObject); ok {
				o.setValue(substrs[1], value)
			}
		} else {
			obj.value = value
			obj.typ = typ
		}
	} else {
		vl.addVar(substrs[0], value, typ)
	}
}

func (vl *SymTable) delVarAll() {
	for _, v := range vl.scope.Objects {
		substrs := strings.Split(v.name, ".")
		if substrs[0] == current {
			delete(vl.scope.Objects, v.name)
		}
	}
}

func (vl *SymTable) getVarByName(name string) (BaseObject, int) {
	if _, obj := vl.scope.Lookup(current + "." + name); obj != nil {
		return obj.value, obj.typ
	}
	return nil, -1
}

func NewFunction(name string, args []string, ctx *Jsp.BlockContext) *Fuction {
	var vars []string

	return &Fuction{name, args, vars, ctx}
}

func (vl *SymTable) addFunc(name string, args []string, ctx *Jsp.BlockContext) {
	vl.Fuctions = append(vl.Fuctions, *NewFunction(name, args, ctx))
}

func (vl *SymTable) getFunc(name string) ([]string, *Jsp.BlockContext) {
	for _, f := range vl.Fuctions {
		if f.name == name {
			return f.args, f.fctx
		}
	}

	return nil, nil
}

func (vl *SymTable) setFuncByName(name, arg string) {
	var i int = 0
	for i < len(vl.Fuctions) {
		if vl.Fuctions[i].name == name {
			vl.Fuctions[i].vars = append(vl.Fuctions[i].vars, arg)
		}
		i++
	}
}

func (vl *SymTable) showFuncAll() {
	fmt.Println("|----------Function------------------------------|")
	fmt.Println("|name \t\t args \t\t vars \t\t |")
	fmt.Printf("|------------------------------------------------|\n")
	for _, f := range vl.Fuctions {
		fmt.Printf("|%-15v[ ", f.name)
		for _, a := range f.args {
			fmt.Printf("%v ", a)
		}
		fmt.Printf("]\t\t[ ")
		for _, v := range f.vars {
			fmt.Printf("%v ", v)
		}
		fmt.Println("]")
	}
	fmt.Printf("|------------------------------------------------|\n\n")
}

func (vl *SymTable) initBuildInFunc() {
	vl.InFuctions.log = func(msg string, key bool) int {
		if key {
			fmt.Printf("[+]log: %v\n", msg)
		} else {
			fmt.Printf("[-]err: %v\n", msg)
		}
		return 0
	}
	vl.InFuctions.len = func(msg string) int {
		return len(msg) - 2
	}
}

func (vl *SymTable) callInFunc(name string, args []Variant) (BaseObject, int) {
	if name == "log" {
		if len(args) == 1 {
			vl.InFuctions.log(args[0].value.getString(), true)
		} else {
			vl.InFuctions.log("The parameters are incorrect", false)
		}
	}
	if name == "len" {
		return &NumberObject{Value: vl.InFuctions.len(args[0].value.getString())}, Jsp.JavaScriptLexerNUMBER
	}
	return nil, -1
}

func NewClass(name string) *Class {
	return &Class{name: name,
		vars:  []string{},
		funcs: map[string]*Fuction{},
	}
}

func (vl *SymTable) getClass(name string) (*Class, bool) {
	for _, c := range vl.Classes {
		if name == c.name {
			return c, true
		}
	}
	return nil, false
}

func (vl *SymTable) addClass(class *Class) {
	vl.Classes = append(vl.Classes, class)
}

func (c *Class) addVar(name string) {
	c.vars = append(c.vars, name)
}

func (c *Class) initFunc(m *Jsp.ConsContext) {
	currentT := current
	defer func() {
		current = currentT
	}()

	name := fmt.Sprintf("%v", m.GetChild(0))
	current = name

	params := m.GetChild(2)
	var args []string
	for _, param := range params.GetChildren() {
		if param.GetChild(0) != nil {
			arg := fmt.Sprintf("%v", param.GetChild(0))
			args = append(args, arg)
		}
	}
	method := NewFunction(name, args, m.GetChild(4).(*Jsp.BlockContext))
	c.funcs[name] = method
}

func (c *Class) addFunc(m *Jsp.MethodContext) {
	currentT := current
	defer func() {
		current = currentT
	}()

	name := fmt.Sprintf("%v", m.GetChild(0))
	current = name

	params := m.GetChild(2)
	var args []string
	for _, param := range params.GetChildren() {
		if param.GetChild(0) != nil {
			arg := fmt.Sprintf("%v", param.GetChild(0))
			args = append(args, arg)
		}
	}
	method := NewFunction(name, args, m.GetChild(4).(*Jsp.BlockContext))
	c.funcs[name] = method
}

func (c *Class) getFunc(name string) (*Fuction, bool) {
	if _, ok := c.funcs[name]; ok {
		return c.funcs[name], true
	}
	return nil, false
}
