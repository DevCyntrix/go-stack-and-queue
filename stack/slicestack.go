package stack

type SliceStack struct {
	readerIndex int
	slice       []interface{}
}

// Creates a new Stack with a dynamic size
func NewSliceStack() *SliceStack {
	return &SliceStack{readerIndex: 0, slice: make([]interface{}, 0)} // This slice has a capacity of size
}

func (s *SliceStack) Push(value interface{}) error {
	s.slice = append(s.slice, value)
	s.readerIndex++
	return nil
}

func (s *SliceStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrStackIsEmpty
	}

	s.readerIndex--
	return s.slice[s.readerIndex], nil
}

func (s *SliceStack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrStackIsEmpty
	}
	i := s.readerIndex - 1
	return s.slice[i], nil
}

func (s *SliceStack) IsEmpty() bool {
	return s.readerIndex <= 0
}
