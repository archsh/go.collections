package collections

const (
    VERSION = "0.0.1"
    AUTHOR = "Mingcai SHEN"
    EMAIL = "archsh@gmail.com"
)


type Collection interface {
    Push(v interface{})
    Pop(v interface{})
    Len() int
    Range(...int) []interface{}
    OnEvicted(func(interface{}))
}


// FIFO: Queue
type Queue struct {

}

// FILO(LIFO): Stack
type Stack struct {

}