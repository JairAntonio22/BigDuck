package main

import (
    "fmt"
    "io/ioutil"
)

func run(filename string) {
    content, err := ioutil.ReadFile(filename)
    checkError(err)
    fmt.Println(string(content))
}
