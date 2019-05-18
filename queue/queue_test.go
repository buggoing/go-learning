package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New()
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < 5; i++ {
		if i != q.Dequeue() {
			t.Error("TestQueue Failed...", i)
		}
	}
}
