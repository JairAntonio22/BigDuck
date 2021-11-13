package structs

/*
    scope enum:
        0 local
        1 global

    type enum:
        2 0010 int
        3 0011 float
        4 0100 bool
        5 0101 string

    memory map
        1 is global bit, 3 type bits, 7 address nibbles

        0010 ... 0000 0000 -> local     int     at address 0
        1011 ... 0000 1010 -> global    float   at address 10
        0100 ... 0001 0110 -> local     bool    at address 22
        1101 ... 0000 1011 -> global    string  at address 11
*/

type memory struct {
    MemI        [2][]int
    MemF        [2][]float64
    MemB        [2][]bool
}

func GetScope(address int) int {
    return address & (0x1 << 31) >> 31
}

func GetType(address int) int {
    return address & (0x7 << 28) >> 28
}

func GetAddress(address int) int {
    return address & 0x0fffffff
}

func (m *memory) InitGlobal(ic, fc, bc int) {
    m.MemI[Global] = make([]int, ic)
    m.MemF[Global] = make([]float64, fc)
    m.MemB[Global] = make([]bool, bc)
}
