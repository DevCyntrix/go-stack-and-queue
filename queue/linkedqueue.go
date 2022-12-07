package queue

type LinkedElement struct {
	next     *LinkedElement
	previous *LinkedElement
	value    interface{}
}

type LinkedQueue struct {
	start *LinkedElement
	end   *LinkedElement
	size  int
}

func NewLinkedQueue() Queue {
	return &LinkedQueue{}
}

func (s *LinkedQueue) Enque(value interface{}) error {
	element := new(LinkedElement)
	element.value = value

	if s.start == nil { // First Item
		s.end = element
	} else {
		s.start.previous = element
		element.next = s.start
	}
	s.start = element
	s.size++

	return nil
}

func (s *LinkedQueue) Deque() (interface{}, error) {
	pop := s.end
	if pop == nil {
		return nil, ErrQueueIsEmpty
	}

	s.end = pop.previous

	if s.end == nil {
		s.start = nil
	}
	return pop.value, nil
}

func (s *LinkedQueue) IsEmpty() bool {
	return s.size == 0
}

func (s *LinkedQueue) Size() int {
	return s.size
}
