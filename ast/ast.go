package ast

type Node interface {
	Literal() string
	ToString() string
}
