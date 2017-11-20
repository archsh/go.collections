package collections

const (
    VERSION = "0.0.1"
    AUTHOR = "Mingcai SHEN"
    EMAIL = "archsh@gmail.com"
)


type Collection interface {
    Push(interface{})
    Pop() interface{}
    Len() int
    Range(...int) []interface{}
    OnEvicted(func(interface{}))
}
