package main

import (
    "bufio"
    "os"
    "strconv"
    "./structs"
)

func run(filename string, debug bool) {
    var code []structs.Tac

    file, err := os.Open(filename)
    checkError(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)

    var curr_code structs.Tac
    count := 0

    for ; scanner.Scan(); count++ {
        state := count % 4

        if state == 0 {
            op, err := strconv.ParseInt(scanner.Text(), 16, 64)
            checkError(err)
            checkError(structs.ValidateOp(int(op)))
            curr_code.Op = int(op)
        } else {
            i := state - 1

            if curr_code.Op == structs.SET && !(state == 3) {
                curr_code.Args[i] = scanner.Text()
            } else {
                op, err := strconv.ParseInt(scanner.Text(), 16, 64)
                checkError(err)
                curr_code.Address[i] = int(op)
            }

            if state == 3 {
                code = append(code, curr_code)
                curr_code.Clear()
            }
        }
    }

    vm := structs.VirtualMachine{Program: code, Debug: debug}
    vm.InitMemory()
    vm.Execute()
}
