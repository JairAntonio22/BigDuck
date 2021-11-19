package structs

import (
    "fmt"
    "os"
    "strconv"
)

type VirtualMachine struct {
    Debug           bool
    Program         []Tac       // Program code
    memory          memory      // Memory manager
    basepc          int         // Pointer to program start
    pc              int         // Program counter
    pcstack         Stack       // Program counter stack
    pvaluequeue     Queue       // Parameter value queue
    paddressqueue   Queue       // Parameter address queue
    init_ptr	    bool
}

func (vm *VirtualMachine) InitMemory() {
    var curr_code Tac

    for ; curr_code.Op != PROGRAM; vm.pc++ {
        curr_code =  vm.Program[vm.pc]

        switch curr_code.Op {
        case ERA:
            vm.memory.InitGlobal(
                curr_code.Address[0],
                curr_code.Address[1],
                curr_code.Address[2])

        case SET:
            va := curr_code.Address[2]
            s := GetScope(va)
            a := GetAddress(va)

            if s != Global {
                fmt.Printf("line %d: Local address used in data segment\n", vm.pc)
                os.Exit(1)
            }

            switch GetType(va) {
            case Int_t:
                value, _ := strconv.ParseInt(curr_code.Args[0], 10, 64)
                vm.memory.MemI[s][a] = int(value)

            case Float_t:
                value, _ := strconv.ParseFloat(curr_code.Args[0], 10)
                vm.memory.MemF[s][a] = value

            case Bool_t:
                if curr_code.Args[0] == "#t" {
                    vm.memory.MemB[s][a] = true
                } else {
                    vm.memory.MemB[s][a] = false
                }

            case String_t:
                value := curr_code.Args[0][1:len(curr_code.Args[0]) - 1]
                vm.memory.Strings[a] = value

            default:
                fmt.Printf("line %d: Invalid address used in data segment\n", vm.pc)
                os.Exit(1)
            }

        case PROGRAM:
            vm.basepc = vm.pc
            break

        default:
            fmt.Printf(
                "line %d: Unexpected operator %s at data segment\n",
		vm.pc,
                TypeToString[curr_code.Op])
            os.Exit(1)
        }
    }
}

