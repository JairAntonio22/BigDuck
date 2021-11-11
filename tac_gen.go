package main

import (
    "fmt"
    "strconv"
    "./structs"
)

func (l *BigDuckListener) TopOp() int {
    if l.opstack.Empty() {
        return structs.NOP
    } else { item := l.opstack.Top()
        top, _ := item.(int)
        return top
    }
}

func (l *BigDuckListener) PopOp() int {
    if l.opstack.Empty() {
        return structs.NOP
    } else {
        item, _ := l.opstack.Pop()
        top, _ := item.(int)
        return top
    }
}

func (l *BigDuckListener) PushOp(op int) {
    if l.opstack.Empty() {
        l.opstack.Push(op)
    } else {
        top := l.TopOp()

        if top == structs.LPAREN {
            if op == structs.RPAREN {
                l.opstack.Pop()
            } else {
                l.opstack.Push(op)
            }

        } else {
            l.opstack.Push(op)
        }
    }
}

func (l *BigDuckListener) GenerateOpTAC() {
    var args [3]string
    var types [3]int
    var i, argc int

    op := l.PopOp()

    if op == structs.NOT {
        i = 0
        argc = 0
    } else {
        i = 1
        argc = 1
    }

    for ; i >= 0; i-- {
        item, _ := l.argstack.Pop()
        args[i], _ = item.(string)
        item, _ = l.typestack.Pop()
        types[i], _ = item.(int)
    }

    if structs.Cube[op][types[0]][types[argc]] == structs.Error_t {
        l.valid = false;
        fmt.Printf("line %d:%d type error mismatch\n", l.curr_line, l.curr_col)
    } else if op != structs.ASG {
        types[2] = structs.Cube[op][types[0]][types[argc]]
        l.typestack.Push(types[2])
    }

    if op == structs.ASG {
        args[2] = args[0]
        args[0] = args[1]
        args[1] = ""
        types[2] = types[0]
        types[0] = types[1]
        types[1] = 0
    } else {
        args[2] = "t" + strconv.Itoa(l.tmpc)
        l.argstack.Push(args[2])
        l.tmpc++
    }

    var address [3]int

    for i = 0; i < 3; i++ {
        scope, _, exists := l.symtable.Lookup(args[i])

        if exists || (len(args[i]) > 0 && args[i][0] == 't') {
            address[i] = l.memmap.GetAddress(scope, args[i], types[i])
        } else if len(args[i]) > 0 {
            address[i] = l.memmap.GetAddress(structs.Global, args[i], types[i])
        }
    }

    l.ir_code = append(
        l.ir_code,
        structs.Tac{
            Op: op,
            Args: args,
            Address: address})
    l.pc++
}

func (l *BigDuckListener) GenerateJmpTAC(jmptype int) {
    if jmptype == structs.JMP {
        l.ir_code = append(l.ir_code, structs.Tac{Op: jmptype})
    } else {
        item, _ := l.argstack.Pop()
        cond, _ := item.(string)
        item, _ = l.typestack.Pop()
        condtype, _ := item.(int)

        if condtype != structs.Bool_t {
            l.valid = false
            fmt.Printf(
                "line %d:%d expected boolean expression\n",
                l.curr_line, l.curr_col)
        }

        address := l.memmap.GetAddress(structs.Local, cond, condtype)

        l.ir_code = append(
            l.ir_code,
            structs.Tac{
                Op: jmptype,
                Args: [3]string{cond, "", ""},
                Address: [3]int{address, 0, 0}})
    }

    l.pc++
}

func (l *BigDuckListener) FillJmpTAC(index, target int) {
    l.ir_code[index].Args[2] = strconv.Itoa(target)
    l.ir_code[index].Address[2] = target
}

func (l *BigDuckListener) GenerateParamTAC() {
    item, _ := l.argstack.Pop()
    param, _ := item.(string)
    item, _ = l.typestack.Pop()
    ptype, _ := item.(int)

    _, sym, _ := l.symtable.Lookup(l.curr_pcall)
    argtype := sym.TypeArgs[l.paramc]

    if structs.Cube[structs.ASG][ptype][argtype] == structs.Error_t {
        l.valid = false
        fmt.Printf(
            "line %d:%d parameter %d expected to be %s, given %s\n",
            l.curr_line,
            l.curr_col,
            l.paramc + 1,
            structs.TypeToString[argtype],
            structs.TypeToString[ptype])
    }

    address := l.memmap.GetAddress(structs.Local, param, ptype)

    l.ir_code = append(
        l.ir_code, structs.Tac{
            Op: structs.PARAM,
            Args: [3]string{"", "", param},
            Address: [3]int{0, 0, address}})
    l.pc++
    l.paramc++
}

func (l *BigDuckListener) GenerateReturnTAC() {
    if l.ret_type == structs.Void_t {
        l.ir_code = append(
            l.ir_code,
            structs.Tac{Op: structs.RETURN})
        l.pc++
    } else {
        item, _ := l.argstack.Pop()
        result, _ := item.(string)
        item, _ = l.typestack.Pop()
        rtype, _ := item.(int)

        if rtype != l.ret_type {
            l.valid = false
            fmt.Printf(
                "line %d:%d return type different from procedure sign\n",
                l.curr_line, l.curr_col)
        }

        address := l.memmap.GetAddress(structs.Local, result, rtype)

        l.ir_code = append(
            l.ir_code,
            structs.Tac{
                Op: structs.RETURN,
                Args: [3]string{"", "", result},
                Address: [3]int{0, 0, address}})
        l.pc++
    }
}

func (l *BigDuckListener) GenerateProcCallRetTAC(procname string) {
    l.ir_code = append(
        l.ir_code,
        structs.Tac{
            Op: structs.GOPROC,
            Args: [3]string{"", "", procname}})
    l.pc++

    _, sym, _ := l.symtable.Lookup(procname)

    tmp := "t" + strconv.Itoa(l.tmpc)
    l.argstack.Push(tmp)
    l.typestack.Push(sym.RetType)
    l.tmpc++

    l.argstack.Push("_" + procname)
    l.typestack.Push(sym.RetType)
    l.PushOp(structs.OpFromString["<-"])
    l.GenerateOpTAC()

    l.argstack.Push(tmp)
    l.typestack.Push(sym.RetType)
}
