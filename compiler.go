package main

import (
    "fmt"
    "os"
    "./parser"
    "github.com/antlr/antlr4/runtime/Go/antlr"
)

func compile(filename string, debug bool) {
    input, err := antlr.NewFileStream(filename)
    checkError(err)

    lexer := parser.NewBigDuckLexer(input)
    stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    parser := parser.NewBigDuckParser(stream)
    listener := &BigDuckListener{
        valid: true, debug: debug, filename: filename}
    tree := parser.Program()

    parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
    parser.BuildParseTrees = true
    antlr.ParseTreeWalkerDefault.Walk(listener, tree)

    if listener.debug {
        if listener.valid {
            fmt.Println("Program is valid")
        } else {
            fmt.Println("Program is not valid")
	    os.Exit(1)
        }
    }
}
