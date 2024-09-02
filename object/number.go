package object

import "strconv"

type Number struct {
	value float64
}

func NewNumber(val float64) *Number { return &Number{value: val} }

func (n *Number) ToString() string {
	return strconv.FormatFloat(n.value, 'f', -1, 64)
}

func (n *Number) Type() ObjectType {
	return NUMBER_OBJ
}

func (n *Number) Value() any { return n.value }

func (n *Number) Increment() {
	n.value++
}

func (n *Number) Decrement() {
	n.value--
}
