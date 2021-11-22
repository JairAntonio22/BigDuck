package structs

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "sort"
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
                fmt.Printf("Local address used in data segment\n")
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
                fmt.Printf("Invalid address used in data segment\n")
                os.Exit(1)
            }

        case PROGRAM:
            vm.basepc = vm.pc
            break

        default:
            fmt.Printf(
                "Unexpected operator %s at data segment\n",
                TypeToString[curr_code.Op])
            os.Exit(1)
        }
    }
}

func (vm *VirtualMachine) Execute() {
    var curr_code Tac
    run_program := true

    reader := bufio.NewReader(os.Stdin)

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
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case OR:
            if t1 == Bool_t && t2 == Bool_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemB[s1][a1] || vm.memory.MemB[s2][a2])
            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case AND:
            if t1 == Bool_t && t2 == Bool_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemB[s1][a1] && vm.memory.MemB[s2][a2])
            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case NOT:
            if t1 == Bool_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = !vm.memory.MemB[s1][a1]
            } else {
                fmt.Printf("Type error mismatch\n")
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
                fmt.Printf("Type error mismatch\n")
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
                fmt.Printf("Type error mismatch\n")
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
                fmt.Printf("Type error mismatch\n")
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
                fmt.Printf("Type error mismatch\n")
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
                fmt.Printf("Type error mismatch\n")
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
                fmt.Printf("Type error mismatch\n")
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
                fmt.Printf("Type error mismatch\n")
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
                fmt.Printf("Type error mismatch\n")
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
                fmt.Printf("Type error mismatch\n")
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
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case LPAREN:
            fmt.Printf("Unexpected operator %s\n", TypeToString[curr_code.Op])
            os.Exit(1)

        case RPAREN:
            fmt.Printf("Unexpected operator %s\n", TypeToString[curr_code.Op])
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
                    fmt.Printf("Type error mismatch\n")
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
                fmt.Printf("Type error mismatch\n")
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
                    fmt.Printf("Type error mismatch\n")
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
                    fmt.Printf("Index out of range\n")
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
                fmt.Printf("Type error mismatch\n")
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
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case READ:
            text, err := reader.ReadString('\n')

            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }

            text = text[:len(text) - 1]

            switch t3 {
            case Int_t:
                value, _ := strconv.ParseInt(text, 10, 64)
                vm.memory.MemI[s3][a3] = int(value)

            case Float_t:
                value, _ := strconv.ParseInt(text, 10, 64)
                vm.memory.MemF[s3][a3] = float64(value)

            case Bool_t:
                if text == "#t" {
                    vm.memory.MemB[s3][a3] = true
                } else {
                    vm.memory.MemB[s3][a3] = false
                }

            default:
                fmt.Printf(
                    "Invalid input, expecting %s\n",
                    TypeToString[t3])
                os.Exit(1)
            }

        case SIN:
            if t1 == Int_t {
                vm.memory.MemF[s3][a3] = math.Sin(float64(
                    vm.memory.MemI[s1][a1]))

            } else if t1 == Float_t {
                vm.memory.MemF[s3][a3] = math.Sin(
                    vm.memory.MemF[s1][a1])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case ASIN:
            if t1 == Int_t {
                vm.memory.MemF[s3][a3] = math.Asin(float64(
                    vm.memory.MemI[s1][a1]))

            } else if t1 == Float_t {
                vm.memory.MemF[s3][a3] = math.Asin(
                    vm.memory.MemF[s1][a1])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case COS:
            if t1 == Int_t {
                vm.memory.MemF[s3][a3] = math.Cos(float64(
                    vm.memory.MemI[s1][a1]))

            } else if t1 == Float_t {
                vm.memory.MemF[s3][a3] = math.Cos(
                    vm.memory.MemF[s1][a1])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case ACOS:
            if t1 == Int_t {
                vm.memory.MemF[s3][a3] = math.Acos(float64(
                    vm.memory.MemI[s1][a1]))

            } else if t1 == Float_t {
                vm.memory.MemF[s3][a3] = math.Acos(
                    vm.memory.MemF[s1][a1])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case TAN:
            if t1 == Int_t {
                vm.memory.MemF[s3][a3] = math.Tan(float64(
                    vm.memory.MemI[s1][a1]))

            } else if t1 == Float_t {
                vm.memory.MemF[s3][a3] = math.Tan(
                    vm.memory.MemF[s1][a1])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case ATAN:
            if t1 == Int_t {
                vm.memory.MemF[s3][a3] = math.Atan(float64(
                    vm.memory.MemI[s1][a1]))

            } else if t1 == Float_t {
                vm.memory.MemF[s3][a3] = math.Atan(
                    vm.memory.MemF[s1][a1])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case ATAN2:
            if t1 == Int_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = math.Atan2(
                    float64(vm.memory.MemI[s1][a1]),
                    float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Int_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = math.Atan2(
                    float64(vm.memory.MemI[s1][a1]), vm.memory.MemF[s2][a2])

            } else if t1 == Float_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = math.Atan2(
                    vm.memory.MemF[s1][a1], float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Float_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = math.Atan2(
                    vm.memory.MemF[s1][a1], vm.memory.MemF[s2][a2])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case EXP:
            if t1 == Int_t {
                vm.memory.MemF[s3][a3] = math.Exp(float64(
                    vm.memory.MemI[s1][a1]))

            } else if t1 == Float_t {
                vm.memory.MemF[s3][a3] = math.Exp(
                    vm.memory.MemF[s1][a1])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case LN:
            if t1 == Int_t {
                vm.memory.MemF[s3][a3] = math.Log(float64(
                    vm.memory.MemI[s1][a1]))

            } else if t1 == Float_t {
                vm.memory.MemF[s3][a3] = math.Exp(
                    vm.memory.MemF[s1][a1])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case SQRT:
            if t1 == Int_t {
                vm.memory.MemF[s3][a3] = math.Sqrt(float64(
                    vm.memory.MemI[s1][a1]))

            } else if t1 == Float_t {
                vm.memory.MemF[s3][a3] = math.Sqrt(
                    vm.memory.MemF[s1][a1])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case POW:
            if t1 == Int_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = math.Pow(
                    float64(vm.memory.MemI[s1][a1]),
                    float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Int_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = math.Pow(
                    float64(vm.memory.MemI[s1][a1]), vm.memory.MemF[s2][a2])

            } else if t1 == Float_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = math.Pow(
                    vm.memory.MemF[s1][a1], float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Float_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = math.Pow(
                    vm.memory.MemF[s1][a1], vm.memory.MemF[s2][a2])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case LOG:
            if t1 == Int_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = Log(
                    float64(vm.memory.MemI[s1][a1]),
                    float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Int_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = Log(
                    float64(vm.memory.MemI[s1][a1]), vm.memory.MemF[s2][a2])

            } else if t1 == Float_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = Log(
                    vm.memory.MemF[s1][a1], float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Float_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = Log(
                    vm.memory.MemF[s1][a1], vm.memory.MemF[s2][a2])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case MOD:
            if t1 == Int_t && t2 == Int_t && t3 == Int_t {
                vm.memory.MemI[s3][a3] = (
                    vm.memory.MemI[s1][a1] %  vm.memory.MemI[s2][a2])

            } else if t1 == Int_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = math.Mod(
                    float64(vm.memory.MemI[s1][a1]), vm.memory.MemF[s2][a2])

            } else if t1 == Float_t && t2 == Int_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = math.Mod(
                    vm.memory.MemF[s1][a1], float64(vm.memory.MemI[s2][a2]))

            } else if t1 == Float_t && t2 == Float_t && t3 == Float_t {
                vm.memory.MemF[s3][a3] = Log(
                    vm.memory.MemF[s1][a1], vm.memory.MemF[s2][a2])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case ABS:
            if t1 == Int_t {
                vm.memory.MemF[s3][a3] = math.Abs(float64(
                    vm.memory.MemI[s1][a1]))

            } else if t1 == Float_t {
                vm.memory.MemF[s3][a3] = math.Abs(
                    vm.memory.MemF[s1][a1])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case CEIL:
            if t1 == Int_t {
                vm.memory.MemF[s3][a3] = math.Ceil(float64(
                    vm.memory.MemI[s1][a1]))

            } else if t1 == Float_t {
                vm.memory.MemF[s3][a3] = math.Ceil(
                    vm.memory.MemF[s1][a1])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case FLOOR:
            if t1 == Int_t {
                vm.memory.MemF[s3][a3] = math.Floor(float64(
                    vm.memory.MemI[s1][a1]))

            } else if t1 == Float_t {
                vm.memory.MemF[s3][a3] = math.Floor(
                    vm.memory.MemF[s1][a1])

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case MEAN:
            if t1 == Int_t && t2 == Int_t && t3 == Float_t {
		var begin int

		if s1 == Local {
		    begin = vm.memory.Sp[t1] + GetAddress(va1)
		} else {
		    begin = GetAddress(va1)
		}

		size := vm.memory.MemI[s2][a2]
		vector := vm.memory.MemI[s1][begin: begin + size]

		sum := 0

		for _, elem := range vector {
		    sum += elem
		}

                vm.memory.MemF[s3][a3] = float64(sum) / float64(size);

            } else if t1 == Float_t && t2 == Int_t && t3 == Float_t {
		var begin int

		if s1 == Local {
		    begin = vm.memory.Sp[t1] + GetAddress(va1)
		} else {
		    begin = GetAddress(va1)
		}

		size := vm.memory.MemI[s2][a2]
		vector := vm.memory.MemF[s1][begin: begin + size]

		sum := 0.0

		for _, elem := range vector {
		    sum += elem
		}

                vm.memory.MemF[s3][a3] = sum / float64(size);

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case MEDIAN:
            if t1 == Int_t && t2 == Int_t && t3 == Float_t {
		var begin int

		if s1 == Local {
		    begin = vm.memory.Sp[t1] + GetAddress(va1)
		} else {
		    begin = GetAddress(va1)
		}

		size := vm.memory.MemI[s2][a2]
		items := vm.memory.MemI[s1][begin: begin + size]
		vector := make([]int, len(items))
		_ = copy(vector, items)

		sort.Slice(vector, func(i, j int) bool {
		    return vector[i] < vector[j]
		})

		mid := len(vector) / 2

		if len(vector) % 2 == 1 {
		    vm.memory.MemF[s3][a3] = float64(vector[mid])
		} else {
		    vm.memory.MemF[s3][a3] = float64(
			vector[mid / 2] + vector[mid + 1]) / 2.0
		}

            } else if t1 == Float_t && t2 == Int_t && t3 == Float_t {
		var begin int

		if s1 == Local {
		    begin = vm.memory.Sp[t1] + GetAddress(va1)
		} else {
		    begin = GetAddress(va1)
		}

		size := vm.memory.MemI[s2][a2]
		items := vm.memory.MemF[s1][begin: begin + size]
		vector := make([]float64, len(items))
		_ = copy(vector, items)

		sort.Slice(vector, func(i, j int) bool {
		    return vector[i] < vector[j]
		})

		if size % 2 == 1 {
		    vm.memory.MemF[s3][a3] = float64(
			vector[len(vector) / 2] + vector[(len(vector) / 2) + 1]) / 2.0
		} else {
		    vm.memory.MemF[s3][a3] = float64(vector[len(vector) / 2])
		}

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }

        case MODE:
            if t1 == Int_t && t2 == Int_t && t3 == Float_t {
		var begin int

		if s1 == Local {
		    begin = vm.memory.Sp[t1] + GetAddress(va1)
		} else {
		    begin = GetAddress(va1)
		}

		size := vm.memory.MemI[s2][a2]
		vector := vm.memory.MemI[s1][begin: begin + size]

		count := make(map[int]int)

		for _, elem := range vector {
		    _, exists := count[elem]

		    if exists {
			count[elem]++
		    } else {
			count[elem] = 1
		    }
		}

		max_key := vector[0]
		max_value := count[vector[0]]

		for key, value := range count {
		    if value > max_value {
			max_key = key
			max_value = value
		    }
		}

                vm.memory.MemF[s3][a3] = float64(max_key);

            } else if t1 == Float_t && t2 == Int_t && t3 == Float_t {
		var begin int

		if s1 == Local {
		    begin = vm.memory.Sp[t1] + GetAddress(va1)
		} else {
		    begin = GetAddress(va1)
		}

		size := vm.memory.MemI[s2][a2]
		vector := vm.memory.MemF[s1][begin: begin + size]

		count := make(map[float64]int)

		for _, elem := range vector {
		    _, exists := count[elem]

		    if exists {
			count[elem]++
		    } else {
			count[elem] = 1
		    }
		}

		max_key := vector[0]
		max_value := count[vector[0]]

		for key, value := range count {
		    if value > max_value {
			max_key = key
			max_value = value
		    }
		}

                vm.memory.MemF[s3][a3] = max_key;

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(1)
            }


        default:
            fmt.Printf( "Unexpected operator at program segment\n")
            os.Exit(1)
        }
    }
}

func Log(base, x float64) float64 {
    return math.Log(x) / math.Log(base)
}
