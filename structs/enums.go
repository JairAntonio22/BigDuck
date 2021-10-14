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

/* 
    operator enumeration to define operators, they are ordered by evaluation
    priority
*/


const (
    ASG = iota  // assignment

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

    operatorEnumCount
)

func OperatorToString(op int) string {
    return []string {
	"ASG",
	"OR", "AND", "NOT",
	"EQ", "NEQ", "LES", "GRE", "LEQ", "GEQ",
	"SUB", "ADD", "DIV", "MUL",
    }[op]
}

func OperatorFromString(s string) int {
    return map[string]int {
        "<-"	: ASG,
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
    }[s]
}

/* 
    operator arity enumeration
*/

const (
    Unary = iota    // assignment
    Binary          // logical or
)

func OperatorPriority(o int) int {
    return map[int]int {
        ASG : 0,

        OR  : 1,

        AND : 2,
        NOT : 3,

        EQ  : 4,
        NEQ : 4,

        LES : 5,
        GRE : 5,
        LEQ : 5,
        GEQ : 5,

        SUB : 6,
        ADD : 6,
        DIV : 7,
        MUL : 7,
    }[o]
}
