package queue

type node struct {
	item int
	next *node
}

// LinkedQueue ...
type LinkedQueue struct {
	first *node
	last  *node
	size  int
}

// IsEmpty ...
func (q *LinkedQueue) IsEmpty() bool {
	return q.size == 0
}

//Enqueue ...
func (q *LinkedQueue) Enqueue(value int) {
	oldLast := q.last
	q.last = &node{item: value, next: nil}

	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}

	q.size++
}

// Dequeue ...
func (q *LinkedQueue) Dequeue() int {
	toReturn := q.first.item
	q.first = q.first.next
	q.size--

	if q.IsEmpty() {
		q.last = nil
	}

	return toReturn
}
