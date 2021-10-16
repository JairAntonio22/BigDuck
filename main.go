package main

import (
    "errors"
    "fmt"
    "os"
    "./parser"
    "github.com/antlr/antlr4/runtime/Go/antlr"
)

func checkError(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }
}

func main() {
    if len(os.Args) != 2 {
        checkError(errors.New("Usage: duck <namefile>"))
    }

    input, err := antlr.NewFileStream(os.Args[1])
    checkError(err)

    lexer := parser.NewBigDuckLexer(input)
    stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    parser := parser.NewBigDuckParser(stream)
    listener := &BigDuckListener{valid: true, debug: true}
    tree := parser.Program()

    parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
    parser.BuildParseTrees = true
    antlr.ParseTreeWalkerDefault.Walk(listener, tree)

    if listener.debug {
        fmt.Printf("Program is ")

        if listener.valid {
            fmt.Println("valid")
        } else {
            fmt.Println("not valid")
        }
    }
}
