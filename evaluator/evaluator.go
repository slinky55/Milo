package evaluator

import (
	"fmt"
	"github.com/slinky55/milo/ast"
	"github.com/slinky55/milo/object"
)

type Evaluator struct {
	Program *ast.Program

	variables map[string]object.Object
}

func New(program *ast.Program) *Evaluator {
	return &Evaluator{
		Program:   program,
		variables: make(map[string]object.Object),
	}
}

func (e *Evaluator) Evaluate() {
	for _, stmt := range e.Program.Statements {
		switch t := stmt.(type) {
		case *ast.ExpressionStatement:
			value, err := e.evalExpression(t.Expr)
			if err != nil {
				println(err.Error())
				continue
			}
			if value.Type() != object.FUNC_OBJ {
				println(value.ToString())
			}
		case *ast.LetStatement:
			value, err := e.evalExpression(t.Expr)
			if err != nil {
				println(err.Error())
				continue
			}

			if value != nil {
				e.variables[t.Ident.Value] = value
			}
		default:
			println("statement type not supported")
		}
	}
}

func (e *Evaluator) evalExpression(node ast.Expression) (object.Object, error) {
	switch expr := node.(type) {
	case *ast.NumberExpr:
		return &object.Number{Value: expr.Value}, nil
	case *ast.BooleanExpr:
		return &object.Boolean{Value: expr.Value}, nil
	case *ast.IdentExpr:
		value, ok := e.variables[expr.Value]
		if !ok {
			return nil, fmt.Errorf("invalid reference: %s is nil", expr.Value)
		}
		return value, nil
	case *ast.CallExpr:
		if fn, ok := builtins[expr.Function.ToString()]; ok {
			fn()
			return nil, nil
		}
		return nil, nil
	default:
		return nil, fmt.Errorf("invalid expression type: %T", expr)
	}
}
