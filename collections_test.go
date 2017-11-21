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
    q.OnEvicted(func(v interface{}){
        t.Log("Evicted:>", v)
    })
    for i:=1; i<1000; i++ {
        q.Push(i)
    }
    for _, x := range q.Range(-10) {
        t.Log("X1:>", x)
    }
    for _, x := range q.Range(50,60) {
        t.Log("X2:>", x)
    }
    for _, x := range q.Range(50) {
        t.Log("X3:>", x)
    }
    for _, x := range q.Range(90,0) {
        t.Log("X4:>", x)
    }
    for v, e := q.Pop(); e == nil && v != nil; v, e = q.Pop() {
        t.Log("V:>", v)
    }
    t2 := time.Now()
    t.Log("T1:>", t1, "T2:>", t2, "Total:>", t2.Sub(t1))
}

func TestNewStack(t *testing.T) {
    t1 := time.Now()
    s, e := NewStack(100)
    if nil != e {
        t.Error("New Stack Failed:>", e)
        t.Failed()
    }
    s.OnEvicted(func(v interface{}){
        t.Log("Evicted:>", v)
    })
    for i:=1; i<1000; i++ {
        s.Push(i)
    }
    for _, x := range s.Range(-10) {
        t.Log("X1:>", x)
    }
    for _, x := range s.Range(50,60) {
        t.Log("X2:>", x)
    }
    for _, x := range s.Range(50) {
        t.Log("X3:>", x)
    }
    for _, x := range s.Range(90,0) {
        t.Log("X4:>", x)
    }
    for v, e := s.Pop(); e == nil && v != nil; v, e = s.Pop() {
        t.Log("V:>", v)
    }
    t2 := time.Now()
    t.Log("T1:>", t1, "T2:>", t2, "Total:>", t2.Sub(t1))
}