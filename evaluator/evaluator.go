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
			if value != nil {
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
			println("unexpected statement: ", t.Literal())
		}
	}
}

func (e *Evaluator) evalExpression(node ast.Expression) (object.Object, error) {
	switch expr := node.(type) {
	case *ast.NumberExpr:
		return object.NewNumber(expr.Value), nil
	case *ast.BooleanExpr:
		return object.NewBoolean(expr.Value), nil
	case *ast.IdentExpr:
		value, ok := e.variables[expr.Value]
		if !ok {
			return nil, fmt.Errorf("invalid reference: %s is nil", expr.Value)
		}
		return value, nil
	case *ast.StringExpr:
		return object.NewString(expr.Value), nil
	case *ast.FunctionExpr:
		return object.NewFunction(expr.Body.Statements, expr.Parameters), nil
	case *ast.PrefixExpression:
		return e.evalPrefixExpression(expr)
	case *ast.BinaryExpression:
		return e.evalBinaryExpression(expr)
	case *ast.CallExpr:
		var args []object.Object

		for _, arg := range expr.Arguments {
			value, err := e.evalExpression(arg)
			if err != nil {
				println(err.Error())
				break
			}
			args = append(args, value)
		}

		if len(args) != len(expr.Arguments) {
			break
		}

		switch idt := expr.Function.(type) {
		case *ast.IdentExpr:
			ident := idt.Value
			if fn, ok := builtins[ident]; ok {
				return fn(args...), nil
			}
		}

		return nil, fmt.Errorf("unknown function: %s", expr.Function.ToString())
	default:
		return nil, fmt.Errorf("invalid expression type: %T", expr)
	}
	return nil, nil
}

func (e *Evaluator) evalPrefixExpression(expr *ast.PrefixExpression) (object.Object, error) {
	right, err := e.evalExpression(expr.Right)
	if err != nil {
		return nil, err
	}
	switch expr.Operator {
	case "!":
		if right.Type() != object.BOOLEAN_OBJ {
			return nil, fmt.Errorf("invalid operand %s for prefix !", right.ToString())
		}
		return object.NewBoolean(!right.(*object.Boolean).Value().(bool)), nil
	case "-":
		if right.Type() != object.NUMBER_OBJ {
			return nil, fmt.Errorf("invalid operand %s for prefix -", right.ToString())
		}
		return object.NewNumber(-right.(*object.Number).Value().(float64)), nil
	case "++":
		if right.Type() != object.NUMBER_OBJ {
			return nil, fmt.Errorf("invalid operand %s for prefix ++", right.ToString())
		}
		num := right.(*object.Number)
		num.Increment()
		return num, nil
	case "--":
		if right.Type() != object.NUMBER_OBJ {
			return nil, fmt.Errorf("invalid operand %s for prefix --", right.ToString())
		}
		num := right.(*object.Number)
		num.Decrement()
		return num, nil
	default:
		return nil, fmt.Errorf("unknown prefix op: %s", expr.Operator)
	}
}

func (e *Evaluator) evalBinaryExpression(expr *ast.BinaryExpression) (object.Object, error) {
	left, err := e.evalExpression(expr.Left)
	if err != nil {
		return nil, err
	}

	right, err := e.evalExpression(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Operator {
	case "+":
		if left.Type() != object.NUMBER_OBJ || right.Type() != object.NUMBER_OBJ {
			return nil, fmt.Errorf("invalid operand(s) for \"+\"")
		}
		return object.NewNumber(left.Value().(float64) + right.Value().(float64)), nil
	case "-":
		if left.Type() != object.NUMBER_OBJ || right.Type() != object.NUMBER_OBJ {
			return nil, fmt.Errorf("invalid operand(s) for \"-\"")
		}
		return object.NewNumber(left.Value().(float64) - right.Value().(float64)), nil
	case "*":
		if left.Type() != object.NUMBER_OBJ || right.Type() != object.NUMBER_OBJ {
			return nil, fmt.Errorf("invalid operand(s) for \"*\"")
		}
		return object.NewNumber(left.Value().(float64) * right.Value().(float64)), nil
	case "/":
		if left.Type() != object.NUMBER_OBJ || right.Type() != object.NUMBER_OBJ {
			return nil, fmt.Errorf("invalid operand(s) for \"/\"")
		}
		return object.NewNumber(left.Value().(float64) / right.Value().(float64)), nil
	case ">":
		if left.Type() != object.NUMBER_OBJ || right.Type() != object.NUMBER_OBJ {
			return nil, fmt.Errorf("invalid operand(s) for \">\"")
		}
		return object.NewBoolean(left.Value().(float64) > right.Value().(float64)), nil
	case "<":
		if left.Type() != object.NUMBER_OBJ || right.Type() != object.NUMBER_OBJ {
			return nil, fmt.Errorf("invalid operand(s) for \"<\"")
		}
		return object.NewBoolean(left.Value().(float64) < right.Value().(float64)), nil
	default:
		return nil, fmt.Errorf("invalid operator for binary expression %s", expr.Operator)
	}
}
