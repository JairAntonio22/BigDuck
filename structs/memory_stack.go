package structs

type memoryStack struct {
    stacki      Stack
    stackf      Stack
    stackb      Stack
}

func (ms memoryStack) Empty() bool {
    return ms.stacki.Size() == 0
}

func (ms memoryStack) Size() int {
    return ms.stacki.Size()
}

func (ms memoryStack) Top() (int, int, int) {
    item := ms.stacki.Top()
    msi, _ := item.(int)

    item = ms.stackf.Top()
    msf, _ := item.(int)

    item = ms.stackb.Top()
    msb, _ := item.(int)

    return msi, msf, msb
}

func (ms *memoryStack) Push(sizei, sizef, sizeb int) {
    ms.stacki.Push(sizei)
    ms.stackf.Push(sizef)
    ms.stackb.Push(sizeb)
}

func (ms *memoryStack) Pop() (int, int, int) {
    item, _ := ms.stacki.Pop()
    msi, _ := item.(int)

    item, _ = ms.stackf.Pop()
    msf, _ := item.(int)

    item, _ = ms.stackb.Pop()
    msb, _ := item.(int)

    return msi, msf, msb
}

