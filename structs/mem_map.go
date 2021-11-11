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
*/

/*
    memory map

        1 is global bit, 3 type bits, 7 address nibbles

        0010 ... 0000 0000 -> local int at address 0
        1011 ... 0000 1010 -> global float at address 10
        0100 ... 0001 0110 -> local bool at address 22
*/

type MemMap struct {
    cache       map[string]int
    typecount   [2]map[int]int
}

func NewMemMap() MemMap {
    return MemMap{
        cache: make(map[string]int),
        typecount: [2]map[int]int{
            make(map[int]int),
            make(map[int]int)}}
}

func (m *MemMap) GetAddress(scope int, var_name string, vtype int) int {
    address, exists := m.cache[var_name]

    if !exists {
        address |= (scope << 31) | (vtype << 28) | m.typecount[scope][vtype]
        m.typecount[scope][vtype]++
        m.cache[var_name] = address
    }

    return address
}

func (m *MemMap) ClearLocalScope() {
    for key, value := range m.cache {
        if value & (1 << 31) == 0 {
            delete(m.cache, key)
        }
    }

    m.typecount[Local][Int_t] = 0
    m.typecount[Local][Float_t] = 0
    m.typecount[Local][Bool_t] = 0
}
