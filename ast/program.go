package ast

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

func (p *Program) AddStatement(s Statement) {
	p.Statements = append(p.Statements, s)
}