func (vm *VirtualMachine) Execute() {
    var curr_code Tac
    run_program := true

    for ; run_program ; vm.pc++ {
        curr_code =  vm.Program[vm.pc]

        if vm.Debug {
            fmt.Printf("%3d ", vm.pc - vm.basepc - 1)
            curr_code.Print()
        }

        va1 := curr_code.Address[0]
        va2 := curr_code.Address[1]
        va3 := curr_code.Address[2]

        s1 := GetScope(va1)
        s2 := GetScope(va2)
        s3 := GetScope(va3)

        t1 := GetType(va1)
        t2 := GetType(va2)
        t3 := GetType(va3)

        var a1, a2, a3 int

	if IsPointer(va1) {
	    ref := vm.memory.MemI[s1][vm.memory.Sp[t1] + GetAddress(va1)]

	    s1 = GetScope(ref)
	    t1 = GetType(ref)
	    a1 = GetAddress(ref)

	} else {
	    if s1 == Local {
		a1 = vm.memory.Sp[t1] + GetAddress(va1)
	    } else {
		a1 = GetAddress(va1)
	    }
	}

	if IsPointer(va2) {
	    ref := vm.memory.MemI[s2][vm.memory.Sp[t2] + GetAddress(va2)]

	    s2 = GetScope(ref)
	    t2 = GetType(ref)
	    a2 = GetAddress(ref)

	} else {
	    if s2 == Local {
		a2 = vm.memory.Sp[t2] + GetAddress(va2)
	    } else {
		a2 = GetAddress(va2)
	    }
	}

	if IsPointer(va3) && !vm.init_ptr {
	    ref := vm.memory.MemI[s3][vm.memory.Sp[t3] + GetAddress(va3)]

	    s3 = GetScope(ref)
	    t3 = GetType(ref)
	    a3 = GetAddress(ref)

	} else {
	    if s3 == Local {
		a3 = vm.memory.Sp[t3] + GetAddress(va3)
	    } else {
		a3 = GetAddress(va3)
	    }

	    if IsPointer(va3) {
		vm.init_ptr = false
	    }
	}

        switch curr_code.Op {
        case ASG:
            if t1 == Int_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = vm.memory.MemI[s1][a1]

            } else if t1 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = vm.memory.MemF[s1][a1]

            } else if t1 == Bool_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = vm.memory.MemB[s1][a1]

            } else if t1 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = float64(vm.memory.MemI[s1][a1])

            } else if t1 == Float_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = int(vm.memory.MemF[s1][a1])

            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case OR:
            if t1 == Bool_t && t2 == Bool_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemB[s1][a1] || vm.memory.MemB[s2][a2])
            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case AND:
            if t1 == Bool_t && t2 == Bool_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemB[s1][a1] && vm.memory.MemB[s2][a2])
            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case NOT:
            if t1 == Bool_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = !vm.memory.MemB[s1][a1]
            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case EQ:
            if t1 == Int_t && t2 == Int_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemI[s1][a1] == vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    float64(vm.memory.MemI[s1][a1]) == vm.memory.MemF[s2][a2])

            } else if t1 == Float_t && t2 == Int_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemF[s1][a1] == float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Float_t && t2 == Float_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemF[s1][a1] == vm.memory.MemF[s2][a2])

            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case NEQ:
            if t1 == Int_t && t2 == Int_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemI[s1][a1] != vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    float64(vm.memory.MemI[s1][a1]) != vm.memory.MemF[s2][a2])

            } else if t1 == Float_t && t2 == Int_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemF[s1][a1] != float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Float_t && t2 == Float_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemF[s1][a1] != vm.memory.MemF[s2][a2])

            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case LES:
            if t1 == Int_t && t2 == Int_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemI[s1][a1] < vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    float64(vm.memory.MemI[s1][a1]) < vm.memory.MemF[s2][a2])

            } else if t1 == Float_t && t2 == Int_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemF[s1][a1] < float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Float_t && t2 == Float_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemF[s1][a1] < vm.memory.MemF[s2][a2])

            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case GRE:
            if t1 == Int_t && t2 == Int_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemI[s1][a1] > vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    float64(vm.memory.MemI[s1][a1]) > vm.memory.MemF[s2][a2])

            } else if t1 == Float_t && t2 == Int_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemF[s1][a1] > float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Float_t && t2 == Float_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemF[s1][a1] > vm.memory.MemF[s2][a2])

            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case LEQ:
            if t1 == Int_t && t2 == Int_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemI[s1][a1] <= vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    float64(vm.memory.MemI[s1][a1]) <= vm.memory.MemF[s2][a2])

            } else if t1 == Float_t && t2 == Int_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemF[s1][a1] <= float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Float_t && t2 == Float_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemF[s1][a1] <= vm.memory.MemF[s2][a2])

            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case GEQ:
            if t1 == Int_t && t2 == Int_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemI[s1][a1] >= vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    float64(vm.memory.MemI[s1][a1]) >= vm.memory.MemF[s2][a2])

            } else if t1 == Float_t && t2 == Int_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemF[s1][a1] >= float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Float_t && t2 == Float_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemF[s1][a1] >= vm.memory.MemF[s2][a2])

            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case SUB:
            if t1 == Int_t && t2 == Int_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = (
                    vm.memory.MemI[s1][a1] - vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = (
                    vm.memory.MemI[s1][a1] - int(vm.memory.MemF[s2][a2]))

            } else if t1 == Float_t && t2 == Int_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = (
                    int(vm.memory.MemF[s1][a1]) - vm.memory.MemI[s2][a2])

            } else if t1 == Float_t && t2 == Float_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = int(
                    vm.memory.MemF[s1][a1] - vm.memory.MemF[s2][a2])

            } else if t1 == Int_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = float64(
                    vm.memory.MemI[s1][a1] - vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = (
                    float64(vm.memory.MemI[s1][a1]) - vm.memory.MemF[s2][a2])

            } else if t1 == Float_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = (
                    vm.memory.MemF[s1][a1] - float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Float_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = (
                    vm.memory.MemF[s1][a1] - vm.memory.MemF[s2][a2])

            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case ADD:
            if t1 == Int_t && t2 == Int_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = (
                    vm.memory.MemI[s1][a1] + vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = (
                    vm.memory.MemI[s1][a1] + int(vm.memory.MemF[s2][a2]))

            } else if t1 == Float_t && t2 == Int_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = (
                    int(vm.memory.MemF[s1][a1]) + vm.memory.MemI[s2][a2])

            } else if t1 == Float_t && t2 == Float_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = int(
                    vm.memory.MemF[s1][a1] + vm.memory.MemF[s2][a2])

            } else if t1 == Int_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = float64(
                    vm.memory.MemI[s1][a1] + vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = (
                    float64(vm.memory.MemI[s1][a1]) + vm.memory.MemF[s2][a2])

            } else if t1 == Float_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = (
                    vm.memory.MemF[s1][a1] + float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Float_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = (
                    vm.memory.MemF[s1][a1] + vm.memory.MemF[s2][a2])

            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case DIV:
            if t1 == Int_t && t2 == Int_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = (
                    vm.memory.MemI[s1][a1] / vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = (
                    vm.memory.MemI[s1][a1] / int(vm.memory.MemF[s2][a2]))

            } else if t1 == Float_t && t2 == Int_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = (
                    int(vm.memory.MemF[s1][a1]) / vm.memory.MemI[s2][a2])

            } else if t1 == Float_t && t2 == Float_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = int(
                    vm.memory.MemF[s1][a1] / vm.memory.MemF[s2][a2])

            } else if t1 == Int_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = float64(
                    vm.memory.MemI[s1][a1] / vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = (
                    float64(vm.memory.MemI[s1][a1]) / vm.memory.MemF[s2][a2])

            } else if t1 == Float_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = (
                    vm.memory.MemF[s1][a1] / float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Float_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = (
                    vm.memory.MemF[s1][a1] / vm.memory.MemF[s2][a2])

            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case MUL:
            if t1 == Int_t && t2 == Int_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = (
                    vm.memory.MemI[s1][a1] * vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = (
                    vm.memory.MemI[s1][a1] * int(vm.memory.MemF[s2][a2]))

            } else if t1 == Float_t && t2 == Int_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = (
                    int(vm.memory.MemF[s1][a1]) * vm.memory.MemI[s2][a2])

            } else if t1 == Float_t && t2 == Float_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = int(
                    vm.memory.MemF[s1][a1] * vm.memory.MemF[s2][a2])

            } else if t1 == Int_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = float64(
                    vm.memory.MemI[s1][a1] * vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = (
                    float64(vm.memory.MemI[s1][a1]) * vm.memory.MemF[s2][a2])

            } else if t1 == Float_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = (
                    vm.memory.MemF[s1][a1] * float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Float_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = (
                    vm.memory.MemF[s1][a1] * vm.memory.MemF[s2][a2])

            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case LPAREN:
            fmt.Printf(
		"line %d: Unexpected operator %s\n",
		vm.pc,
		TypeToString[curr_code.Op])
            os.Exit(1)

        case RPAREN:
            fmt.Printf(
		"line %d: Unexpected operator %s\n",
		vm.pc,
		TypeToString[curr_code.Op])
            os.Exit(1)

        case JMP:
            vm.pc = vm.basepc + curr_code.Address[2]

        case JMT:
            if t1 == Bool_t && vm.memory.MemB[s1][a1] == true {
                vm.pc = vm.basepc + curr_code.Address[2]
            }

        case JMF:
            if t1 == Bool_t && vm.memory.MemB[s1][a1] == false {
                vm.pc = vm.basepc + curr_code.Address[2]
            }

        case GOPROC:
            vm.pcstack.Push(vm.pc)
            vm.pc = vm.basepc + curr_code.Address[2]
            vm.memory.PushContext()

            for !vm.paddressqueue.Empty() {
                item, _ := vm.paddressqueue.Pop()

                va, _ := item.(int)
                t := GetType(va)
                a := vm.memory.Sp[t] + GetAddress(va)

                item, _ = vm.pvaluequeue.Pop()

                switch t {
                case Int_t:
                    value, _ := item.(int)
                    vm.memory.MemI[Local][a] = value

                case Float_t:
                    value, _ := item.(float64)
                    vm.memory.MemF[Local][a] = value

                case Bool_t:
                    value, _ := item.(bool)
                    vm.memory.MemB[Local][a] = value

                default:
                    fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                    os.Exit(1)
                }
            }

        case ERA:
            vm.memory.InitLocal(
                curr_code.Address[0],
                curr_code.Address[1],
                curr_code.Address[2])

        case PARAM:
            vm.paddressqueue.Push(va3)

            if t1 == Int_t && t3 == Int_t {
                vm.pvaluequeue.Push(vm.memory.MemI[s1][a1])

            } else if t1 == Float_t && t3 == Float_t {
                vm.pvaluequeue.Push(vm.memory.MemF[s1][a1])

            } else if t1 == Bool_t && t3 == Bool_t {
                vm.pvaluequeue.Push(vm.memory.MemB[s1][a1])

            } else if t1 == Int_t && t3 == Float_t {
                vm.pvaluequeue.Push(float64(vm.memory.MemI[s1][a1]))

            } else if t1 == Float_t && t3 == Int_t {
                vm.pvaluequeue.Push(int(vm.memory.MemF[s1][a1]))

            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case RETURN:
            if va1 != 0 && va3 != 0 {
                if t1 == Int_t && t3 == Int_t {
                    vm.memory.MemI[s3][a3] = vm.memory.MemI[s1][a1]

                } else if t1 == Float_t && t3 == Float_t {
                    vm.memory.MemF[s3][a3] = vm.memory.MemF[s1][a1]

                } else if t1 == Bool_t && t3 == Bool_t {
                    vm.memory.MemB[s3][a3] = vm.memory.MemB[s1][a1]

                } else if t1 == Int_t && t3 == Float_t {
                    vm.memory.MemF[s3][a3] = float64(vm.memory.MemI[s1][a1])

                } else if t1 == Float_t && t3 == Int_t {
                    vm.memory.MemI[s3][a3] = int(vm.memory.MemF[s1][a1])

                } else {
                    fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                    os.Exit(1)
                }
            }

            item, _ := vm.pcstack.Pop()
            vm.pc, _ = item.(int)

            if vm.pcstack.Size() == 0 {
                run_program = false
            } else {
                vm.memory.PopContext()
            }

        case ENDPROC:
            item, _ := vm.pcstack.Pop()
            vm.pc, _ = item.(int)

            if vm.pcstack.Size() == 0 {
                run_program = false
            } else {
                vm.memory.PopContext()
            }

	case ASSERT:
            if t1 == Int_t {
		index := vm.memory.MemI[s1][a1]
		limit := vm.memory.MemI[s3][a3]
		vm.init_ptr = true

		if index < 0 || index >= limit {
                    fmt.Printf("line %d: Index out of range\n", vm.pc)
                    os.Exit(1)
		}
	    }

        case PRINT:
            if t3 == Int_t {
                fmt.Print(vm.memory.MemI[s3][a3], " ")

            } else if t3 == Float_t {
                fmt.Print(vm.memory.MemF[s3][a3], " ")

            } else if t3 == Bool_t {
                fmt.Print(vm.memory.MemB[s3][a3], " ")

            } else if t3 == Float_t {
                fmt.Print(vm.memory.MemI[s3][a3], " ")

            } else if t3 == Int_t {
                fmt.Print(vm.memory.MemF[s3][a3], " ")

            } else if t3 == String_t {
                fmt.Print(vm.memory.Strings[a3], " ")

            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        case PRINTLN:
            if t3 == Int_t {
                fmt.Println(vm.memory.MemI[s3][a3])

            } else if t3 == Float_t {
                fmt.Println(vm.memory.MemF[s3][a3])

            } else if t3 == Bool_t {
                fmt.Println(vm.memory.MemB[s3][a3])

            } else if t3 == Float_t {
                fmt.Println(vm.memory.MemI[s3][a3])

            } else if t3 == Int_t {
                fmt.Println(vm.memory.MemF[s3][a3])

            } else if t3 == String_t {
                fmt.Println(vm.memory.Strings[a3], " ")

            } else {
                fmt.Printf("line %d: Type error mismatch\n", vm.pc)
                os.Exit(1)
            }

        default:
            fmt.Printf(
		"line %d: Unexpected operator at program segment\n",
		vm.pc)
            os.Exit(1)
        }
    }
}
