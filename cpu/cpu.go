package cpu

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/optclblast/bytecode-vm/stack"
)

// VM's cpu object
type CPU struct {
	iPointer int
	regs     [16]*Register // VM registers
	stack    stack.Stack   // Our stack
	mem      [0xFFFF]byte  // RAM
	STDIN    *bufio.Reader
	STDOUT   *bufio.Writer
}

func NewCPU() *CPU {
	var regs [16]*Register

	for i := range regs {
		regs[i] = new(Register)
	}

	return &CPU{
		regs:   regs,
		stack:  stack.AllocStack(1024),
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
	if len(data) > 0xFFFF {
		return fmt.Errorf("error programm is to large")
	}

	c.Reset()

	for i := 0; i < len(data); i++ {
		c.mem[i] = data[i]
	}

	return nil
}

func (c *CPU) Reset() {
	c.mem = [0xFFFF]byte{}

	var regs [16]*Register

	for i := range regs {
		regs[i] = new(Register)
	}

	c.stack = stack.AllocStack(1024)
}

func (c *CPU) Run() error {
	run := true
	for run {
		if c.iPointer >= 0xffff {
			return fmt.Errorf("reading beyond RAM")
		}

		op := opcode.NewOpcode(c.mem[c.iPointer])
	}

	return nil
}
