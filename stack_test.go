package stackarray_test

import (
	"testing"

	stackarray "github.com/DevCyntrix/go-stack-array"
)

func TestStackPushPop(t *testing.T) {

	stack := stackarray.New(2)
	err := stack.Push("test")
	if err != nil {
		t.Fail()
		return
	}
	err = stack.Push("test1")
	if err != nil {
		t.Fail()
		return
	}
	err = stack.Push("test2")
	if err == nil {
		t.Fail()
		return
	}
	value, err := stack.Pop()
	if err != nil {
		t.Fail()
		return
	}
	if value != "test1" {
		t.Fail()
	}
	value, err = stack.Pop()
	if err != nil {
		t.Fail()
		return
	}
	if value != "test" {
		t.Fail()
	}

	value, err = stack.Pop()
	if err == nil {
		t.Fail()
		return
	}
	if value != nil {
		t.Fail()
	}
}

func TestStackTop(t *testing.T) {
	stack := stackarray.New(2)
	err := stack.Push("test")
	if err != nil {
		t.Fail()
		return
	}
	value, err := stack.Top()
	if err != nil {
		t.Fail()
	}
	if value != "test" {
		t.Fail()
	}
	value, err = stack.Pop()
	if err != nil {
		t.Fail()
	}
	if value != "test" {
		t.Fail()
	}

	value, err = stack.Top()
	if err == nil {
		t.Fail()
	}
	if value != nil {
		t.Fail()
	}
}