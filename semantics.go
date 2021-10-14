package main

import (
    "./structs"
)

// variables to keep semantic analysis state

type SemanticAnalizer struct {
    symtable    structs.SymTable
    symqueue    structs.Queue
    typequeue   structs.Queue
    dimqueue    structs.Queue
    curr_proc   string
    var_decl    bool
    args        bool
    scope       int
    ret_type    int
    argc        int
}

func (s SemanticAnalizer) Init() {
    s.symtable = structs.NewSymTable()
    s.scope = structs.Global
    s.ret_type = structs.Void_t
}


func (s SemanticAnalizer) enqueueID(id string) {
    s.symqueue.Push(id)
}
