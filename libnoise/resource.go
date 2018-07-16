package main

import (
	"sync"
	"unsafe"
)

var Resources = NewResourceRegistry()

type ResourceRegistry struct {
	holder sync.Map // uintptr -> unsafe.Pointer
}

func NewResourceRegistry() *ResourceRegistry {
	return &ResourceRegistry {
	}
}

func (r *ResourceRegistry) Add(p unsafe.Pointer) {
	r.holder.Store(uintptr(p), p)
}

func (r *ResourceRegistry) Delete(p unsafe.Pointer) {
	r.holder.Delete(uintptr(p))
}

func MarshalPointer(p unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(^uintptr(p))
}

func UnmarshalPointer(p unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(^uintptr(p))
}
