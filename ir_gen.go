package main

import (
    "fmt"
    "strconv"
    "./structs"
    "./parser"
)

//  generation listeners

func (l *BigDuckListener) TopOp() int {
    if l.opstack.Empty() {
        return structs.NOP
    } else {
        item := l.opstack.Top()
        top, _ := item.(int)
        return top
    }
}

func (l *BigDuckListener) PopOp() int {
    if l.opstack.Empty() {
        return structs.NOP
    } else {
        item, _ := l.opstack.Pop()
        top, _ := item.(int)
        return top
    }
}

func (l *BigDuckListener) PushOp(op int) {
    if l.opstack.Empty() {
        l.opstack.Push(op)
    } else {
        top := l.TopOp()

        if top == structs.LPAREN {
            if op == structs.RPAREN {
                l.opstack.Pop()
            } else {
                l.opstack.Push(op)
            }

        } else {
            l.opstack.Push(op)
        }
    }
}

func (l *BigDuckListener) GenerateOpTAC() {
    var args [2]string
    var target string
    var i int

    op := l.PopOp()

    if op == structs.NOT {
        i = 0
    } else {
        i = 1
    }

    for ; i >= 0; i-- {
        item, _ := l.argstack.Pop()
        args[i], _ = item.(string)
    }

    if op == structs.ASG {
        target = args[0]
        args[0] = args[1]
        args[1] = ""
    } else {
        target = "t" + strconv.Itoa(l.tmpc)
        l.argstack.Push(target)
        l.tmpc++
    }

    l.ir_code = append(
        l.ir_code,
        structs.Tac{Op: op, Arg1: args[0], Arg2: args[1], Target: target})
    l.pc++
}

func (l *BigDuckListener) GenerateJmpTAC(jmptype int) {
    if jmptype == structs.JMP {
        l.ir_code = append(l.ir_code, structs.Tac{Op: jmptype})
        l.pc++
    } else {
        item, _ := l.argstack.Pop()
        cond, _ := item.(string)

        l.ir_code = append(l.ir_code, structs.Tac{Op: jmptype, Arg1: cond})
        l.pc++
    }
}

func (l *BigDuckListener) FillJmpTAC(index, target int) {
    l.ir_code[index].Target = strconv.Itoa(target)
}

// IR generation listeners

func (l *BigDuckListener) ExitProgram(c *parser.ProgramContext) {
    l.ir_code = append(l.ir_code, structs.Tac{Op: structs.NOP})
    if l.debug {
        for index, code := range l.ir_code {
            fmt.Printf("%3d ", index);
            code.Print()
        }

        fmt.Println()
    }
}

// Assignments

func (l *BigDuckListener) EnterAssignment(c *parser.AssignmentContext) {
    l.argstack.Push(c.ID().GetText())
    l.PushOp(structs.OpFromString["<-"])
}

func (l *BigDuckListener) ExitAssignment(c *parser.AssignmentContext) {
    l.GenerateOpTAC()
}

// Bool exprs

func (l *BigDuckListener) EnterNextBool(c *parser.NextBoolContext) {
    if c.GetText() != "" {
        l.PushOp(structs.OpFromString["or"])
    }
}

// Prod exprs

func (l *BigDuckListener) EnterNextAnd(c *parser.NextAndContext) {
    if c.GetText() != "" {
        l.PushOp(structs.OpFromString["and"])
    }
}

func (l *BigDuckListener) ExitAnd_expr(c *parser.And_exprContext) {
    if l.TopOp() == structs.OR {
        l.GenerateOpTAC()
    }
}

// Not exprs

func (l *BigDuckListener) EnterNot_expr(c *parser.Not_exprContext) {
    if c.NOT() != nil {
        l.PushOp(structs.OpFromString["not"])
    }
}

func (l *BigDuckListener) ExitNot_expr(c *parser.Not_exprContext) {
    if c.NOT() != nil {
        l.GenerateOpTAC()
    }
}

// Bool terms

func (l *BigDuckListener) EnterBool_term(c *parser.Bool_termContext) {
    if c.Bool_expr() != nil {
        l.PushOp(structs.LPAREN)
    }
}

