/*
    symbol_table.go contains the implementation for a table to keep track of
    declared symbols on a specific scope
*/

package structs

import "fmt"

// type enumeration to keep track of symbol types

const (
    Int_t = iota
    Float_t
    Bool_t
    Tensor_t
    Proc_t
)

func TypeToString(t int) string {
    return []string {"int", "float", "bool", "tensor", "proc"}[t]
}

func TypeFromString(s string) int {
    return map[string]int {
        "int": Int_t,
        "float": Float_t,
        "bool": Bool_t,
        "tensor": Tensor_t,
        "proc": Proc_t,
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
            make(map[string]Symbol),
        },
    }
}

func (s SymTable) Print() {
    for scope := 0; scope < scopeEnumCount; scope++ {
        fmt.Printf("=== %s ===\n", ScopeToString(scope))

        for symbol, _ := range s.table[scope] {
            fmt.Println(symbol)
        }

        fmt.Println()
    }
}

func (s *SymTable) Insert(scope int, name string, sym Symbol) {
    s.table[scope][name] = sym
}

func (s *SymTable) Lookup(name string) (scope int, sym Symbol, exists bool) {
    for scope = 0; scope < scopeEnumCount; scope++ {
        sym, exists = s.table[scope][name]

        if exists {
            break
        }
    }

    return
}

func (s *SymTable) ClearLocalScope() {
    for key, _ := range s.table[0] {
        delete(s.table[0], key)
    }
}
