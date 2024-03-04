package cpu

import "unsafe"

type Register struct {
	object unsafe.Pointer
	len    uint32
}

func (r *Register) SetValue()