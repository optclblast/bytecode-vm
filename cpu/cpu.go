package cpu

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"unsafe"

	"github.com/optclblast/drain-machine/opcode"
)

const (
	R_MP = iota
	R_NP
	R_EP
)

// VM's cpu object
type CPU struct {
	pc_r    atomic.Int64 // pc register
	stack_p atomic.Int64 // stack pointer
	heap_p  atomic.Int64 // heap pointer
	cons_p  atomic.Int64 // constants area pointer
	mark_p  atomic.Int64 // mark pointer

	pp int64 // where the stack lives

	regs [3]*Register // VM registers [ mp, np, ep]
	mem  [0xFFFF]byte // RAM

	STDIN  *bufio.Reader
	STDOUT *bufio.Writer
}

func NewCPU() *CPU {
	var regs [3]*Register

	for i := range regs {
		regs[i] = new(Register)
	}

	return &CPU{
		regs: regs,

		STDIN:  bufio.NewReader(os.Stdin),
		STDOUT: bufio.NewWriter(os.Stdout),
	}
}

func (c *CPU) LoadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error open file %s. %w", path, err)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error read from file %s. %w", path, err)
	}

	return c.LoadProgramm(bytes)
}

func (c *CPU) LoadProgramm(data []byte) error {
	if len(data) > 0xFFFF/3*2 { // the program requires more than 2/3 of all awailable memory, so there will be no space for stack
		return fmt.Errorf("error program is to large")
	}

	c.Reset()

	c.pp = int64(len(data))

	for i := 0; i < len(data); i++ {
		c.mem[i] = data[i]
	}

	return nil
}

// Reset the CPU state
func (c *CPU) Reset() {
	c.mem = [0xFFFF]byte{}

	var regs [16]*Register

	for i := range regs {
		regs[i] = new(Register)
	}

	c.pc_r = atomic.Int64{}
	c.heap_p = atomic.Int64{}
	c.stack_p = atomic.Int64{}
	c.cons_p = atomic.Int64{}
	c.pp = 0
}

func (c *CPU) Run() error {
	run := true
	for run {
		if c.pc_r.Load() >= 0xFFFF {
			return fmt.Errorf("reading beyond RAM")
		}

		op := opcode.NewOpcode(c.mem[c.pc_r.Load()])
		c.pc_r.Add(1)

		switch int(op.Value()) {
		case opcode.ABI:
			p := c.stack_p.Load()

			raw := c.mem[p-7 : p+1]

			value := byteArrayToInt(raw)

			if value < 0 {
				value *= -1
			}

			bytes := [8]byte(intToByteArray(value))

			for i, b := range bytes {
				c.mem[p+int64(i)] = b
			}
		}
	}

	return nil
}

func intToByteArray(num int64) []byte {
	size := int(unsafe.Sizeof(num))
	arr := make([]byte, size)

	for i := 0; i < size; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}
	return arr
}

func byteArrayToInt(arr []byte) int64 {
	val := int64(0)
	size := len(arr)

	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}

	return val
}
