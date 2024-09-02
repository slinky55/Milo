package object

import "github.com/slinky55/milo/ast"

type Function struct {
	stmts  []ast.Statement
	params []string
}

func NewFunction(stmts []ast.Statement, params []*ast.IdentExpr) *Function {
	fn := &Function{
		stmts: stmts,
	}
	for _, param := range params {
		fn.params = append(fn.params, param.Value)
	}
	return fn
}

func (f *Function) ToString() string { return "function" }
func (f *Function) Type() ObjectType { return FUNC_OBJ }
func (f *Function) Value() any       { return f.stmts }
