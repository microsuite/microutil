package queue

import (
	"testing"
)

func checkQueueLen(t *testing.T, q *Queue, expected int) {
	if q.Len() != expected {
		t.Errorf("Expected queue length of %d, got %d", expected, q.Len())
	}
}

func checkValue(t *testing.T, n *Node, v interface{}) {
	if n == nil || n.Value == nil || v == nil {
		t.Errorf("Expected node value of %v, got nil", v)
	}

	if n.Value != v {
		t.Errorf("Expected node value of %v, got %v", v, n.Value)
	}
}

func TestQueue(t *testing.T) {
	q := New()
	checkQueueLen(t, q, 0)

	// Add a value
	q.Enqueue("hello")
	checkQueueLen(t, q, 1)

	q.Enqueue("world")
	checkQueueLen(t, q, 2)

	q.Enqueue(2024)
	checkQueueLen(t, q, 3)

	n1 := q.Dequeue()
	checkValue(t, n1, "hello")
	checkQueueLen(t, q, 2)

	n2 := q.Dequeue()
	checkValue(t, n2, "world")
	checkQueueLen(t, q, 1)

	n3 := q.Dequeue()
	checkValue(t, n3, 2024)
	checkQueueLen(t, q, 0)

	// Dequeue from empty queue
	n4 := q.Dequeue()
	if n4 != nil {
		t.Errorf("Expected nil, got %v", n4)
	}
}
