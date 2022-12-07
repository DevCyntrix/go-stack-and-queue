package stack

type LinkedElement struct {
	next  *LinkedElement
	value interface{}
}

type LinkedStack struct {
	top *LinkedElement
}

func NewLinkedStack() Stack {
	return &LinkedStack{}
}

func (s *LinkedStack) Push(value interface{}) error {
	element := new(LinkedElement)
	element.next = s.top
	element.value = value
	s.top = element
	return nil
}

func (s *LinkedStack) Pop() (interface{}, error) {
	top := s.top
	if top == nil {
		return nil, ErrStackIsEmpty
	}
	s.top = top.next
	return top.value, nil
}

func (s *LinkedStack) Top() (interface{}, error) {
	top := s.top
	if top == nil {
		return nil, ErrStackIsEmpty
	}
	return top.value, nil
}

func (s *LinkedStack) IsEmpty() bool {
	return s.top == nil
}
