package concurrent

import (
	"testing"
	"unsafe"
)

func TestTaggedPointer(t *testing.T) {
	var x int64 = 123
	var tag uint = 3

	taggedPointer := makeTagged(unsafe.Pointer(&x), tag)

	if taggedPointer.GetPointer() != unsafe.Pointer(&x) {
		t.Errorf("pointer is wrong, got %v expect %v", taggedPointer.GetPointer(), unsafe.Pointer(&x))
	}

	if taggedPointer.GetTag() != tag {
		t.Error("tag is wrong")
	}

	tag = 1
	taggedPointer.SetTag(tag)
	if taggedPointer.GetPointer() != unsafe.Pointer(&x) {
		t.Errorf("pointer is wrong, got %v expect %v", taggedPointer.GetPointer(), unsafe.Pointer(&x))
	}

	if taggedPointer.GetTag() != tag {
		t.Error("tag is wrong")
	}
}
