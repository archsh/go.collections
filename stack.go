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

func (self *Stack) Push( v interface{}) error {
    if self.ll != nil {
        if self.ll.Len() >= self.size {
            e, r := self.Pop()
            if nil != r {
                return r
            }
            if nil != self.on_evicted {
                self.on_evicted(e)
            }
        }
        self.ll.PushBack(v)
    }
    return nil
}

func (self *Stack) Pop() (interface{}, error) {
    if self.ll == nil {
        return nil,nil
    }
    e := self.ll.Back()
    if nil == e {
        return nil,nil
    }
    v := e.Value
    self.ll.Remove(e)
    return v,nil
}

func (self *Stack) Len() int {
    if self.ll == nil {
        return -1
    }
    return self.ll.Len()
}
func (self *Stack) Range(pos ...int) ([]interface{},error) {
    return nil,nil
}

func (self *Stack) OnEvicted(f func(interface{})) {
    self.on_evicted = f
}