package cpu

import (
	"bufio"
	"os"

	"github.com/optclblast/bytecode-vm/stack"
)

// VM's cpu object
type CPU struct {
	regs   [16]*Register // VM registers
	stack  stack.Stack   // Our stack
	mem    [0xFFFF]byte  // RAM
	STDIN  *bufio.Reader
	STDOUT *bufio.Writer
}

func NewCPU() *CPU {
	return &CPU{
		stack:  stack.Allocstack(1024),
		STDIN:  bufio.NewReader(os.Stdin),
		STDOUT: bufio.NewWriter(os.Stdout),
	}
}
