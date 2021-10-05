package structs

import (
    "errors"
)

type Stack struct {
    data LinkedList
}

func (s Stack) Print() {
    s.data.Print()
}

func (s Stack) Empty() bool {
    return s.data.size == 0
}

func (s Stack) Size() int {
    return s.data.size
}

func (s Stack) Top() interface{} {
    return s.data.back.item
}

func (s *Stack) Push(item interface{}) {
    s.data.PushBack(item)
}

func (s *Stack) Pop() (item interface{}, err error) {
    item, err = s.data.PopBack()

    if err != nil {
        err = errors.New("Stack.Pop: empty stack")
    }

    return
}
