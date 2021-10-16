package main

import (
    "fmt"
    "strconv"
    "./structs"
    "./parser"
)

//  generation listeners

func (l *BigDuckListener) TopOp() int {
    item := l.opstack.Top()
    top, _ := item.(int)
    return top
}

func (l *BigDuckListener) PopOp() int {
    item, _ := l.opstack.Pop()
    op, _ := item.(int)
    return op
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

func (l *BigDuckListener) generateTAC() {
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

// IR generation listeners

func (l *BigDuckListener) ExitProgram(c *parser.ProgramContext) {
    for index, code := range l.ir_code {
        fmt.Printf("%d\t", index);
        code.Print()
    }

    fmt.Println()
}

// Assignments

func (l *BigDuckListener) EnterAssignment(c *parser.AssignmentContext) {
    l.argstack.Push(c.ID().GetText())
    l.PushOp(structs.OpFromString["<-"])
}

func (l *BigDuckListener) ExitAssignment(c *parser.AssignmentContext) {
    l.generateTAC()
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
        l.generateTAC()
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
        l.generateTAC()
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
        l.generateTAC()
    }
}

// Rel ops

func (l *BigDuckListener) EnterOpRel(c *parser.OpRelContext) {
    l.PushOp(structs.OpFromString[c.RelOp().GetText()])
}

func (l *BigDuckListener) ExitRel_expr(c *parser.Rel_exprContext) {
    l.generateTAC()
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
        l.generateTAC()
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
        l.generateTAC()
    }
}

// if statements

func (l *BigDuckListener) EnterBodyCond(c *parser.BodyCondContext) {
    /*
    l.ir_code = append(
        l.ir_code,
        structs.Tac{Op: op, Arg1: args[0], Arg2: args[1], Target: target})
    l.jmpstack.Push(l.pc - 1)
    */
}
