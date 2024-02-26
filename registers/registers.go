package registers

import "unsafe"

type Register struct {
	Value unsafe.Pointer
}

var Registers map[byte]Register = map[byte]Register{
	// registers
	0x01: {},
	0x02: {},
	0x03: {},
}
