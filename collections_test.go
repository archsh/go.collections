package collections

import (
    "testing"
    "time"
)


func TestNewQueue(t *testing.T) {
    t1 := time.Now()
    q, e := NewQueue(100)
    if nil != e {
        t.Error("New Queue Failed:>", e)
        t.Failed()
    }
    for i:=1; i<1000; i++ {
        q.Push(i)
    }
    for v, e := q.Pop(); e == nil && v != nil; v, e = q.Pop() {
        t.Log("V:>", v)
    }
    t2 := time.Now()
    t.Log("T1:>", t1, "T2:>", t2, "Total:>", t2.Sub(t1))
}

func TestNewStack(t *testing.T) {
    t1 := time.Now()
    q, e := NewStack(100)
    if nil != e {
        t.Error("New Stack Failed:>", e)
        t.Failed()
    }
    for i:=1; i<1000; i++ {
        q.Push(i)
    }
    for v, e := q.Pop(); e == nil && v != nil; v, e = q.Pop() {
        t.Log("V:>", v)
    }
    t2 := time.Now()
    t.Log("T1:>", t1, "T2:>", t2, "Total:>", t2.Sub(t1))
}