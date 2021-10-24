package core

type CircleQueue struct {
	data []interface{}
	capa int64
	head int64
	rear int64
}

func NewCircleQueue(capa int64) *CircleQueue {
	return &CircleQueue{
		data: make([]interface{}, capa),
		capa: capa + 1,
		head: 0,
		rear: 0,
	}
}

func (q *CircleQueue) Push(v interface{}) int64 {
	rear := (q.rear + 1) % q.capa
	if rear == q.head {
		return -1
	}

	q.data[q.rear] = v
	q.rear = rear
	return rear
}

func (q *CircleQueue) Pop() interface{} {
	if q.head == q.rear {
		return nil
	}

	v := q.data[q.head]
	q.head = (q.head + 1) % q.capa
	return v
}

func (q *CircleQueue) Len() int64 {
	return q.rear - q.head
}

func (q *CircleQueue) Empty() bool {
	return q.head == q.rear
}

func (q *CircleQueue) Range(f func(idx int64, value interface{})) {
	for i := q.head; i < q.rear; i = (i + 1) % q.capa {
		f(i, q.data[i])
	}
}
