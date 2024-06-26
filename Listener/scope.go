package listener

type Scope struct {
	Outer   *Scope
	Objects map[string]*Variant
}

func NewScope(outer *Scope) *Scope {
	return &Scope{
		Outer:   outer,
		Objects: make(map[string]*Variant),
	}
}

func (s *Scope) HasName(name string) bool {
	_, ok := s.Objects[name]
	return ok
}

func (s *Scope) Lookup(name string) (*Scope, *Variant) {
	for ; s != nil; s = s.Outer {
		if obj := s.Objects[name]; obj != nil {
			return s, obj
		}
	}
	return nil, nil
}

func (s *Scope) Insert(v *Variant) (alt *Variant) {
	if alt = s.Objects[v.name]; alt == nil {
		s.Objects[v.name] = v
	}
	return
}
