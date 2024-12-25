package container

// Queue implements a FIFO (First In First Out) queue data structure using a circular buffer.
// Elements are added to the back and removed from the front of the queue.
// This implementation provides better performance than a linked list by:
// 1. Improving cache locality with contiguous memory
// 2. Reducing allocations and pointer chasing
// 3. Using a circular buffer to avoid shifting elements
type Queue[T any] struct {
	buf    []T
	head   int // index of the first element
	tail   int // index of next position to write
	length int // current number of elements
}

// NewQueue creates and returns a new empty queue.
// The type parameter T can be any type.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		buf: make([]T, 16), // Start with reasonable capacity
	}
}

// IsEmpty returns true if the queue contains no elements.
func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int {
	return q.length
}

// Push adds a new element to the back of the queue.
func (q *Queue[T]) Push(data T) {
	if q.length == len(q.buf) {
		// Queue is full, resize
		q.resize()
	}

	q.buf[q.tail] = data
	q.tail = (q.tail + 1) % len(q.buf)
	q.length++
}

// Pop removes and returns the element at the front of the queue.
// Panics if the queue is empty.
func (q *Queue[T]) Pop() T {
	if q.IsEmpty() {
		panic("Queue is empty")
	}

	data := q.buf[q.head]
	var zero T
	q.buf[q.head] = zero // Help GC
	q.head = (q.head + 1) % len(q.buf)
	q.length--

	// Shrink the buffer if it's too large
	if len(q.buf) > 16 && q.length < len(q.buf)/4 {
		q.shrink()
	}

	return data
}

// resize doubles the size of the internal buffer
func (q *Queue[T]) resize() {
	newBuf := make([]T, len(q.buf)*2)

	// Copy elements to the start of the new buffer
	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	q.head = 0
	q.tail = q.length
	q.buf = newBuf
}

// shrink reduces the size of the internal buffer by half
func (q *Queue[T]) shrink() {
	newBuf := make([]T, len(q.buf)/2)

	// Copy elements to the start of the new buffer
	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	q.head = 0
	q.tail = q.length
	q.buf = newBuf
}

// Values returns a slice of all elements in the queue in the order they were added.
func (q *Queue[T]) Values() []T {
	values := make([]T, q.Len())
	for i := range q.Len() {
		values[i] = q.buf[(q.head+i)%len(q.buf)]
	}
	return values
}

// NewQueueFromSlice creates and returns a new queue initialized with the elements from the given slice.
func NewQueueFromSlice[T any](slice []T) *Queue[T] {
	q := NewQueue[T]()
	for _, elem := range slice {
		q.Push(elem)
	}
	return q
}
