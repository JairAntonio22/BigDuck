package structs

import (
    "fmt"
    "os"
    "strconv"
)

type VirtualMachine struct {
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
            vaddress := curr_code.Address[2]
            scope := GetScope(vaddress)
            address := GetAddress(vaddress)

            if scope != Global {
                fmt.Printf("Local address used in data segment")
                os.Exit(0)
            }

            switch GetType(vaddress) {
            case Int_t:
                value, _ := strconv.ParseInt(curr_code.Args[0], 10, 64)
                vm.memory.MemI[scope][address] = int(value)

            case Float_t:
                value, _ := strconv.ParseFloat(curr_code.Args[0], 10)
                vm.memory.MemF[scope][address] = value

            case Bool_t:
                if curr_code.Args[0] == "#t" {
                    vm.memory.MemB[scope][address] = true
                } else {
                    vm.memory.MemB[scope][address] = false
                }

            default:
                fmt.Printf("Invalid address used in data segment")
                os.Exit(0)
            }

        case PROGRAM:
            break

        default:
            fmt.Printf(
                "Unexpected operator %s at data segment\n",
                TypeToString[curr_code.Op])
            os.Exit(0)
        }
    }

    vm.basepc = vm.pc
}

func (vm *VirtualMachine) Execute() {
    var curr_code Tac

    for ; vm.pc < len(vm.Program); vm.pc++ {
        curr_code =  vm.Program[vm.pc]
        curr_code.Print()

        switch curr_code.Op {
        case ASG:
            vaddress1 := curr_code.Address[0]
            vaddress2 := curr_code.Address[2]

            scope1 := GetScope(vaddress1)
            scope2 := GetScope(vaddress2)

            type1 := GetType(vaddress1)
            type2 := GetType(vaddress2)

            address1 := GetAddress(vaddress1)
            address2 := GetAddress(vaddress2)


            if type1 == Int_t && type2 == Int_t {
                vm.memory.MemI[scope2][address2] = vm.memory.MemI[scope1][address1]

            } else if type1 == Float_t && type2 == Float_t {
                vm.memory.MemF[scope2][address2] = vm.memory.MemF[scope1][address1]

            } else if type1 == Bool_t && type2 == Bool_t {
                vm.memory.MemB[scope2][address2] = vm.memory.MemB[scope1][address1]

            } else if type1 == Int_t && type2 == Float_t {
                vm.memory.MemI[scope2][address2] = int(vm.memory.MemF[scope1][address1])

            } else if type1 == Float_t && type2 == Int_t {
                vm.memory.MemF[scope2][address2] = float64(vm.memory.MemI[scope1][address1])

            } else {
                fmt.Printf("Type error mismatch")
                os.Exit(0)
            }

        case OR:
            vaddress1 := curr_code.Address[0]
            vaddress2 := curr_code.Address[1]
            vaddress3 := curr_code.Address[2]

            scope1 := GetScope(vaddress1)
            scope2 := GetScope(vaddress2)
            scope3 := GetScope(vaddress3)

            type1 := GetType(vaddress1)
            type2 := GetType(vaddress2)
            type3 := GetType(vaddress3)

            address1 := GetAddress(vaddress1)
            address2 := GetAddress(vaddress2)
            address3 := GetAddress(vaddress3)

            if type1 == Bool_t && type2 == Bool_t && type3 == Bool_t {
                vm.memory.MemB[scope3][address3] = (
                    vm.memory.MemB[scope1][address1] || vm.memory.MemB[scope2][address2])
            } else {
                fmt.Printf("Type error mismatch")
                os.Exit(0)
            }

        case AND:
            vaddress1 := curr_code.Address[0]
            vaddress2 := curr_code.Address[1]
            vaddress3 := curr_code.Address[2]

            scope1 := GetScope(vaddress1)
            scope2 := GetScope(vaddress2)
            scope3 := GetScope(vaddress3)

            type1 := GetType(vaddress1)
            type2 := GetType(vaddress2)
            type3 := GetType(vaddress3)

            address1 := GetAddress(vaddress1)
            address2 := GetAddress(vaddress2)
            address3 := GetAddress(vaddress3)

            if type1 == Bool_t && type2 == Bool_t && type3 == Bool_t {
                vm.memory.MemB[scope3][address3] = (
                    vm.memory.MemB[scope1][address1] && vm.memory.MemB[scope2][address2])
            } else {
                fmt.Printf("Type error mismatch")
                os.Exit(0)
            }

        case NOT:
            vaddress1 := curr_code.Address[0]
            vaddress2 := curr_code.Address[2]

            scope1 := GetScope(vaddress1)
            scope2 := GetScope(vaddress2)

            type1 := GetType(vaddress1)
            type2 := GetType(vaddress2)

            address1 := GetAddress(vaddress1)
            address2 := GetAddress(vaddress2)

            if type1 == Bool_t && type2 == Bool_t {
                vm.memory.MemB[scope2][address2] = !vm.memory.MemB[scope1][address1]
            } else {
                fmt.Printf("Type error mismatch")
                os.Exit(0)
            }

        case EQ:

        case NEQ:

        case LES:
        case GRE:
        case LEQ:
        case GEQ:

        case SUB:
        case ADD:
        case DIV:
        case MUL:

        case LPAREN:
        case RPAREN:

        case JMP:
        case JMT:
        case JMF:

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

        case SET:
        case PROGRAM:

        default:
            fmt.Printf(
                "Unexpected operator %s at program segment\n",
                TypeToString[curr_code.Op])
            os.Exit(0)
        }
    }
}
