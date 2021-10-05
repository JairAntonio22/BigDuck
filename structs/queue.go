package structs

import "errors"

type Queue struct {
    data LinkedList
}

func (q Queue) Print() {
    q.data.Print()
}

func (q Queue) Empty() bool {
    return q.data.size == 0
}

func (q Queue) Size() int {
    return q.data.size
}

func (q Queue) Front() interface{} {
    return q.data.front.item
}

func (q Queue) Back() interface{} {
    return q.data.back.item
}

func (q *Queue) Push(item interface{}) {
    q.data.PushBack(item)
}

func (q *Queue) Pop() (item interface{}, err error) {
    item, err = q.data.PopFront()

    if err != nil {
        err = errors.New("Queue.Pop: empty queue")
    }

    return
}
