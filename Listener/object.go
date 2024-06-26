package listener

import "fmt"

type BaseObject interface {
	getString() string
}

type UndefObject struct {
	BaseObject
	Value string
}

type ThisObjetc struct {
	BaseObject
	Value  string
	Object BaseObject
}

type StringObject struct {
	BaseObject
	Value string
}

type NumberObject struct {
	BaseObject
	Value int
}

type BooleanObject struct {
	BaseObject
	Value bool
}

type ArrayObject struct {
	BaseObject
	Value []BaseObject
}

type ObjObject struct {
	BaseObject
	Value map[string]BaseObject
}

type VariantObject struct {
	BaseObject
	Value string
}

type CobjObject struct {
	BaseObject
	ClassName string
	Value     map[string]BaseObject
}

func (n *UndefObject) getString() string {
	return "undefined"
}

func (t *ThisObjetc) setObject(obj BaseObject) {
	t.Object = obj
}
func (t *ThisObjetc) setValue(value string) {
	t.Value = value
}

func (s *StringObject) getString() string {
	return s.Value
}

func (n *NumberObject) getString() string {
	return fmt.Sprintf("%d", n.Value)
}

func (b *BooleanObject) getString() string {
	if b.Value {
		return "true"
	}
	return "false"
}

func (a *ArrayObject) getString() string {
	str := "["
	for _, n := range a.Value {
		str += fmt.Sprintf(" %v,", n.getString())
	}
	str = str[:len(str)-1] + " ]"

	return str
}

func (a *ArrayObject) setValue(num BaseObject) {
	a.Value = append(a.Value, num)
}

func (a *ArrayObject) getValue(index int) BaseObject {
	return a.Value[index]
}

func (o *ObjObject) getString() string {
	str := "{"
	for k, v := range o.Value {
		str += fmt.Sprintf(" %s:%s,", k, v.getString())
	}
	str = str[:len(str)-1] + " }"

	return str
}

func (o *ObjObject) addValue(name string, obj BaseObject) {
	o.Value[name] = obj
}

func (o *ObjObject) setValue(name string, obj BaseObject) {
	for i, _ := range o.Value {
		if i == name {
			o.Value[i] = obj
		}
	}
}

func (v *VariantObject) getString() string {
	return v.Value
}

func (o *CobjObject) getString() string {
	str := "{"
	for k, v := range o.Value {
		str += fmt.Sprintf(" %s:%s,", k, v.getString())
	}
	str = str[:len(str)-1] + " }"

	return str
}

func (o *CobjObject) addValue(name string, obj BaseObject) {
	o.Value[name] = obj
}

func (o *CobjObject) setValue(name string, obj BaseObject) {
	for i, _ := range o.Value {
		if i == name {
			o.Value[i] = obj
		}
	}
}
