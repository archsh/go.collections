package collections

const (
    VERSION = "0.0.1"
    AUTHOR = "Mingcai SHEN"
    EMAIL = "archsh@gmail.com"
)

type Collection interface {
    Push(interface{}) error
    Pop() (interface{},error)
    Len() int
    Range(...int) ([]interface{},error)
    OnEvicted(func(interface{}))
}
