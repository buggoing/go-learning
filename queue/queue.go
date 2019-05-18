package queue

type Item struct {
	item interface{}
	pre  *Item
}

type Queue struct {
	head  *Item
	tail  *Item
	depth uint64
}

func New() *Queue {
	q := new(Queue)
	q.depth = 0
	return q
}

func (q *Queue) Enqueue(item interface{}) {
	if q.depth == 0 {
		qItem := Item{item, nil}
		q.head = &qItem
		q.tail = &qItem
		q.depth++
		return
	}
	qItem := Item{item, nil}
	q.head.pre = &qItem
	q.head = &qItem
	q.depth++

}

func (q *Queue) Dequeue() (item interface{}) {
	if q.depth > 0 {
		qItem := q.tail
		q.tail = q.tail.pre
		q.depth--
		return qItem.item
	}

	return nil
}
