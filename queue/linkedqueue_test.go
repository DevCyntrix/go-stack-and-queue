package queue_test

import (
	"testing"

	"github.com/DevCyntrix/go-stack-and-queue/queue"
)

func TestLinkedQueue(t *testing.T) {

	q := queue.NewLinkedQueue()
	err := q.Inque("test")
	if err != nil {
		t.Error(err)
	}
	err = q.Inque("test1")
	if err != nil {
		t.Error(err)
	}
	err = q.Inque("test2")
	if err != nil {
		t.Error(err)
	}
	err = q.Inque("test3")
	if err != nil {
		t.Error(err)
	}

	// 1
	value, err := q.Deque()
	if err != nil {
		t.Error(err)
	}
	if value != "test" {
		t.Errorf("wrong return dequed value (exspected: %s, provided: %s)", "test", value)
	}

	// 2
	value, err = q.Deque()
	if err != nil {
		t.Error(err)
	}
	if value != "test1" {
		t.Errorf("wrong return dequed value (exspected: %s, provided: %s)", "test1", value)
	}

	// 3
	value, err = q.Deque()
	if err != nil {
		t.Error(err)
	}
	if value != "test2" {
		t.Errorf("wrong return dequed value (exspected: %s, provided: %s)", "test2", value)
	}

	// 4
	value, err = q.Deque()
	if err != nil {
		t.Error(err)
	}
	if value != "test3" {
		t.Errorf("wrong return dequed value (exspected: %s, provided: %s)", "test3", value)
	}

	// 5
	value, err = q.Deque()
	if err == nil {
		t.Error(err)
	}
	if value != nil {
		t.Error("wrong return dequed value")
	}
}
