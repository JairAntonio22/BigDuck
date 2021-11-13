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
        "%8s | %10s%10s%10s | %10x%10x%10x\n",
        OpToString[t.Op],
        t.Args[0],
        t.Args[1],
        t.Args[2],
        t.Address[0],
        t.Address[1],
        t.Address[2])
}

func (t Tac) GetArgs() string {
    result := fmt.Sprintf("%x ", t.Op)

    for i := 0; i < 3; i++ {
        if t.Args[i] == "" {
            result += fmt.Sprintf("0 ")
        } else {
            result += fmt.Sprintf("%s ", t.Args[i])
        }
    }

    return result
}

func (t Tac) GetAddress() string {
    return fmt.Sprintf(
        "%x %x %x %x ", t.Op, t.Address[0], t.Address[1], t.Address[2])
}

func (t *Tac) Clear() {
    t.Op = 0

    for i := 0; i < 3; i++ {
        t.Args[i] = ""
        t.Address[i] = 0
    }
}
