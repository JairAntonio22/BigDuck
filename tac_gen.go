package main

import (
    "strconv"
    "./structs"
)

func (l *BigDuckListener) TopOp() int {
    if l.opstack.Empty() {
        return structs.NOP
    } else {
        item := l.opstack.Top()
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
    var target string
    var i int

    op := l.PopOp()

    if op == structs.NOT {
        i = 0
    } else {
        i = 1
    }

    for ; i >= 0; i-- {
        item, _ := l.argstack.Pop()
        args[i], _ = item.(string)
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

    target := "p" + strconv.Itoa(l.paramc)
    l.paramc++

    l.ir_code = append(
        l.ir_code,
        structs.Tac{Op: structs.PARAM, Arg1: param, Target: target})
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

        l.ir_code = append(
            l.ir_code,
            structs.Tac{Op: structs.RETURN, Target: result})
        l.pc++
    }
}

func (l *BigDuckListener) GenerateProcCallRetTAC(procname string) {
    l.ir_code = append(l.ir_code, structs.Tac{
        Op: structs.GOPROC, Target: procname})
    l.pc++

    tmp := "t" + strconv.Itoa(l.tmpc)
    l.argstack.Push(tmp)
    l.tmpc++

    l.argstack.Push("__" + procname)
    l.PushOp(structs.OpFromString["<-"])
    l.GenerateOpTAC()

    l.argstack.Push(tmp)
}
