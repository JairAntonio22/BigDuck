
/*
    utils.go contains commonly used functions between the project that can not
    be grouped into a single file with similar functionality
*/

package main

import (
    "fmt"
    "os"
)

func checkError(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }
}

