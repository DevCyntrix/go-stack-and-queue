package queue

type Queue interface {
	Inque(value interface{}) error
	Deque() (interface{}, error)
	IsEmpty() bool
	Size() int
}
