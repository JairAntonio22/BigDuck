/*
    tac.go contains the implemenation for a three-address code instruction
    for code transformations through compilation process
*/

package structs

import "fmt"

type Tac struct {
    Op      int
    Arg1    string
    Arg2    string
    Target  string
}

func (t Tac) Print() {
    fmt.Printf(
        "%s\t%s\t%s\t%s\n",
        OperatorToString(t.Op),
        t.Arg1,
        t.Arg2,
        t.Target)
}
