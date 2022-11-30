package queue

type SliceQueue struct {
	readerIndex int
	slice       []interface{}
}

func NewSliceQueue() Queue {
	return &SliceQueue{readerIndex: 0, slice: make([]interface{}, 0)}
}

func (q *SliceQueue) optimize() {
	if q.readerIndex < 16 {
		return
	}
	q.slice = q.slice[q.readerIndex:] // Maybe this won't reduce the size of the queue. (allocate and copy)
	q.readerIndex = 0
}

func (q *SliceQueue) Inque(value interface{}) error {
	q.slice = append(q.slice, value)
	return nil
}

func (q *SliceQueue) Deque() (interface{}, error) {
	if q.readerIndex == len(q.slice) {
		return nil, ErrQueueIsEmpty
	}
	value := q.slice[q.readerIndex]
	q.readerIndex++
	q.optimize()
	return value, nil
}

func (q *SliceQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *SliceQueue) Size() int {
	return len(q.slice) - q.readerIndex
}
