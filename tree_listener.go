/*
    symbol_analysis.go uses BigDuckListener and a symbol table to perform the
    static analysis for declared symbols
*/

package main

import (
    "./structs"
    "./parser"
)

// BigDuckListener structure

type BigDuckListener struct {
    *parser.BaseBigDuckListener
    // flag to know whether the input is well formed or not
    valid       bool
    debug       bool

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
    jmpstack    structs.Stack

    pc          int
    tmpc        int
}
