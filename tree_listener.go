/*
    symbol_analysis.go uses BigDuckListener and a symbol table to perform the
    static analysis for declared symbols
*/

package main

import (
    "fmt"
    "strconv"
    "./structs"
    "./parser"
    "github.com/antlr/antlr4/runtime/Go/antlr"
)

// BigDuckListener structure

type BigDuckListener struct {
    *parser.BaseBigDuckListener
    // flag to know whether the input is well formed or not
    valid       bool

    // variables to keep semantic analysis state

    symtable    structs.SymTable

    symqueue    structs.Queue
    typequeue   structs.Queue
    dimqueue    structs.Queue

    in_decl     bool
    in_args     bool

    scope       int
    argc        int

    ret_type    int
    curr_proc   string

    // variables to generarte IR code

    ir_code     []structs.Tac

    opstack     structs.Stack
    argstack    structs.Stack
    // gotoqueue   structs.Queue

    pc          int
    tmpc        int
}

func (l *BigDuckListener) EnterProgram(c *parser.ProgramContext) {
    l.symtable = structs.NewSymTable()
    l.ret_type = structs.Void_t
    l.scope = structs.Global
    l.valid = true
}

func (l *BigDuckListener) VisitErrorNode(node antlr.ErrorNode) {
    l.valid = false
}

// variable declaration listeners

func (l *BigDuckListener) EnterVar_decl(c *parser.Var_declContext) {
    l.in_decl = true
    l.symqueue.Push(c.ID().GetText())
}

func (l *BigDuckListener) ExitVar_decl(c *parser.Var_declContext) {
    l.in_decl = false
}

func (l *BigDuckListener) EnterNextVar(c *parser.NextVarContext) {
    if c.ID() != nil {
        l.symqueue.Push(c.ID().GetText())
    }
}

func (l *BigDuckListener) EnterDim(c *parser.DimContext) {
    if l.in_decl {
        n, err := strconv.Atoi(c.Num_expr().GetText())

        if err != nil {
            l.valid = false
            fmt.Printf(
                "line %d:%d tensor dimension must be constant\n",
                c.GetStart().GetLine(), c.GetStart().GetColumn())
        } else {
            l.dimqueue.Push(n)
        }
    }
}

func (l *BigDuckListener) EnterScalar(c *parser.ScalarContext) {
    var var_dim []int

    if l.dimqueue.Empty() {
        var_dim = append(var_dim, 1)
    } else {
        for !l.dimqueue.Empty() {
            item, _ := l.dimqueue.Pop()

            if n, ok := item.(int); ok {
                var_dim = append(var_dim, n)
            }
        }
    }

    for !l.symqueue.Empty() {
        item, _ := l.symqueue.Pop()
        name, _ := item.(string)
        _, _, exists := l.symtable.Lookup(name)

        if exists {
            l.valid = false
            fmt.Printf(
                "line %d:%d duplicate symbol %s\n",
                c.GetStart().GetLine(), c.GetStart().GetColumn(), name)
        } else {
            stype := structs.TypeFromString(c.GetText())

            l.symtable.Insert(
                l.scope,
                name,
                structs.Symbol {
                    Stype: stype,
                    Dim: var_dim})

            if l.in_args {
                l.typequeue.Push(stype)
            }
        }
    }
}

// procedure declaration listener

func (l *BigDuckListener) EnterSign(c *parser.SignContext) {
    l.symtable.Insert(l.scope, c.ID().GetText(), structs.Symbol{})
    l.curr_proc = c.ID().GetText()
    l.scope = structs.Local
}

func (l *BigDuckListener) ExitProc_decl(c *parser.Proc_declContext) {
    var typeArgs []int

    for !l.typequeue.Empty() {
        item, _ := l.typequeue.Pop()
        stype, _ := item.(int)
        typeArgs = append(typeArgs, stype)
    }

    l.symtable.Update(
        l.curr_proc,
        structs.Symbol{
            Stype: structs.Proc_t,
            Dim: []int{1},
            NumArgs: l.argc,
            TypeArgs: typeArgs,
            RetType: l.ret_type},)


    l.symtable.Print()
    l.symtable.ClearLocalScope()
    l.scope = structs.Global
    l.ret_type = structs.Void_t
    l.argc = 0
}

// procedure arguments declaration listener

