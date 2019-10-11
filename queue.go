package concurrent

import (
	"sync/atomic"
	"unsafe"
)

type queueNode struct {
	next  *queueNode
	value interface{}
}

type Queue struct {
	tail  *queueNode
	head  *queueNode
	empty *queueNode
	len   int32
}

func NewQueue() *Queue {
	emptyNode := &queueNode{}
	return &Queue{
		empty: emptyNode,
		tail:  emptyNode,
		head:  emptyNode,
	}
}

func (q *Queue) Enqueue(value interface{}) {
	node := &queueNode{
		value: value,
		next:  nil,
	}
	var (
		tail *queueNode
		next *queueNode
	)

	for {
		tail = (*queueNode)(
			atomic.LoadPointer((*unsafe.Pointer)(
				unsafe.Pointer(
					&q.tail))))
		next = (*queueNode)(
			atomic.LoadPointer((*unsafe.Pointer)(
				unsafe.Pointer(
					&tail.next))))
		if tail == (*queueNode)(
			atomic.LoadPointer((*unsafe.Pointer)(
				unsafe.Pointer(
					&q.tail)))) {
			if next == nil {
				if atomic.CompareAndSwapPointer(
					(*unsafe.Pointer)(unsafe.Pointer(&tail.next)),
					unsafe.Pointer(nil),
					unsafe.Pointer(node),
				) {
					break
				}
			} else {
				atomic.CompareAndSwapPointer(
					(*unsafe.Pointer)(unsafe.Pointer(&q.tail)),
					unsafe.Pointer(tail),
					unsafe.Pointer(next),
				)
			}
		}
	}

	atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&q.tail)),
		unsafe.Pointer(tail),
		unsafe.Pointer(node),
	)
	atomic.AddInt32(&q.len, 1)
}

func (q *Queue) Dequeue() (interface{}, bool) {
	var (
		head  *queueNode
		tail  *queueNode
		next  *queueNode
		value interface{}
	)

	for {
		head = (*queueNode)(
			atomic.LoadPointer((*unsafe.Pointer)(
				unsafe.Pointer(
					&q.head))))
		tail = (*queueNode)(
			atomic.LoadPointer((*unsafe.Pointer)(
				unsafe.Pointer(
					&q.tail))))
		next = (*queueNode)(
			atomic.LoadPointer((*unsafe.Pointer)(
				unsafe.Pointer(
					&head.next))))

		if head == (*queueNode)(
			atomic.LoadPointer((*unsafe.Pointer)(
				unsafe.Pointer(
					&q.head)))) {
			if head == tail {
				if next == nil {
					return nil, false
				}
				atomic.CompareAndSwapPointer(
					(*unsafe.Pointer)(unsafe.Pointer(&q.tail)),
					unsafe.Pointer(tail),
					unsafe.Pointer(next),
				)
			} else {
				value = next.value
				if atomic.CompareAndSwapPointer(
					(*unsafe.Pointer)(unsafe.Pointer(&q.head)),
					unsafe.Pointer(head),
					unsafe.Pointer(next),
				) {
					break
				}
			}
		}
	}
	atomic.AddInt32(&q.len, -1)
	return value, true
}

func (q *Queue) Len() int {
	len := atomic.LoadInt32(&q.len)
	return int(len)
}
