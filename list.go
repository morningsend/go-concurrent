package concurrent

type listNode struct {
	next *listNode
	prev *listNode

	data interface{}
}

type LinkedList struct {
	head *listNode
	tail *listNode
}

func NewLinkedList() *LinkedList {
	sentinel1 := &listNode{}
	sentinel2 := &listNode{}
	sentinel1.next = sentinel2
	sentinel2.prev = sentinel1

	return &LinkedList{
		head: sentinel1,
		tail: sentinel2,
	}
}

func (l *LinkedList) Append(element interface{}) {

}

func (l *LinkedList) Prepend(element interface{}) {

}

func (l *LinkedList) Len() int {
	return 0
}
func (l *LinkedList) NewIterator() *ListIterator {
	return &ListIterator{}
}

type ListIterator struct {
}
