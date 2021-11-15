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
                os.Exit(0)
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

            default:
                fmt.Printf("Invalid address used in data segment\n")
                os.Exit(0)
            }

        case PROGRAM:
            vm.basepc = vm.pc
            break

        default:
            fmt.Printf(
                "Unexpected operator %s at data segment\n",
                TypeToString[curr_code.Op])
            os.Exit(0)
        }
    }
}

func (vm *VirtualMachine) Execute() {
    var curr_code Tac

    for ; vm.pc < len(vm.Program); vm.pc++ {
        curr_code =  vm.Program[vm.pc]

        if vm.Debug {
            fmt.Printf("%3d ", vm.pc - vm.basepc)
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

        a1 := vm.memory.Sp[Prev][t1] + GetAddress(va1)
        a2 := vm.memory.Sp[Prev][t2] + GetAddress(va2)
        a3 := vm.memory.Sp[Prev][t3] + GetAddress(va3)

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
                os.Exit(0)
            }

        case OR:
            if t1 == Bool_t && t2 == Bool_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemB[s1][a1] || vm.memory.MemB[s2][a2])
            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(0)
            }

        case AND:
            if t1 == Bool_t && t2 == Bool_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = (
                    vm.memory.MemB[s1][a1] && vm.memory.MemB[s2][a2])
            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(0)
            }

        case NOT:
            if t1 == Bool_t && t3 == Bool_t {
                vm.memory.MemB[s3][a3] = !vm.memory.MemB[s1][a1]
            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(0)
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
                os.Exit(0)
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
                os.Exit(0)
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
                os.Exit(0)
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
                os.Exit(0)
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
                os.Exit(0)
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
                os.Exit(0)
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
                os.Exit(0)
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
                os.Exit(0)
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
                os.Exit(0)
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
                os.Exit(0)
            }

        case LPAREN:
            fmt.Printf("Unexpected operator %s\n", TypeToString[curr_code.Op])
            os.Exit(0)

        case RPAREN:
            fmt.Printf("Unexpected operator %s\n", TypeToString[curr_code.Op])
            os.Exit(0)

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

        case PROC:
        case GOPROC:
            vm.pc = vm.basepc + curr_code.Address[2]
            vm.memory.ChangeContext()

        case ERA:
            vm.memory.InitLocal(
                curr_code.Address[0],
                curr_code.Address[1],
                curr_code.Address[2])

        case PARAM:
        case RETURN:
        case ENDPROC:

        case PRINT:
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

            } else {
                fmt.Printf("Type error mismatch\n")
                os.Exit(0)
            }

        default:
            fmt.Printf("Unexpected operator at program segment\n")
            os.Exit(0)
        }
    }
}