func (l *BigDuckListener) ExitBool_term(c *parser.Bool_termContext) {
    if c.Bool_expr() == nil {
        if c.ID() != nil {
            l.argstack.Push(c.ID().GetText())

        } else if c.TRUE() != nil {
            l.argstack.Push("#t")

        } else if c.FALSE() != nil {
            l.argstack.Push("#f")
        }

    } else {
        l.PushOp(structs.RPAREN)
    }

    top := l.TopOp()

    if top == structs.AND {
        l.GenerateOpTAC()
    }
}

// Rel ops

func (l *BigDuckListener) EnterOpRel(c *parser.OpRelContext) {
    l.PushOp(structs.OpFromString[c.RelOp().GetText()])
}

func (l *BigDuckListener) ExitRel_expr(c *parser.Rel_exprContext) {
    l.GenerateOpTAC()
}

// Sum exprs

func (l *BigDuckListener) EnterNextSum(c *parser.NextSumContext) {
    if c.GetText() != "" {
        if c.GetText()[0] == '+' {
            l.PushOp(structs.OpFromString["+"])

        } else if c.GetText()[0] == '-' {
            l.PushOp(structs.OpFromString["-"])
        }
    }
}

// Prod exprs

func (l *BigDuckListener) EnterNextProd(c *parser.NextProdContext) {
    if c.GetText() != "" {
        if c.GetText()[0] == '*' {
            l.PushOp(structs.OpFromString["*"])

        } else if c.GetText()[0] == '/' {
            l.PushOp(structs.OpFromString["/"])
        }
    }
}

func (l *BigDuckListener) ExitProd_expr(c *parser.Prod_exprContext) {
    if l.TopOp() == structs.ADD || l.TopOp() == structs.SUB {
        l.GenerateOpTAC()
    }
}

// Factors

func (l *BigDuckListener) EnterFactor(c *parser.FactorContext) {
    if c.Num_expr() != nil {
        l.PushOp(structs.LPAREN)
    }
}

func (l *BigDuckListener) ExitFactor(c *parser.FactorContext) {
    if c.Num_expr() == nil {
        if c.ID() != nil {
            l.argstack.Push(c.ID().GetText())

        } else if c.CTE_INT() != nil {
            l.argstack.Push(c.CTE_INT().GetText())

        } else if c.CTE_FLOAT() != nil {
            l.argstack.Push(c.CTE_FLOAT().GetText())
        }

    } else {
        l.PushOp(structs.RPAREN)
    }

    top := l.TopOp()

    if top == structs.MUL || top == structs.DIV {
        l.GenerateOpTAC()
    }
}

// if statements

func (l *BigDuckListener) EnterBodyCond(c *parser.BodyCondContext) {
    l.jmpstack.Push(l.pc)
    l.GenerateJmpTAC(structs.JMF)
}

func (l *BigDuckListener) EnterEndIfBlock(c *parser.EndIfBlockContext) {
    item, _ := l.jmpstack.Pop()
    index, _ := item.(int)

    if c.Alter() == nil {
        l.FillJmpTAC(index, l.pc)
    } else {
        l.FillJmpTAC(index, l.pc + 1)
    }
}

func (l *BigDuckListener) EnterAlter(c *parser.AlterContext) {
    l.jmpstack.Push(l.pc)
    l.GenerateJmpTAC(structs.JMP)
}

func (l *BigDuckListener) ExitAlter(c *parser.AlterContext) {
    item, _ := l.jmpstack.Pop()
    index, _ := item.(int)
    l.FillJmpTAC(index, l.pc)
}

// loops

func (l *BigDuckListener) EnterInfLoop(c *parser.InfLoopContext) {
    l.jmpstack.Push(l.pc)
}

func (l *BigDuckListener) EnterWhileStyle(c *parser.WhileStyleContext) {
    l.jmpstack.Push(l.pc)
}

func (l *BigDuckListener) ExitWhileStyle(c *parser.WhileStyleContext) {
    l.jmpstack.Push(l.pc)
    l.GenerateJmpTAC(structs.JMF)
}

func (l *BigDuckListener) ExitLoop_stmt(c *parser.Loop_stmtContext) {
    l.GenerateJmpTAC(structs.JMP)
    item, _ := l.jmpstack.Pop()
    target, _ := item.(int)
    l.FillJmpTAC(l.pc - 1, target)
}
