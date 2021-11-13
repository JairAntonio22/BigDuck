package main

import (
    "fmt"
    "strings"
    "os"
)

func checkError(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }
}

func main() {
    switch len(os.Args) {
    case 2:
        if !strings.Contains(os.Args[1], ".duck") {
            fmt.Println("Source code files must have .duck extension")
        } else {
            compile(os.Args[1], true)
        }

    case 3:
        if !strings.Contains(os.Args[2], ".quack") {
            fmt.Println("Executable files must have .quack extension")
        } else if os.Args[1] != "run" {
            fmt.Printf("%s option not recognized", os.Args[1])
        } else {
            run(os.Args[2])
        }

    default:
        fmt.Println("Compile: duck filename.duck")
        fmt.Println("Run: duck run filename.quack")
    }
}
