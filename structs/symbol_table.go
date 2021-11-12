/*
    symbol_table.go contains the implementation for a table to keep track of
    declared symbols on a specific scope
*/

package structs

import "fmt"

// symbol table implementation

type Symbol struct {
    Stype       int

    // for vars
    Dim         []int

    // for procs
    Argc        int
    TypeArgs    []int
    RetType     int
    Startpoint  int
    Ic          int
    Fc          int
    Bc          int
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
                "%s\t%16s %8s %+v\n",
                ScopeToString[scope],
                name,
                TypeToString[s.table[scope][name].Stype],
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
