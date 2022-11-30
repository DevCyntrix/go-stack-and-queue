package stack

import (
	"errors"
)

var (
	ErrStackIsFull  = errors.New("the stack is full")
	ErrStackIsEmpty = errors.New("the stack is empty")
)

type Stack interface {
	Push(value interface{}) error
	Pop() (interface{}, error)
	Top() (interface{}, error)
	IsEmpty() bool
}
