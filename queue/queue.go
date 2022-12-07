package queue

type Queue interface {
	Enque(value interface{}) error
	Deque() (interface{}, error)
	IsEmpty() bool
	Size() int
}
