/*
    symbol_table.go contains the implementation for a table to keep track of
    declared symbols on a specific scope
*/

package structs

import "fmt"

// type enumeration to keep track of symbol types

type TypeEnum int

const (
    Int_t TypeEnum = iota
    Float_t
    Bool_t
    Tensor_t
    Proc_t
)

func (t TypeEnum) String() string {
    return []string {"int", "float", "bool", "tensor", "proc"}[t]
}

func TypeFromString(s string) TypeEnum {
    return map[string]TypeEnum {
        "int": Int_t,
        "float": Float_t,
        "bool": Bool_t,
        "tensor": Tensor_t,
        "proc": Proc_t,
    }[s]
}

var ttoa = map[TypeEnum]string {
}

// symbol table implementation

type symbolRegistry struct {
    stype   TypeEnum
    local   map[string]symbolRegistry
}

func (s *symbolRegistry) print() {
    fmt.Printf("%10s ", s.stype.String())

    if s.local == nil {
        return

    } else {
        fmt.Printf("\n")

        for key, registry := range s.local {
            fmt.Printf("\t%20s ", key)
            registry.print()
            fmt.Printf("\n")
        }
    }
}

type SymbolTable struct {
    data    map[string]symbolRegistry
}

func NewSymbolTable() SymbolTable {
    return SymbolTable{data: make(map[string]symbolRegistry)}
}

func (s *SymbolTable) Print() {
    for key, registry := range s.data {
        fmt.Printf("%20s ", key)
        registry.print()
        fmt.Printf("\n")
    }
}

func (s *SymbolTable) Insert(scope, name string, stype TypeEnum) bool {
    var exists bool

    if scope == "global" {
        _, exists := s.data[name]

        if !exists {
            if stype == Proc_t {
                s.data[name] = symbolRegistry{
                    stype: stype,
                    local: make(map[string]symbolRegistry),
                }
            } else {
                s.data[name] = symbolRegistry{stype: stype}
            }
        }

    } else {
        _, exists := s.data[scope].local[name]

        if !exists {
            s.data[scope].local[name] = symbolRegistry{stype: stype}
        }
    }

    return exists
}

func (s *SymbolTable) Exists(scope, name string) bool {
    var exists bool

    if scope == "global" {
        _, exists = s.data[name]
    } else {
        _, exists = s.data[scope].local[name]
    }

    return exists
}
