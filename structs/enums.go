/*
    enums.go contains all enums used inside structs package
*/

package structs

// type enumeration to keep track of symbol types

const (
    Error_t = iota
    Void_t
    Int_t
    Float_t
    Bool_t
    Tensor_t
    Proc_t
    typeEnumCount
)

var TypeToString [typeEnumCount]string = [typeEnumCount]string {
    "error", "void", "int", "float", "bool", "tensor", "proc",
}

var TypeFromString map[string]int = map[string]int {
    "error" : Error_t,
    "void"  : Void_t,
    "int"   : Int_t,
    "float" : Float_t,
    "bool"  : Bool_t,
    "tensor": Tensor_t,
    "proc"  : Proc_t,
}

// scope enumeration to keep track of symbol scope

const (
    Local = iota
    Global
    scopeEnumCount
)

var ScopeToString [scopeEnumCount]string = [scopeEnumCount]string {
    "local", "global",
}

/* 
    operator enumeration to define operators, they are ordered by evaluation
    priority
*/


const (
    NOP = iota  // No operation

    ASG         // assignment

    OR          // logical or
    AND         // logical and
    NOT         // logical not

    EQ          // equal
    NEQ         // not equal

    LES         // less than
    GRE         // greater than
    LEQ         // less or equal
    GEQ         // greater or equal

    SUB         // substraction
    ADD         // addition
    DIV         // division
    MUL         // multiplication

    LPAREN      // left parenthesis
    RPAREN      // right parenthesis

    JMP         // unconditional jump
    JMT         // jump on true
    JMF         // jump on false

    opEnumCount
)

var OpToString [opEnumCount]string = [opEnumCount]string{
    "NOP",
    "ASG",
    "OR", "AND", "NOT",
    "EQ", "NEQ",
    "LES", "GRE", "LEQ", "GEQ",
    "SUB", "ADD", "DIV", "MUL",
    "LPA", "RPA",
    "JMP", "JMT", "JMF",
}

var OpFromString map[string]int = map[string]int {
    "<-"    : ASG,
    "or"    : OR,
    "and"   : AND,
    "not"   : NOT,
    "="     : EQ,
    "/="    : NEQ,
    "<"     : LES,
    ">"     : GRE,
    "<="    : LEQ,
    ">="    : GEQ,
    "-"     : SUB,
    "+"     : ADD,
    "/"     : DIV,
    "*"     : MUL,
    "LPA"   : LPAREN,
    "RPA"   : RPAREN,
    "JMP"   : JMP,
    "JMT"   : JMT,
    "JMF"   : JMF,
}

//enumeration to keep track of loop style

const (
    InfStyle = iota
    ForStyle
    WhileStyle
)
