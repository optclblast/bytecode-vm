package cpu

import (
	"testing"
	"unsafe"
)

func TestCPURun(t *testing.T) {
	c := &CPU{}

	c.stack_p.Add(8)

	c.mem[0] = 0x01

	i := -100

	iB := *((*[8]byte)(unsafe.Pointer(&i)))

	for i, ib := range iB {
		c.mem[i+1] = ib
	}

	c.Run()
}
