package object

type ObjectType string

const (
	NUMBER_OBJ  = "NUMBER"
	STRING_OBJ  = "STRING"
	BOOLEAN_OBJ = "BOOLEAN"
	FUNC_OBJ    = "FUNC"
)

type Object interface {
	Type() ObjectType
	ToString() string
}
