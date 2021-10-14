package structs

import (
    "errors"
    "fmt"
)

type node struct {
    item interface{}
    next *node
    prev *node
}

type LinkedList struct {
    front *node
    back  *node
    size  int
}

func (list LinkedList) Print() {
    if list.size == 0 {
        fmt.Println("(empty)")
    } else {
        fmt.Printf("[ ")

        for it := list.front; it != nil; it = it.next {
            fmt.Printf("%v ", it.item)
        }

        fmt.Printf("]\n")
    }
}

func (list *LinkedList) PushFront(item interface{}) {
    if list.size == 0 {
        list.front = &node{item: item}
        list.back = list.front

    } else {
        list.front.prev = &node{
            item: item,
            next: list.front,
        }

        list.front = list.front.prev
    }

    list.size++
}

func (list *LinkedList) PushBack(item interface{}) {
    if list.size == 0 {
        list.front = &node{item: item}
        list.back = list.front

    } else {
        list.back.next = &node{
            item: item,
            prev: list.back,
        }

        list.back = list.back.next
    }

    list.size++
}

func (list *LinkedList) PopFront() (item interface{}, err error) {
    if list.size == 0 {
        err = errors.New("LinkedList.PopFront: empty list")
    } else {
        item = list.front.item
        list.front = list.front.next
        list.size--

        if list.size > 0 {
            list.front.prev = nil
        }
    }

    return
}

func (list *LinkedList) PopBack() (item interface{}, err error) {
    if list.size == 0 {
        err = errors.New("LinkedList.PopBack: empty list")
    } else {
        item = list.back.item
        list.back = list.back.prev
        list.size--

        if list.size > 0 {
            list.back.next = nil
        }
    }

    return
}
