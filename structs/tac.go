/*
    tac.go contains the implemenation for a three-address code instruction
    for code transformations through compilation process
*/

package structs

import "fmt"

type Tac struct {
    Op      int
    Args    [3]string
    Address [3]int
}

func (t Tac) Print() {
    fmt.Printf(
        "%8s | %8s %8s %8s | %8x %8x %8x\n",
        OpToString[t.Op],
        t.Args[0],
        t.Args[1],
        t.Args[2],
        t.Address[0],
        t.Address[1],
        t.Address[2])
}
