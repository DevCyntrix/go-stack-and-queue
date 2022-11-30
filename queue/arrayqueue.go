package queue

import (
	"errors"
	"fmt"
)

// FIFO
type ArrayQueue struct {
	readerIndex int
	writerIndex int
	array       []interface{}
}

var (
	ErrQueueIsEmpty    = errors.New("the queue is empty")
	ErrArrayOutOfBound = errors.New("array is out of bound")
)

func NewArrayQueue(size int) *ArrayQueue {
	return &ArrayQueue{readerIndex: 0, writerIndex: 0, array: make([]interface{}, size)}
}

func (q *ArrayQueue) Inque(value interface{}) error {

	if q.writerIndex >= q.Size() && (q.writerIndex%q.Size() == q.readerIndex%q.Size()) {
		return ErrArrayOutOfBound
	}

	q.array[q.writerIndex%q.Size()] = value
	q.writerIndex++
	fmt.Println(q.array)
	return nil
}

func (q *ArrayQueue) Deque() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueIsEmpty
	}
	value := q.array[q.readerIndex%q.Size()]
	q.readerIndex++
	return value, nil
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.writerIndex == q.readerIndex
}

func (q *ArrayQueue) Size() int {
	return len(q.array)
}
