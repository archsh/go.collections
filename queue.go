package collections

import (
    "container/list"
    //"fmt"
    "errors"
)

// FIFO: Queue
type Queue struct {
    ll *list.List
    size int
    on_evicted func(interface{})
}

func NewQueue(size int) (*Queue, error) {
    if size < 1 {
        return nil, errors.New("size can not be less than 1.")
    }
    q := &Queue{
        ll: list.New(),
        size: size,
    }
    return q, nil
}

func (self *Queue) Push( v interface{}) {
    if self.ll != nil {
        self.ll.PushBack(v)
    }
}

func (self *Queue) Pop() interface{} {
    if self.ll == nil {
        return nil
    }
    e := self.ll.Front()
    if nil == e {
        return nil
    }
    v := e.Value
    self.ll.Remove(e)
    return v
}

func (self *Queue) Len() int {
    if self.ll == nil {
        return -1
    }
    return self.ll.Len()
}
func (self *Queue) Range(pos ...int) []interface{} {
    return nil
}

func (self *Queue) OnEvicted(f func(interface{})) {
    self.on_evicted = f
}