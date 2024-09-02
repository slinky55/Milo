package object

type String struct {
	value string
}

func NewString(val string) *String { return &String{value: val} }

func (s *String) ToString() string {
	return s.value
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) Value() any { return s.value }
