package structs

type stackPointer struct {
    stacki      Stack
    stackf      Stack
    stackb      Stack
}

func (sp stackPointer) empty() bool {
    return sp.stacki.Size() == 0
}

func (sp stackPointer) size() int {
    return sp.stacki.Size()
}

func (sp stackPointer) top() (int, int, int) {
    item := sp.stacki.Top()
    spi, _ := item.(int)

    item = sp.stackf.Top()
    spf, _ := item.(int)

    item = sp.stackb.Top()
    spb, _ := item.(int)

    return spi, spf, spb
}

func (sp stackPointer) push(sizei, sizef, sizeb int) {
    sp.stacki.Push(sizei)
    sp.stackf.Push(sizef)
    sp.stackb.Push(sizeb)
}

func (sp stackPointer) pop() (int, int, int) {
    item, _ := sp.stacki.Pop()
    spi, _ := item.(int)

    item, _ = sp.stackf.Pop()
    spf, _ := item.(int)

    item, _ = sp.stackb.Pop()
    spb, _ := item.(int)

    return spi, spf, spb
}

