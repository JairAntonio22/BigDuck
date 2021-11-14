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

const (
    Prev = iota
    Curr
    Size
)

type memory struct {
    MemI        [2][]int
    MemF        [2][]float64
    MemB        [2][]bool
    Sp          [3]map[int]int
    pcallc      int
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

    m.Sp[Prev] = make(map[int]int)
    m.Sp[Curr] = make(map[int]int)
    m.Sp[Size] = make(map[int]int)

    m.Sp[Prev][Int_t] = 0
    m.Sp[Prev][Float_t] = 0
    m.Sp[Prev][Bool_t] = 0

    m.Sp[Curr][Int_t] = 0
    m.Sp[Curr][Float_t] = 0
    m.Sp[Curr][Bool_t] = 0

    m.Sp[Size][Int_t] = 0
    m.Sp[Size][Float_t] = 0
    m.Sp[Size][Bool_t] = 0
}

func (m *memory) InitLocal(ic, fc, bc int) {
    m.MemI[Local] = append(m.MemI[Local], make([]int, ic)...)
    m.MemF[Local] = append(m.MemF[Local], make([]float64, fc)...)
    m.MemB[Local] = append(m.MemB[Local], make([]bool, bc)...)

    m.Sp[Size][Int_t]   = ic
    m.Sp[Size][Float_t] = fc
    m.Sp[Size][Bool_t]  = bc
}

func (m *memory) ChangeContext() {
    m.Sp[Prev][Int_t]   = m.Sp[Curr][Int_t]
    m.Sp[Prev][Float_t] = m.Sp[Curr][Float_t]
    m.Sp[Prev][Bool_t]  = m.Sp[Curr][Bool_t]

    m.Sp[Curr][Int_t]   += m.Sp[Size][Int_t]
    m.Sp[Curr][Float_t] += m.Sp[Size][Float_t]
    m.Sp[Curr][Bool_t]  += m.Sp[Size][Bool_t]
}
