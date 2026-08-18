package opcode

func (o Opcode) Valid() bool {
	switch o {
	case LOAD,
		ADD,
		PRINT,
		HALT,
		SUB,
		MUL,
		DIV,
		JMP,
		JZ,
		STORE,
		LOADM,
		PUSH,
		POP,
		CALL,
		RET:
		return true

	default:
		return false
	}
}