func (l *BigDuckListener) EnterArgs(c *parser.ArgsContext) {
    if c.ID() != nil {
        l.symqueue.Push(c.ID().GetText())
        l.in_decl = true
        l.in_args = true
        l.argc++
    }
}

func (l *BigDuckListener) ExitArgs(c *parser.ArgsContext) {
    l.in_decl = false
    l.in_args = false
}

func (l *BigDuckListener) EnterNextTypes(c *parser.NextTypesContext) {
    if c.ID() != nil {
        l.symqueue.Push(c.ID().GetText())
        l.argc++
    }
}

func (l *BigDuckListener) EnterNextArg(c *parser.NextArgContext) {
    if c.ID() != nil {
        l.symqueue.Push(c.ID().GetText())
        l.argc++
    }
}

// procedure arguments declaration listener

func (l *BigDuckListener) EnterRet_type(c *parser.Ret_typeContext) {
    l.in_decl = true
    l.ret_type = structs.TypeFromString(c.Scalar().GetText())
}





// IR generation listeners

func (l *BigDuckListener) EnterBool_term(c *parser.Bool_termContext) {
    if c.Bool_expr() == nil && c.Rel_expr() == nil {
        if c.ID() != nil {
            l.argstack.Push(c.ID().GetText())

        } else if c.TRUE() != nil {
            l.argstack.Push("#t")

        } else if c.FALSE() != nil {
            l.argstack.Push("#f")
        }
    }
}

func (l *BigDuckListener) EnterFactor(c *parser.FactorContext) {
    if c.Num_expr() == nil {
        if c.ID() != nil {
            l.argstack.Push(c.ID().GetText())

        } else if c.CTE_INT() != nil {
            l.argstack.Push(c.CTE_INT().GetText())

        } else if c.CTE_FLOAT() != nil {
            l.argstack.Push(c.CTE_FLOAT().GetText())
        }
    }
}

func (l *BigDuckListener) generateTAC() {
    item, _ := l.opstack.Pop()
    op, _ := item.(int)

    var args [2]string
    var target string
    var i int

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
    l.ir_code[l.pc].Print()
    l.pc++
}

func (l *BigDuckListener) PushOp(op int) {
    if l.opstack.Empty() {
        l.opstack.Push(op)
    } else {
        item := l.opstack.Top()
        top, _ := item.(int)

        if structs.OperatorPriority(top) < structs.OperatorPriority(op) {
            l.opstack.Push(op)
        } else {
            l.generateTAC()
            l.opstack.Push(op)
        }
    }
}

func (l *BigDuckListener) EnterAssignment(c *parser.AssignmentContext) {
    fmt.Println(c.GetText())
    l.argstack.Push(c.ID().GetText())
    l.PushOp(structs.OperatorFromString("<-"))
}

func (l *BigDuckListener) ExitAssignment(c *parser.AssignmentContext) {
    for !l.opstack.Empty() {
        l.generateTAC()
    }
}

func (l *BigDuckListener) EnterNextBool(c *parser.NextBoolContext) {
    if c.GetText() != "" {
        l.PushOp(structs.OperatorFromString("or"))
    }
}

func (l *BigDuckListener) EnterNextAnd(c *parser.NextAndContext) {
    if c.GetText() != "" {
        l.PushOp(structs.OperatorFromString("and"))
    }
}

func (l *BigDuckListener) EnterNot_expr(c *parser.Not_exprContext) {
    if c.NOT() != nil {
        l.PushOp(structs.OperatorFromString("or"))
    }
}

func (l *BigDuckListener) EnterRelOp(c *parser.RelOpContext) {
    l.PushOp(structs.OperatorFromString(c.GetText()))
}

func (l *BigDuckListener) EnterNextSum(c *parser.NextSumContext) {
    if c.GetText() != "" {
        if c.GetText()[0] == '+' {
            l.PushOp(structs.OperatorFromString("+"))

        } else if c.GetText()[0] == '-' {
            l.PushOp(structs.OperatorFromString("-"))
        }
    }
}

func (l *BigDuckListener) EnterNextProd(c *parser.NextProdContext) {
    if c.GetText() != "" {
        if c.GetText()[0] == '*' {
            l.PushOp(structs.OperatorFromString("*"))

        } else if c.GetText()[0] == '/' {
            l.PushOp(structs.OperatorFromString("/"))
        }
    }
}
