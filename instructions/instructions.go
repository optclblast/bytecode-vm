package instructions

import (
	"fmt"
	"os"
	"strconv"
	"unsafe"

	"github.com/optclblast/bytecode-vm/stack"
)

type Instruction struct {
	Opcode byte
	Fn     func(st stack.Stack, args []string)
}

var InstructionsSet map[string]*Instruction = map[string]*Instruction{
	// ldc loads a value into a stack
	"ldc": &Instruction{
		Opcode: 0x12,
		Fn: func(st stack.Stack, args []string) {
			intVal, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}

			//fmt.Println(*(*int32)(unsafe.Pointer(&intVal))) //dbg
			st.Push(unsafe.Pointer(&intVal))
		},
	},
	// iadd add value to a value from the stack
	"iadd": &Instruction{
		Opcode: 0xa9,
		Fn: func(st stack.Stack, args []string) {
			intVal, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}

			val := (*int32)(st.Pop())
			newVal := *val + int32(intVal)

			//fmt.Println(newVal) //dbg

			st.Push(unsafe.Pointer(&newVal))
		},
	},
	// out prints stack value into stdout
	"out": &Instruction{
		Opcode: 0x1f,
		Fn: func(st stack.Stack, args []string) {
			value := *(*int32)(st.Pop())
			fmt.Fprint(os.Stdout, value, "\n")
		},
	},
	// divide the stack value onto argument value
	"idiv": &Instruction{
		Opcode: 0x2f,
		Fn: func(st stack.Stack, args []string) {
			divisor, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			value := *(*int32)(st.Pop())
			res := value / int32(divisor)

			st.Push(unsafe.Pointer(&res))
		},
	},
}
