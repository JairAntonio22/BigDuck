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
        6 0110 pointer

    memory map
        1 addressing mode bit, 1 is global bit, 3 type bits, 7 address nibbles

        0010 ... 0000 0000 -> local     int     at address 0
        1011 ... 0000 1010 -> global    float   at address 10
        0100 ... 0001 0110 -> local     bool    at address 22
        1101 ... 0000 1011 -> global    string  at address 11
        1110 ... 0000 1111 -> global    pointer at address 15
*/

type memory struct {
    Strings     map[int]string
    Pointers    map[int]int
    MemI        [2][]int
    MemF        [2][]float64
    MemB        [2][]bool
    memstack    memoryStack
    framesize   map[int]int
    Sp          map[int]int
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
    m.Strings = make(map[int]string, ic)
    m.MemI[Global] = make([]int, ic)
    m.MemF[Global] = make([]float64, fc)
    m.MemB[Global] = make([]bool, bc)
    m.framesize = make(map[int]int)
    m.Sp = make(map[int]int)
}

func (m *memory) InitLocal(ic, fc, bc int) {
    m.MemI[Local] = append(m.MemI[Local], make([]int, ic)...)
    m.MemF[Local] = append(m.MemF[Local], make([]float64, fc)...)
    m.MemB[Local] = append(m.MemB[Local], make([]bool, bc)...)

    m.memstack.Push(ic, fc, bc)
}

func (m *memory) PushContext() {
    m.Sp[Int_t]     += m.framesize[Int_t]
    m.Sp[Float_t]   += m.framesize[Float_t]
    m.Sp[Bool_t]    += m.framesize[Bool_t]

    ic, fc, bc := m.memstack.Top()
    m.framesize[Int_t]      = ic
    m.framesize[Float_t]    = fc
    m.framesize[Bool_t]     = bc
}

func (m *memory) PopContext() {
    ic, fc, bc := m.memstack.Pop()

    m.MemI[Local] = append(
        []int(nil), m.MemI[Local][:len(m.MemI[Local]) - ic]...)
    m.MemF[Local] = append(
        []float64(nil), m.MemF[Local][:len(m.MemF[Local]) - fc]...)
    m.MemB[Local] = append(
        []bool(nil), m.MemB[Local][:len(m.MemB[Local]) - bc]...)

    ic, fc, bc = m.memstack.Top()

    m.framesize[Int_t]      = ic
    m.framesize[Float_t]    = fc
    m.framesize[Bool_t]     = bc

    m.Sp[Int_t]     -= m.framesize[Int_t]
    m.Sp[Float_t]   -= m.framesize[Float_t]
    m.Sp[Bool_t]    -= m.framesize[Bool_t]
}
