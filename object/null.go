package object

var NULL = &Null{}

type Null struct {
}

func (n *Null) ToString() string { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Value() any       { return nil }
