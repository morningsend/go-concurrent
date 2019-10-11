package concurrent

import (
	"sync/atomic"
	"unsafe"
)

type stackNode struct {
	next  *stackNode
	value interface{}
	len   int32
}

type Stack struct {
	top *stackNode
	len int32
}

func NewStack() *Stack {
	return &Stack{
		top: nil,
	}
}
func (s *Stack) Push(value interface{}) {
	newNode := &stackNode{value: value, len: 1}

	for {
		t := (*stackNode)(
			atomic.LoadPointer(
				(*unsafe.Pointer)(
					unsafe.Pointer(&s.top))))

		newNode.next = t
		if t != nil {
			newNode.len = t.len + 1
		}
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&s.top)),
			unsafe.Pointer(t),
			unsafe.Pointer(newNode),
		) {
			atomic.AddInt32(&s.len, 1)
			break
		}
	}
}

func (s *Stack) Pop() (value interface{}, ok bool) {
	var t *stackNode
	for {
		t = (*stackNode)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&s.top))))
		if t == nil {
			return nil, false
		}

		next := t.next
		value = t.value

		if ok = atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&s.top)),
			unsafe.Pointer(t),
			unsafe.Pointer(next),
		); ok {
			break
		}
	}

	// garbage collection
	t.next = nil
	return
}

func (s *Stack) Len() int {
	t := (*stackNode)(
		atomic.LoadPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&s.top))))
	if t == nil {
		return 0
	}

	return int(t.len)
}
