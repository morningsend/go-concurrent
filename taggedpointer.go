package concurrent

import (
	"unsafe"
)

type taggedPointer uint

func (p taggedPointer) GetTag() uint {
	return uint(p) & 0x7
}

func (p *taggedPointer) SetTag(tag uint) {
	*p = taggedPointer((uint(*p) & (^(uint(0x7)))) | (tag & 0xF))
}

func (p taggedPointer) WithTag(tag uint) taggedPointer {
	return taggedPointer((uint(p) & (^(uint(0x7)))) | (tag & 0xF))
}
func (p taggedPointer) GetPointer() unsafe.Pointer {
	return (unsafe.Pointer)(uintptr(uint(p) & (^(uint(0x7)))))
}

func makeTagged(p unsafe.Pointer, tag uint) taggedPointer {
	tg := taggedPointer(uintptr(p))
	tg.SetTag(tag)
	return tg
}

func tagged(p unsafe.Pointer) taggedPointer {
	return taggedPointer(uintptr(p))
}

func (p taggedPointer) Pointer() unsafe.Pointer {
	return (unsafe.Pointer)((uintptr)(uint(p)))
}
