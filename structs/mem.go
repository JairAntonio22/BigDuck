package structs

/*
    String Constants
    00 00 00 00

    Local Scope
    02 00 00 00 int
    03 00 00 00 float
    04 00 00 00 bool

    Global Scope
    12 00 00 00 int
    13 00 00 00 float
    14 00 00 00 bool
*/

type Mem struct {
    strings     []string
    size        map[string][]int
    init        map[string][]int
}

func (m *Mem) Init() {
}

func (m Mem) At(vir_add string) (int, interface{}) {
    //pool := vir_add[0:2]


    return Error_t, nil
}
