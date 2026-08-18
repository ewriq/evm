package cpu

import 	(
	oc "evm/internal/opcode"
	register "evm/internal/register"
)

func Decode(data []byte) oc.Instruction {
	if len(data) < 4 {
		panic("invalid instruction size")
	}

	instruction := oc.Instruction{
		Opcode: oc.Opcode(data[0]),
		Arg1:   data[1],
		Arg2:   data[2],
		Arg3:   data[3],
	}

	if !instruction.Opcode.Valid() {
		panic("invalid opcode")
	}

	switch instruction.Opcode {

	case oc.LOAD:
		if !register.ValidRegister(instruction.Arg1) {
			panic("invalid register")
		}

	case oc.ADD:
		if !register.ValidRegister(instruction.Arg1) ||
			!register.ValidRegister(instruction.Arg2) ||
			!register.ValidRegister(instruction.Arg3) {
			panic("invalid register")
		}

	case oc.SUB:
		if !register.ValidRegister(instruction.Arg1) ||
			!register.ValidRegister(instruction.Arg2) ||
			!register.ValidRegister(instruction.Arg3) {
			panic("invalid register")
		}

	case oc.MUL:
		if !register.ValidRegister(instruction.Arg1) ||
			!register.ValidRegister(instruction.Arg2) ||
			!register.ValidRegister(instruction.Arg3) {
			panic("invalid register")
		}

	case oc.DIV:
		if !register.ValidRegister(instruction.Arg1) ||
			!register.ValidRegister(instruction.Arg2) ||
			!register.ValidRegister(instruction.Arg3) {
			panic("invalid register")
		}

	case oc.PRINT:
		if !register.ValidRegister(instruction.Arg1) {
			panic("invalid register")
		}

	case oc.HALT:
		break

	case oc.JMP:
		break

	case oc.JZ:
		if !register.ValidRegister(instruction.Arg1) {
			panic("invalid register")
		}

	case oc.STORE:
		if !register.ValidRegister(instruction.Arg1) {
			panic("invalid register")
		}

	case oc.LOADM:
		if !register.ValidRegister(instruction.Arg1) {
			panic("invalid register")
		}

	case oc.PUSH:
		if !register.ValidRegister(instruction.Arg1) {
			panic("invalid register")
		}

	case oc.POP:
		if !register.ValidRegister(instruction.Arg1) {
			panic("invalid register")
		}

	case oc.CALL:
		break

	case oc.RET:
		break
	default:
		panic("unsupported opcode")
	}

	return instruction
}

func Encode(instruction oc.Instruction) []byte {
	return []byte{
		byte(instruction.Opcode),
		instruction.Arg1,
		instruction.Arg2,
		instruction.Arg3,
	}
}