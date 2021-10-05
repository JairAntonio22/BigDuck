package main

import (
    "errors"
    "fmt"
    "os"
    "./parser"
    "github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
    if len(os.Args) != 2 {
        checkError(errors.New("Usage: duck <namefile>"))
    }

    input, err := antlr.NewFileStream(os.Args[1])
    checkError(err)

    lexer := parser.NewBigDuckLexer(input)
    stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    parser := parser.NewBigDuckParser(stream)
    listener := &BigDuckListener{valid: true}
    tree := parser.Program()

    parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
    parser.BuildParseTrees = true

    antlr.ParseTreeWalkerDefault.Walk(listener, tree)

    fmt.Println("")

    if listener.valid {
        fmt.Println("accepted")

    } else {
        fmt.Println("rejected")
    }
}
