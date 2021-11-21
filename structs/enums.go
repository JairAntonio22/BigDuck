/*
    enums.go contains all enums used inside structs package
*/

package structs

import "errors"

// type enumeration to keep track of symbol types

const (
    Error_t = iota
    Void_t
    Int_t
    Float_t
    Bool_t
    String_t
    Pointer_t
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
    instructions enumeration to define operators and instructions used by
    virtual machine
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

    GOPROC      // go to procedure
    ERA         // indicates the size of local memory to be used
    PARAM       // assigns value to parameter passed to procedure
    RETURN      // returns value from procedures
    ENDPROC     // clears procedure call and local memory

    ASSERT      // checks whether a given value is 

    PRINT       // prints on stdout
    PRINTLN     // prints on stdout with newline character

    READ        // reads input on stdin

    SIN
    ASIN

    COS
    ACOS

    TAN
    ATAN
    ATAN2

    EXP
    LN

    SQRT
    POW
    LOG

    MOD

    ABS

    CEIL
    FLOOR

    SET         // Sets value to address
    PROGRAM     // Indicates program segment

    opEnumCount
)

var OpToString [opEnumCount]string = [opEnumCount]string{
    "NOP",
    "ASG",
    "OR", "AND", "NOT",
    "EQ", "NEQ",
    "LES", "GRE", "LEQ", "GEQ",
    "SUB", "ADD", "DIV", "MUL",
    "LPAREN", "RPAREN",
    "JMP", "JMT", "JMF",
    "GOPROC", "ERA", "PARAM", "RETURN", "ENDPROC",
    "ASSERT",
    "PRINT", "PRINTLN",
    "READ",
    "SIN", "ASIN", "COS", "ACOS", "TAN", "ATAN", "ATAN2",
    "EXP", "LN",
    "SQRT", "POW", "LOG",
    "MOD",
    "ABS",
    "CEIL", "FLOOR",
    "SET", "PROGRAM",
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
    "JMP"   : JMP,
    "JMT"   : JMT,
    "JMF"   : JMF,
    "sin"   : SIN,
    "asin"  : ASIN,
    "cos"   : COS,
    "acos"  : ACOS,
    "tan"   : TAN,
    "atan"  : ATAN,
    "atan2" : ATAN2,
    "exp"   : EXP,
    "ln"    : LN,
    "sqrt"  : SQRT,
    "pow"   : SQRT,
    "log"   : LOG,
    "mod"   : MOD,
    "abs"   : ABS,
    "ceil"  : CEIL,
    "floor" : FLOOR,
}

var IsUnaryOp map[int]bool = map[int]bool {
    NOT     : true,
    SIN     : true,
    ASIN    : true,
    COS     : true,
    ACOS    : true,
    TAN     : true,
    ATAN    : true,
    EXP     : true,
    LN      : true,
    ABS     : true,
    CEIL    : true,
    FLOOR   : true,
}

func ValidateOp(op int) error {
    if op < 0 || op >= opEnumCount {
        return errors.New("Unexpected operator found")
    } else {
        return nil
    }
}

// enumeration to keep track of loop style

const (
    InfStyle = iota
    ForStyle
    WhileStyle
)
