package object

import "fmt"

type Boolean struct {
	Value bool
}

func (b *Boolean) ToString() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}
