package main

import (
	"fmt"
	"strings"
)

type Operand interface {
	String() string
}

type Register struct {
	Name string
}

type Imm struct {
	Value int
}

func (r Register) String() string {
	return "%eax"
}

func (i Imm) String() string {
	return fmt.Sprintf("$%d", i.Value)
}

type Instruction interface {
	String() string
}

type Mov struct {
	Src Operand
	Dst Operand
}

func (m Mov) String() string {
	return fmt.Sprintf("	movl %s, %s", m.Src.String(), m.Dst.String())
}

type Ret struct{}

func (r Ret) String() string {
	return "	ret"
}

type Func struct {
	Name string
	Inst []Instruction
}

func (f Func) String() string {
	var builder strings.Builder
	builder.WriteString("	.globl " + f.Name)
	builder.WriteString("\n")
	builder.WriteString(f.Name + ":\n")
	for _, inst := range f.Inst {
		builder.WriteString(inst.String() + "\n")
	}
	return builder.String()
}

type Prog struct {
	Func
}

func (p Prog) String() string {
	var builder strings.Builder
	builder.WriteString(p.Func.String())
	builder.WriteString("\n")
	builder.WriteString(`	.ident	"GCC: (x86_64-posix-seh-rev0, Built by MinGW-W64 project) 8.1.0"` + "\n")
	return builder.String()
}
