package main

import (
    "fmt"
    "strconv"
    "./structs"
    "./parser"
    "github.com/antlr/antlr4/runtime/Go/antlr"
)


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
            stype := structs.TypeFromString[c.GetText()]

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


    if l.debug {
        l.symtable.Print()
        fmt.Println()
    }

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
    l.ret_type = structs.TypeFromString[c.Scalar().GetText()]
}
