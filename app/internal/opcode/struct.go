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
	Arg1   uint16
	Arg2   uint16
	Arg3   uint16
	Flags  byte
}
// LOAD R0 10 