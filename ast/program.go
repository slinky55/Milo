package ast

import "strings"

type Program struct {
	Statements []Statement
}

func (p *Program) Literal() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].Literal()
	} else {
		return ""
	}
}

func (p *Program) ToString() string {
	var out strings.Builder

	for _, stmt := range p.Statements {
		out.WriteString(stmt.Literal())
	}

	return out.String()
}

func (p *Program) AddStatement(s Statement) {
	p.Statements = append(p.Statements, s)
}
