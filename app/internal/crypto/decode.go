package crypto

import (
	"encoding/binary"

	oc "evm/internal/opcode"
	register "evm/internal/register"
)

const InstructionSize = 8

func Decode(data []byte) oc.Instruction {
	if len(data) < InstructionSize {
		panic("invalid instruction size")
	}

	instruction := oc.Instruction{
		Opcode: oc.Opcode(data[0]),

		Arg1: binary.LittleEndian.Uint16(data[1:3]),
		Arg2: binary.LittleEndian.Uint16(data[3:5]),
		Arg3: binary.LittleEndian.Uint16(data[5:7]),

		Flags: data[7],
	}

	if !instruction.Opcode.Valid() {
		panic("invalid opcode")
	}

	switch instruction.Opcode {

	case oc.LOAD:
		if !register.ValidRegister(byte(instruction.Arg1)) {
			panic("invalid register")
		}

	case oc.ADD, oc.SUB, oc.MUL, oc.DIV:
		if !register.ValidRegister(byte(instruction.Arg1)) ||
			!register.ValidRegister(byte(instruction.Arg2)) ||
			!register.ValidRegister(byte(instruction.Arg3)) {
			panic("invalid register")
		}

	case oc.PRINT, oc.PUSH, oc.POP:
		if !register.ValidRegister(byte(instruction.Arg1)) {
			panic("invalid register")
		}

	case oc.JZ:
		if !register.ValidRegister(byte(instruction.Arg1)) {
			panic("invalid register")
		}

	case oc.STORE, oc.LOADM:
		if !register.ValidRegister(byte(instruction.Arg1)) {
			panic("invalid register")
		}

	case oc.HALT:
	case oc.JMP:
	case oc.CALL:
	case oc.RET:

	default:
		panic("unsupported opcode")
	}

	return instruction
}

