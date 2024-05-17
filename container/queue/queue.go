package queue

// Node is a node in a linked queue.
type Node struct {
	next  *Node
	Value interface{}
}

// Queue represents a linked queue.
type Queue struct {
	head, tail *Node
	len        int
}

// Init initializes or clears queue q.
func (q *Queue) Init() *Queue {
	q.head = new(Node)
	q.tail = new(Node)
	q.head = q.tail
	q.len = 0
	return q
}

// New returns an initialized queue.
func New() *Queue { return new(Queue).Init() }

// Len returns the number of nodes of queue.
func (q *Queue) Len() int { return q.len }

// Enqueue adds a new node of at the tail of queue.
func (q *Queue) Enqueue(v interface{}) {
	if q.head == nil {
		q.Init()
	}
	node := &Node{Value: v}
	q.tail.next = node
	q.tail = node
	q.len++
}

// Dequeue returns the head node from queue and removes it.
func (q *Queue) Dequeue() *Node {
	if q.len == 0 {
		return nil
	}
	node := q.head.next
	q.head.next = node.next
	if q.tail == node {
		q.tail = q.head
	}
	q.len--
	return node
}
