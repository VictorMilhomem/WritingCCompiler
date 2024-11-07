package main

type CodeGenerator struct {
}

func (c *CodeGenerator) Generate(ast AST) string {
	prog := c.convertProgram(ast.Program)
	return prog.String()
}

func (c *CodeGenerator) convertProgram(prog Program) Prog {
	fn := c.convertFunction(prog.Function)
	return Prog{fn}
}

func (c *CodeGenerator) convertFunction(f Function) Func {
	name := f.Name
	stmt := f.Body.(Return)
	mov, ret := c.convertStatement(stmt)
	var instructions []Instruction = []Instruction{mov, ret}
	return Func{
		name,
		instructions,
	}
}

func (c *CodeGenerator) convertStatement(stmt Return) (Mov, Ret) {
	expr, _ := stmt.Expression.(NumberInteger)
	val := c.convertExpr(expr)

	return Mov{
		val,
		Register{},
	}, Ret{}
}

func (c *CodeGenerator) convertExpr(val NumberInteger) Imm {
	return Imm{Value: val.Value}
}
