package opcode

type Opcode byte

const (
	LOAD Opcode = iota
	ADD
	PRINT
	HALT
	SUB
	MUL
	DIV
	JMP
	JZ
	STORE
	LOADM
	PUSH
	POP
	CALL
	RET
)

type Instruction struct {
	Opcode Opcode
	Arg1   byte
	Arg2   byte
	Arg3   byte
}

// LOAD R0 10 