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
        return nil, errors.New("size can not be less than 1. ")
    }
    s := &Stack{
        ll: list.New(),
        size: size,
    }
    return s, nil
}

func (self *Stack) Push( v interface{}) error {
    if nil == self.ll {
        return errors.New("Not initialized! ")
    }
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
    return nil
}

func (self *Stack) Pop() (interface{}, error) {
    if self.ll == nil {
        return nil,errors.New("Not initialized! ")
    }
    e := self.ll.Back()
    if nil == e {
        return nil,errors.New("Empty! ")
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
func (self *Stack) Range(pos ...int) []interface{} {
    if nil == self.ll {
        return nil
    }
    var i, begin, limit int
    //i := 0
    //begin := 0
    //limit := self.ll.Len()
    var result []interface{}
    switch len(pos) {
    case 0:
        begin = 0
        limit = self.ll.Len()
    case 1:
        begin = pos[0]
        limit = self.ll.Len()
    case 2:
        begin = pos[0]
        limit = pos[1]
    default:
        return nil
    }
    if begin < 0 {
        begin = self.ll.Len() + begin
    }
    if limit <= 0 {
        return nil
    }
    for e := self.ll.Back(); e != nil; e = e.Prev() {
        if limit <= 0 {
            break
        }
        if i>= begin {
            result = append(result, e.Value)
        }
        i += 1
    }
    return result
}

func (self *Stack) OnEvicted(f func(interface{})) {
    self.on_evicted = f
}