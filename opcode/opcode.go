package opcode

var (
	EXIT = 0x00

	// The P-Code Instruction Set (src: https://homepages.cwi.nl/~steven/pascal/book/10pcode.html)
	ABI  = 0x01 // (i)	i	 	Absolute value of integer
	ABR  = 0x02 // (r)	r	 	Absolute value of real
	ADI  = 0x03 // (i,i)	i	 	Adds two integers on the top of the stack and leaves an integer result
	ADR  = 0x04 // (r,r)	r	 	Adds two reals on the top of the stack and leaves a real result
	CHKc = 0x05 // No change	PQ	Checks value is between upper and lower bounds
	CHR  = 0x06 // (i)	c	 	Converts integer to character
	CSP  = 0x07 // Special	Q		Call standard procedure
	CUP  = 0x08 // Special	PQ		Call user procedure
	DECc = 0x09 // (x)	x	Q	Decrement
	DIF  = 0x0A // (s, s)	s	 	Set difference
	DVI  = 0x0B // (i,i)	i	 	Integer division
	DVR  = 0x0C // (r,r)	r	 	Real division
	ENT  = 0x0D // Special		PQ	Enter block
	EOF  = 0x0E // (a)	b	 	Test on end of file
	EQUc = 0x0F // (x,x)	b	Q	Compare on equal
	FJP  = 0x10 // (b)	 	 	False jump
	FLO  = 0x11 // (i,r)	r,r	 	Float next to the top
	FLT  = 0x12 // (i)	r	 	Float top of the stack
	GEQc = 0x13 // (x,x)	b	Q	Compare on greater or equal
	INCc = 0x14 // (x)	x	Q	Increment
	INDc = 0x15 // (a)	x	Q	Indexed fetch
	INN  = 0x16 // (i,s)	b	 	Test set membership
	INT  = 0x17 // (s,s)	s	 	Set intersection
	IOR  = 0x18 // (b,b)	b	 	Boolean inciusive OR
	IXA  = 0x19 // (a,i)	a	Q	Compute indexed address
	LAO  = 0x1A //  	a	Q	Load base level address
	LCA  = 0x1B //  	a	Q	Load address of constant
	LCI  = 0x1C //  	x	PQ	Load constant indirect - assembler generated
	LDA  = 0x1D // 	a	PQ	Load address with level P
	LDCc = 0x1E // 	x	Q	Load constant
	LDOc = 0x1F //  	x	Q	Load contents of base level address
	LEQc = 0x20 // (x,x)	b	Q	Compare on less than or equal
	LESc = 0x21 // (x,x)	b	Q	Compare on less than
	LODc = 0x22 //  	x	PQ	Load contents of address
	MOD  = 0x23 // (i,i)	i	 	Modulo
	MOV  = 0x24 // (a,a)	 	Q	Move
	MPI  = 0x25 // (i,i)	i	 	Integer multiplication
	MPR  = 0x26 // (r,r)	r	 	Real multiplication
	MST  = 0x27 // Special	P	Mark stack
	NEQc = 0x28 // (x,x)	b	Q	Compare on not equal
	NGI  = 0x29 // (i)	i	 	Integer sign inversion
	NGR  = 0x2A // (r)	r	 	Real sign inversion
	NOT  = 0x2B // (b)	b	 	Boolean not
	ODD  = 0x2C // (i)	b	 	Test on odd
	ORDc = 0x2D // (x)	i	 	Convert to integer
	RETc = 0x2E // Special	 	Return from block
	SBI  = 0x2F // (i,i)	i	 	Integer subtraction
	SBR  = 0x30 // (r,r)	r	 	Real subtraction
	SGS  = 0x31 // (i)	s	 	Generate singleton set
	SQI  = 0x32 // (i)	i	 	Squareinteger
	SQR  = 0x33 // (r)	r	 	Square real
	SROc = 0x34 // (x)	 	Q	Store at base level address
	STOc = 0x35 // (a,x)	 	 	Store at base level address
	STP  = 0x36 // No effect             Stop
	STRc = 0x37 // (x)	 	PQ	Store at level P
	TRC  = 0x38 // (r)	i	 	Truncate
	UJC  = 0x39 // No effect             Error in case statement
	UJP  = 0x3A // No effect        Q    Unconditional jump
	UNI  = 0x3B // (s, s)	s	 	Set union
	XJP  = 0x3C // (i)	 	Q	Indexed jump
)

type Opcode struct {
	instruction byte
}

func NewOpcode(instruction byte) *Opcode {
	o := new(Opcode)

	o.instruction = instruction

	return o
}

func (o *Opcode) Value() byte {
	return o.instruction
}
