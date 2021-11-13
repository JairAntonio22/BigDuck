/*
    mem_map.go contains the implementation of a memory mapper for variables
*/

package structs

/*
    scope enum:
        0 local
        1 global

    type enum:
        2 0010 int
        3 0011 float
        4 0100 bool

    memory map
        1 is global bit, 3 type bits, 7 address nibbles

        0010 ... 0000 0000 -> local int at address 0
        1011 ... 0000 1010 -> global float at address 10
        0100 ... 0001 0110 -> local bool at address 22
*/

import (
    "fmt"
    "strconv"
)

type MemMap struct {
    memcache    map[string]int
    varcache    map[int]string
    Typecount   [2]map[int]int
}

func NewMemMap() MemMap {
    return MemMap{
        memcache: make(map[string]int),
        varcache: make(map[int]string),
        Typecount: [2]map[int]int{
            make(map[int]int),
            make(map[int]int)}}
}

func (m *MemMap) GetAddress(scope int, var_name string, vtype int) int {
    address, exists := m.memcache[var_name]

    if !exists {
        address |= (scope << 31) | (vtype << 28) | m.Typecount[scope][vtype]
        m.Typecount[scope][vtype]++
        m.memcache[var_name] = address
        m.varcache[address] = var_name
    }

    return address
}

func (m *MemMap) GetDataSegment() []Tac {
    var ir_code []Tac
    ir_code = append(ir_code, Tac{ Op: ERA})

    ir_code[0].Args[0] = strconv.Itoa(m.Typecount[Global][Int_t])
    ir_code[0].Args[1] = strconv.Itoa(m.Typecount[Global][Float_t])
    ir_code[0].Args[2] = strconv.Itoa(m.Typecount[Global][Bool_t])
    ir_code[0].Address[0] = m.Typecount[Global][Int_t]
    ir_code[0].Address[1] = m.Typecount[Global][Float_t]
    ir_code[0].Address[2] = m.Typecount[Global][Bool_t]

    for key, value := range m.memcache {
        if value & (1 << 31) != 0 {
            ir_code = append(
                ir_code,
                Tac{
                    Op: SET,
                    Args: [3]string{key, "", fmt.Sprintf("%8x", value)},
                    Address: [3]int{}})
        }
    }

    return ir_code
}

func (m *MemMap) ClearLocalScope() {
    for key, value := range m.memcache {
        if value & (1 << 31) == 0 {
            delete(m.memcache, key)
        }
    }

    m.Typecount[Local][Int_t] = 0
    m.Typecount[Local][Float_t] = 0
    m.Typecount[Local][Bool_t] = 0
}
