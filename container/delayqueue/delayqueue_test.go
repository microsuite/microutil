package delayqueue

import (
	"testing"
	"time"

	"github.com/microsuite/microutil/container/pqueue"
)

func TimeToS(t time.Time) int64 {
	return t.UnixNano() / int64(time.Second)
}

func TestDelayQueue(t *testing.T) {
	quit := make(chan struct{})

	dq := NewDelayQueue(10)

	go func() {
		dq.Poll(quit, func() int64 {
			return TimeToS(time.Now().UTC())
		})
	}()

	item1 := &pqueue.Item{Value: "Item 1", Priority: 3}
	item2 := &pqueue.Item{Value: "Item 2", Priority: 4}

	dq.Add(item1, 3)
	dq.Add(item2, 10)

}
