package cpu

type ALUOperation byte


func ALU(op ALUOperation, a, b uint32) uint32 {
	switch op {
	case ALU_ADD:
		return a + b

	case ALU_SUB:
		return a - b

	case ALU_MUL:
		return a * b

	case ALU_DIV:
		if b == 0 {
			panic("division by zero")
		}

		return a / b

	default:
		panic("invalid ALU operation")
	}
}

//Alu Aritmetik logic unitesi 