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

// variables to keep semantic analysis state

var symtable = structs.NewSymTable()
var scope = structs.Global
var symqueue structs.Queue
var dim structs.Queue
var var_decl bool

// BigDuckListener structure

type BigDuckListener struct {
    valid bool
    *parser.BaseBigDuckListener
}

func (this *BigDuckListener) VisitErrorNode(node antlr.ErrorNode) {
    this.valid = false
}

// variable declaration listeners

func (this *BigDuckListener) EnterVar_decl(c *parser.Var_declContext) {
    var_decl = true
    symqueue.Push(c.ID().GetText())
}

func (this *BigDuckListener) ExitVar_decl(c *parser.Var_declContext) {
    var_decl = false
}

func (this *BigDuckListener) EnterNextVar(c *parser.NextVarContext) {
    if c.ID() != nil {
        symqueue.Push(c.ID().GetText())
    }
}

// type declaration listener

func (this *BigDuckListener) EnterVar_type(c *parser.Var_typeContext) {
    //fmt.Println(c.GetText())
}

func (this *BigDuckListener) EnterDim(c *parser.DimContext) {
    if var_decl {
        n, err := strconv.Atoi(c.Num_expr().GetText())

        if err != nil {
            this.valid = false
            s := c.GetStart()
            fmt.Printf(
                "line %d:%d tensor dimension must be constant\n",
                s.GetLine(), s.GetColumn())

        } else {
            dim.Push(n)
        }
    }
}

func (this *BigDuckListener) EnterScalar(c *parser.ScalarContext) {
    var var_dim []int

    if dim.Empty() {
        var_dim = append(var_dim, 1)
    } else {

        for !dim.Empty() {
            item, _ := dim.Pop()

            if n, ok := item.(int); ok {
                var_dim = append(var_dim, n)
            }
        }
    }

    for !symqueue.Empty() {
        item, _ := symqueue.Pop()
        name, _ := item.(string)

        _, _, exists := symtable.Lookup(name)

        if exists {
            this.valid = false
            s := c.GetStart()
            fmt.Printf(
                "line %d:%d duplicate symbol %s\n",
                s.GetLine(), s.GetColumn(), name)

        } else {
            stype := structs.TypeFromString(c.GetText())
            new_sym := structs.Symbol {
                Stype: stype,
                Dim: var_dim,
            }

            symtable.Insert(scope, name, new_sym)
        }
    }
}

// procedure declaration listener

func (this *BigDuckListener) EnterProc_decl(c *parser.Proc_declContext) {
    scope = structs.Local
}

func (this *BigDuckListener) ExitProc_decl(c *parser.Proc_declContext) {
    symtable.Print()
    symtable.ClearLocalScope()
    scope = structs.Global
}

func (this *BigDuckListener) EnterSign(c *parser.SignContext) {
    //fmt.Println(c.ID())
}

// procedure arguments declaration listener

func (this *BigDuckListener) EnterArgs(c *parser.ArgsContext) {
    if c.ID() != nil {
        //fmt.Println(c.ID())
    }
}
