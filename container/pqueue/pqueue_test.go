package pqueue

import (
	"fmt"
	"testing"
)

func TestCreatePriorityQueue(t *testing.T) {
	if pq := NewPriorityQueue(10); pq == nil {
		t.Error("NewPriorityQueue failed")
	}
}

func TestPushAndPop(t *testing.T) {
	pq := NewPriorityQueue(10)
	if pq == nil {
		t.Error("NewPriorityQueue failed")
	}

	item1 := &Item{Value: "Item 1", Priority: 3}
	item2 := &Item{Value: "Item 2", Priority: 4}

	pq.Push(item1)
	pq.Push(item2)

	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
}
