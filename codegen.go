package main

import (
	"fmt"
)

func generateAssembly(ast AST) string {
	function := ast.Program.Function

	assembly := fmt.Sprintf("	.globl	%s\n%s:\n", function.Name, function.Name)

	assembly += generateStatementAssembly(function.Body) + "\n"

	assembly += `	.ident	"GCC: (x86_64-posix-seh-rev0, Built by MinGW-W64 project) 8.1.0"` + "\n"
	return assembly
}

func generateStatementAssembly(stmt Statement) string {
	switch s := stmt.(type) {
	case Return:
		return generateExpressionAssembly(s.Expression) + "\n	ret"
	default:
		return ""
	}
}

func generateExpressionAssembly(expr Expression) string {
	switch e := expr.(type) {
	case NumberInteger:
		return fmt.Sprintf("	movl	$%d, %%eax", e.Value)
	default:
		return ""
	}
}
