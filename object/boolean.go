package object

import "fmt"

type Boolean struct {
	value bool
}

func NewBoolean(val bool) *Boolean { return &Boolean{value: val} }

func (b *Boolean) ToString() string {
	return fmt.Sprintf("%t", b.value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) Value() any {
	return b.value
}
