package opcode

var (
	// EXIT terminates the programm.
	EXIT = 0x00

	// LCD puts a value onto a stack
	LDC = 0x01

	// SR stores a value into a register
	SR = 0x02

	// GOTO jump to a label
	GOTO = 0x10

	// JMP jumps to a lable if 0-register value is 1
	JMP = 0x11

	// JMP jumps to a lable if 0-register value is not 1
	JMP_N = 0x12

	// XOR_OP performs an XOR operation against two registers.
	XOR_OP = 0x20

	// ADD_OP performs an ADD operation against two registers.
	ADD_OP = 0x21

	// SUB_OP performs an MINUS operation against two registers.
	SUB_OP = 0x22

	// MUL_OP performs a MULTIPLY operation against two registers.
	MUL_OP = 0x23

	// DIV_OP performs a DIVIDE operation against two registers.
	DIV_OP = 0x24

	// INC_OP increments the given register by one.
	INC_OP = 0x25

	// DEC_OP decrements the given register by one.
	DEC_OP = 0x26

	// AND_OP performs a logical AND operation against two registers.
	AND_OP = 0x27

	// OR_OP performs a logical OR operation against two registers.
	OR_OP = 0x28
)
