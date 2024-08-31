package object

import "strconv"

type Number struct {
	Value float64
}

func (n *Number) ToString() string {
	return strconv.FormatFloat(n.Value, 'f', -1, 64)
}

func (n *Number) Type() ObjectType {
	return NUMBER_OBJ
}
