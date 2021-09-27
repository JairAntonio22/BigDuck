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

type BigDuckListener struct {
    *parser.BaseBigDuckListener
}

var validInput = true

func (this *BigDuckListener) VisitErrorNode(node antlr.ErrorNode) {
    validInput = false
}

func main() {
    if len(os.Args) != 2 {
        checkError(errors.New("Usage: duck <namefile>"))
    }

    input, err := antlr.NewFileStream(os.Args[1])
    checkError(err)

    lexer := parser.NewBigDuckLexer(input)
    stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

    p := parser.NewBigDuckParser(stream)
    p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
    p.BuildParseTrees = true

    tree := p.Programa()
    antlr.ParseTreeWalkerDefault.Walk(&BigDuckListener{}, tree)

    if validInput {
        fmt.Println("accepted")

    } else {
        fmt.Println("rejected")
    }
}
