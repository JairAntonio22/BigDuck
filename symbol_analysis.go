/*
    symbol_analysis.go uses BigDuckListener and a symbol table to perform the
    static analysis for declared symbols
*/

package main

import (
    "./structs"
    "./parser"
    "github.com/antlr/antlr4/runtime/Go/antlr"
)

// variables to keep semantic analysis state

var symtable structs.SymbolTable = structs.NewSymbolTable()
var scope = "global"
var symqueue structs.Queue

// BigDuckListener structure

type BigDuckListener struct {
    valid bool
    *parser.BaseBigDuckListener
}

func (this *BigDuckListener) VisitErrorNode(node antlr.ErrorNode) {
    this.valid = false
}

func (this *BigDuckListener) ExitProgram(c *parser.ProgramContext) {
    symtable.Print()
}

// variable declaration listeners

func (this *BigDuckListener) EnterVar_decl(c *parser.Var_declContext) {
    symqueue.Push(c.ID().GetText())
}

func (this *BigDuckListener) EnterNextVar(c *parser.NextVarContext) {
    if c.ID() != nil {
        symqueue.Push(c.ID().GetText())
    }
}

// type declaration listener

func (this *BigDuckListener) EnterScalar(c *parser.ScalarContext) {
    // fmt.Println(c.GetText())

    curr_type := structs.TypeFromString(c.GetText())

    for !symqueue.Empty() {
        if elem, err := symqueue.Pop(); err == nil {
            if symbol, ok := elem.(string); ok {
                symtable.Insert(scope, symbol, curr_type)
            }
        }
    }
}

// procedure declaration listener

func (this *BigDuckListener) EnterSign(c *parser.SignContext) {
    scope = c.ID().GetText()
    //fmt.Println(c.ID())
}

// procedure arguments declaration listener

func (this *BigDuckListener) EnterArgs(c *parser.ArgsContext) {
    /*
    if c.ID() != nil {
        fmt.Println(c.ID())
    }
    */
}

