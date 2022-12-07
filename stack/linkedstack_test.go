package stack_test

import (
	"testing"

	stack "github.com/DevCyntrix/go-stack-and-queue/stack"
)

func TestLinkedtackPushPop(t *testing.T) {

	stack := stack.NewLinkedStack()
	err := stack.Push("test")
	if err != nil {
		t.Error("Cannot push items on the stack")
	}
	err = stack.Push("test1")
	if err != nil {
		t.Error("Cannot push items on the stack")
	}
	err = stack.Push("test2")
	if err != nil {
		t.Error("Cannot push items on the stack")
	}
	value, err := stack.Pop()
	if err != nil {
		t.Fail()
	}
	if value != "test2" {
		t.Errorf("Got wrong value (exspected: %s, provided: %s)", "test2", value)
	}
	value, err = stack.Pop()
	if err != nil {
		t.Fail()
	}
	if value != "test1" {
		t.Errorf("Got wrong value (exspected: %s, provided: %s)", "test1", value)
	}
	value, err = stack.Pop()
	if err != nil {
		t.Fail()
	}
	if value != "test" {
		t.Errorf("Got wrong value (exspected: %s, provided: %s)", "test", value)
	}

	value, err = stack.Pop()
	if err == nil {
		t.Fail()
	}
	if value != nil {
		t.Fail()
	}
}

func TestLinkedStackTop(t *testing.T) {
	stack := stack.NewLinkedStack()
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
