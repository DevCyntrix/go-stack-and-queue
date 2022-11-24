package stackarray

import (
	"errors"
)

var (
	StackIsFull  = errors.New("The stack is full")
	StackIsEmpty = errors.New("The stack is empty")
)

// LIFO (Last in First out)
type Stack struct {
	index int
	array []interface{}
}

// Creates a new Stack with a constant size
func New(size int32) *Stack {
	return &Stack{array: make([]interface{}, size)} // This slice has a capacity of size
}

// Pushs the element on the top of the stack, otherwise it returns the stack is full error
func (stack *Stack) Push(value interface{}) error {
	if stack.index >= len(stack.array) {
		return StackIsFull
	}

	stack.array[stack.index] = value
	stack.index++
	return nil
}

// Returns the last element which were put in otherwise it returns the stack is empty error
func (stack *Stack) Pop() (interface{}, error) {
	if stack.index <= 0 {
		return nil, StackIsEmpty
	}
	stack.index--
	return stack.array[stack.index], nil
}
