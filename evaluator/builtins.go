package evaluator

import "github.com/slinky55/milo/object"

type Builtin func(...object.Object) object.Object

var builtins = map[string]Builtin{
	"print": Print,
}

func Print(args ...object.Object) object.Object {
	println(args[0].ToString())
	return nil
}
