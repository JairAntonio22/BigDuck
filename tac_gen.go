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
    var args [2]string
    var types [2]int
    var target string
    var i, j int

    op := l.PopOp()

    if op == structs.NOT {
        i = 0
        j = 0
    } else {
        i = 1
        j = 1
    }

    for ; i >= 0; i-- {
        item, _ := l.argstack.Pop()
        args[i], _ = item.(string)
        item, _ = l.typestack.Pop()
        types[i], _ = item.(int)
    }

    if structs.Cube[op][types[0]][types[j]] == structs.Error_t {
        l.valid = false;
        fmt.Printf("line %d:%d type error mismatch\n", l.curr_line, l.curr_col)
    } else if op != structs.ASG {
        l.typestack.Push(structs.Cube[op][types[0]][types[j]])
    }

    if op == structs.ASG {
        target = args[0]
        args[0] = args[1]
        args[1] = ""
    } else {
        target = "t" + strconv.Itoa(l.tmpc)
        l.argstack.Push(target)
        l.tmpc++
    }

    l.ir_code = append(
        l.ir_code,
        structs.Tac{Op: op, Arg1: args[0], Arg2: args[1], Target: target})
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

        l.ir_code = append(l.ir_code, structs.Tac{Op: jmptype, Arg1: cond})
    }

    l.pc++
}

func (l *BigDuckListener) FillJmpTAC(index, target int) {
    l.ir_code[index].Target = strconv.Itoa(target)
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

    target := "p" + strconv.Itoa(l.paramc)
    l.paramc++

    l.ir_code = append(
        l.ir_code, structs.Tac{Op: structs.PARAM, Arg1: param, Target: target})
    l.pc++
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

        l.ir_code = append(
            l.ir_code,
            structs.Tac{Op: structs.RETURN, Target: result})
        l.pc++
    }
}

func (l *BigDuckListener) GenerateProcCallRetTAC(procname string) {
    l.ir_code = append(
        l.ir_code,
        structs.Tac{Op: structs.GOPROC, Target: procname})
    l.pc++

    _, sym, _ := l.symtable.Lookup(procname)

    tmp := "t" + strconv.Itoa(l.tmpc)
    l.argstack.Push(tmp)
    l.typestack.Push(sym.RetType)
    l.tmpc++

    l.argstack.Push("__" + procname)
    l.typestack.Push(sym.RetType)
    l.PushOp(structs.OpFromString["<-"])
    l.GenerateOpTAC()

    l.argstack.Push(tmp)
    l.typestack.Push(sym.RetType)
}
