/*
    symbol_table.go contains the implementation for a table to keep track of
    declared symbols on a specific scope
*/

package structs

import "fmt"

// type enumeration to keep track of symbol types

const (
    Void_t = iota
    Error_t = iota
    Int_t
    Float_t
    Bool_t
    Tensor_t
    Proc_t
    typeEnumCount
)

func TypeToString(t int) string {
    return []string {
        "error", "void", "int", "float", "bool", "tensor", "proc",
    }[t]
}

func TypeFromString(s string) int {
    return map[string]int {
        "error" : Error_t,
        "void"  : Void_t,
        "int"   : Int_t,
        "float" : Float_t,
        "bool"  : Bool_t,
        "tensor": Tensor_t,
        "proc"  : Proc_t,
    }[s]
}

// scope enumeration to keep track of symbol scope

const (
    Local = iota
    Global
    scopeEnumCount
)

func ScopeToString(t int) string {
    return []string {"local", "global"}[t]
}

// symbol table implementation

type Symbol struct {
    Stype       int

    // for vars
    Dim         []int

    // for procs
    NumArgs     int
    TypeArgs    []int
    RetType     int
}

type SymTable struct {
    table       [2]map[string]Symbol
}

func NewSymTable() SymTable {
    return SymTable{
        table: [2]map[string]Symbol{
            make(map[string]Symbol),
            make(map[string]Symbol)}}
}

func (s SymTable) Print() {
    fmt.Println("===== SymbolTable =====")

    for scope := 0; scope < scopeEnumCount; scope++ {
        for name, _ := range s.table[scope] {
            fmt.Printf(
                "%8s %16s %8s %v\n",
                ScopeToString(scope),
                name,
                TypeToString(s.table[scope][name].Stype),
                s.table[scope][name])
        }
    }
}

func (s *SymTable) Insert(scope int, name string, sym Symbol) {
    s.table[scope][name] = sym
}

func (s *SymTable) Lookup(name string) (int, Symbol, bool) {
    for scope := 0; scope < scopeEnumCount; scope++ {
        sym, exists := s.table[scope][name]

        if exists {
            return scope, sym, exists
        }
    }

    return 0, Symbol{}, false
}

func (s *SymTable) Update(name string, sym Symbol) {
    for scope := 0; scope < scopeEnumCount; scope++ {
        _, exists := s.table[scope][name]

        if exists {
            s.table[scope][name] = sym
        }
    }
}

func (s *SymTable) ClearLocalScope() {
    for key, _ := range s.table[0] {
        delete(s.table[0], key)
    }
}
