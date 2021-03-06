package main

import (
    "os"
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

func (l *BigDuckListener) GenerateOpTAC(pointer int) {
    var args [3]string
    var types [3]int
    var i, argc int

    op := l.PopOp()

    is_unary, _ := structs.IsUnaryOp[op]

    if is_unary {
        argc = 0
    } else {
        argc = 1
    }

    for i = argc; i >= 0; i-- {
        item, _ := l.argstack.Pop()
        args[i], _ = item.(string)
        item, _ = l.typestack.Pop()
        types[i], _ = item.(int)
    }

    if structs.Cube[op][types[0]][types[argc]] == structs.Error_t {
        l.valid = false;
        fmt.Printf("line %d:%d type error mismatch\n", l.curr_line, l.curr_col)
        fmt.Println(args[0], args[1])
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
	    if i + 1 == pointer {
		l.memmap.RegisterPointer(scope, args[i])
	    }

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

        var address int

        if cond == "#t" || cond == "#f" {
            address = l.memmap.GetAddress(structs.Global, cond, condtype)
        } else {
            address = l.memmap.GetAddress(structs.Local, cond, condtype)
        }

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

    _, procsym, _ := l.symtable.Lookup(l.curr_pcall)

    argtype := procsym.TypeArgs[l.paramc]

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

    scope, _, exists := l.symtable.Lookup(param)
    var address int

    if exists {
        address = l.memmap.GetAddress(scope, param, ptype)
    } else {
        address = l.memmap.GetAddress(structs.Global, param, ptype)
    }

    l.ir_code = append(
        l.ir_code, structs.Tac{
            Op: structs.PARAM,
            Args: [3]string{param, "", fmt.Sprintf("%x", procsym.Paddress[l.paramc])},
            Address: [3]int{address, 0, procsym.Paddress[l.paramc]}})
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

        scope := structs.Local
        _, _, exists := l.symtable.Lookup(result)

        if !exists || (len(result) >= 1 && result[0] != 't') {
            scope = structs.Global
        }

        address1 := l.memmap.GetAddress(scope, result, rtype)
        address2 := l.memmap.GetAddress(structs.Global, "_" + l.curr_proc, rtype)

        l.ir_code = append(
            l.ir_code,
            structs.Tac{
                Op: structs.RETURN,
                Args: [3]string{result, "", "_" + l.curr_proc},
                Address: [3]int{address1, 0, address2}})
        l.pc++
    }
}

func (l *BigDuckListener) GenerateProcCallRetTAC(procname string) {
    _, sym, _ := l.symtable.Lookup(procname)

    l.ir_code = append(
        l.ir_code,
        structs.Tac{
            Op: structs.GOPROC,
            Args: [3]string{"", "", procname},
            Address: [3]int{0, 0, sym.Startpoint}})
    l.pc++

    tmp := "t" + strconv.Itoa(l.tmpc)
    l.argstack.Push(tmp)
    l.typestack.Push(sym.RetType)
    l.tmpc++

    l.argstack.Push("_" + procname)
    l.typestack.Push(sym.RetType)
    l.PushOp(structs.OpFromString["<-"])
    l.GenerateOpTAC(0)

    l.argstack.Push(tmp)
    l.typestack.Push(sym.RetType)
}

func (l *BigDuckListener) GeneratePrintTAC() {
    item, _ := l.argstack.Pop()
    param, _ := item.(string)
    item, _ = l.typestack.Pop()
    ptype, _ := item.(int)

    scope, sym, exists := l.symtable.Lookup(param)
    var address int

    if exists {
        address = l.memmap.GetAddress(scope, param, ptype)
    } else {
        address = l.memmap.GetAddress(structs.Global, param, ptype)
    }

    if exists {
        if len(sym.Dim) == 1 {
            for i := 0; i < sym.Dim[0]; i++ {
                l.ir_code = append(
                    l.ir_code, structs.Tac{
                        Op: structs.PRINT,
                        Args: [3]string{"", "", param},
                        Address: [3]int{0, 0, address + i}})
                l.pc++
            }

        } else if len(sym.Dim) == 2 {
            for i := 0; i < sym.Dim[0]; i++ {
                for j := 0; j < sym.Dim[1]; j++ {
                    l.ir_code = append(
                        l.ir_code, structs.Tac{
                            Op: structs.PRINT,
                            Args: [3]string{"", "", param},
                            Address: [3]int{0, 0, address + i * sym.Dim[0] + j}})
                    l.pc++
                }

                l.ir_code[l.pc - 1].Op = structs.PRINTLN
            }

        } else if len(sym.Dim) > 2 {
            msg := "\"No available print for higher dimensional tensors\""
            address = l.memmap.GetAddress(structs.Global, msg, structs.String_t)
            l.ir_code = append(
                l.ir_code, structs.Tac{
                    Op: structs.PRINT,
                    Args: [3]string{"", "", msg},
                    Address: [3]int{0, 0, address}})
            l.pc++
        } else {
            l.ir_code = append(
                l.ir_code, structs.Tac{
                    Op: structs.PRINT,
                    Args: [3]string{"", "", param},
                    Address: [3]int{0, 0, address}})
            l.pc++
        }

    } else {
	l.ir_code = append(
	    l.ir_code, structs.Tac{
		Op: structs.PRINT,
		Args: [3]string{"", "", param},
		Address: [3]int{0, 0, address}})
	l.pc++
    }
}

func (l *BigDuckListener) GenerateObjFile() {
    content := ""

    for _, code := range l.data_seg {
        content += code.GetArgs()
    }

    content += fmt.Sprintf("%s", structs.Tac{Op: structs.PROGRAM}.GetAddress())

    for _, code := range l.ir_code {
        content += code.GetAddress()
    }

    file, err := os.Create(l.filename[:len(l.filename) - 5] + ".quack")
    checkError(err)
    defer file.Close()
    file.WriteString(content)
}
