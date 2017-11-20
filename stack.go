package collections

import (
    "container/list"
    //"fmt"
    "errors"
)

// FILO(LIFO): Stack
type Stack struct {
    ll *list.List
    size int
    on_evicted func(interface{})
}

func NewStack(size int) (*Stack, error) {
    if size < 1 {
        return nil, errors.New("size can not be less than 1.")
    }
    s := &Stack{
        ll: list.New(),
        size: size,
    }
    return s, nil
}

func (self *Stack) Push( v interface{}) {
    if self.ll != nil {
        self.ll.PushBack(v)
    }
}

func (self *Stack) Pop() interface{} {
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

func (self *Stack) Len() int {
    if self.ll == nil {
        return -1
    }
    return self.ll.Len()
}
func (self *Stack) Range(pos ...int) []interface{} {
    return nil
}

func (self *Stack) OnEvicted(f func(interface{})) {
    self.on_evicted = f
}