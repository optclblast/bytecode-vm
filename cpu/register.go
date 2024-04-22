package cpu

import (
	"reflect"
	"sync"
	"unsafe"
)

type Register struct {
	m sync.RWMutex

	value Value
}

type Value struct {
	object     unsafe.Pointer
	objectType reflect.Type
	len        uintptr
}

func (r *Register) SetValue(obj any, size uintptr) {
	r.m.Lock()
	defer r.m.Unlock()

	value := new(Value)

	value.object = unsafe.Pointer(&obj)
	value.objectType = reflect.TypeOf(obj)
	value.len = size
}

func (r *Register) Value() Value {
	r.m.RLock()
	defer r.m.RUnlock()

	return r.value
}
