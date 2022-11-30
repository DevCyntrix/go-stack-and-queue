package stack

// LIFO (Last in First out)
type ArrayStack struct {
	index int
	array []interface{}
}

// Creates a new Stack with a constant size
func NewArrayStack(size int32) *ArrayStack {
	return &ArrayStack{array: make([]interface{}, size)} // This slice has a capacity of size
}

// Pushs the element on the top of the stack, otherwise it returns the stack is full error
func (stack *ArrayStack) Push(value interface{}) error {
	if stack.index >= len(stack.array) {
		return ErrStackIsFull
	}
	stack.array[stack.index] = value
	stack.index++
	return nil
}

// Returns the last element which were put in otherwise it returns the stack is empty error
func (stack *ArrayStack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, ErrStackIsEmpty
	}
	stack.index--
	return stack.array[stack.index], nil
}

// Returns the element on the top of the stack
func (stack *ArrayStack) Top() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, ErrStackIsEmpty
	}
	i := stack.index - 1
	return stack.array[i], nil
}

// Returns true if the stack has no elements
func (stack *ArrayStack) IsEmpty() bool {
	return stack.index <= 0
}
