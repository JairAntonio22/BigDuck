package structs

import (
    "fmt"
    "os"
    "strconv"
)

type VirtualMachine struct {
    Program     []Tac       // Program code
    memory      memory      // Memory manager
    basepc      int         // Pointer to program start
    pc          int         // Program counter
}

func (vm *VirtualMachine) InitMemory() {
    var curr_code Tac

    for ; curr_code.Op != PROGRAM; vm.pc++ {
        curr_code =  vm.Program[vm.pc]
        curr_code.Print()

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
}
