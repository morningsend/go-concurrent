package concurrent

import (
	"sync/atomic"
	"unsafe"
)

type dequeNode struct {
	value   interface{}
	prev    taggedPointer
	next    taggedPointer
	deleted bool
}

const (
	dequeNodeDefaultTag uint = 0
	dequeNodeDeletedTag uint = 1
)

func newDequeNode(value interface{}) *dequeNode {
	return &dequeNode{value: value}
}

type Deque struct {
	head *dequeNode
	tail *dequeNode
}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) AddFirst(value interface{}) {
	node := newDequeNode(value)

	prev := (*dequeNode)(
		atomic.LoadPointer((*unsafe.Pointer)(
			unsafe.Pointer(
				&d.head))))

	nextTagged := (tagged(atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&prev.next)),
	)))

	for {
		if prev.next != nextTagged.WithTag(dequeNodeDefaultTag) {
			nextTagged = (tagged(atomic.LoadPointer(
				(*unsafe.Pointer)(unsafe.Pointer(&prev.next)),
			)))
			continue
		}

		node.prev = tagged(unsafe.Pointer(prev))
		node.next = nextTagged.WithTag(dequeNodeDefaultTag)

		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&prev.next)),
			nextTagged.WithTag(dequeNodeDefaultTag).Pointer(),
			makeTagged(unsafe.Pointer(node), dequeNodeDefaultTag).Pointer(),
		) {
			break
		}

	}

}

func (d *Deque) AddLast(value interface{}) {

}

func (d *Deque) RemoveFirst() (value interface{}, ok bool) {
	return
}

func (d *Deque) RemoveLast() (value interface{}, ok bool) {
	return
}

func (d *Deque) addCommon(node, next *dequeNode) {
	for {
		link1 := tagged(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&next.prev))))
		if link1.GetTag() == dequeNodeDeletedTag ||
			node.next != makeTagged(unsafe.Pointer(next), dequeNodeDefaultTag) {
			break
		}

		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&next.prev)),
			link1.Pointer(),
			makeTagged(unsafe.Pointer(node), dequeNodeDefaultTag).Pointer(),
		) {

			if node.prev.GetTag() == dequeNodeDeletedTag {

			}
			break
		}
	}
}
