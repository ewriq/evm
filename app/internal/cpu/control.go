package cpu
import (	 
oc "evm/internal/opcode"
)



func DecodeControl(opcode oc.Opcode) ControlSignal {
	switch opcode {

	case oc.LOAD:
		return ControlSignal{
			Register: true,
		}

	case oc.ADD, oc.SUB, oc.MUL, oc.DIV:
		return ControlSignal{
			ALU:      true,
			Register: true,
		}

	case oc.PRINT:
		return ControlSignal{
			Register: true,
		}

	case oc.JMP, oc.JZ:
		return ControlSignal{
			Register: opcode == oc.JZ,
			PC:       true,
		}

	case oc.STORE, oc.LOADM:
		return ControlSignal{
			Memory:   true,
			Register: true,
		}

	case oc.PUSH, oc.POP:
		return ControlSignal{
		Stack:    true,
		Register: true,
		}

	case oc.CALL, oc.RET:
		return ControlSignal{
			Stack: true,
			PC:    true,
			FP:    true,
		}

	case oc.HALT:
		return ControlSignal{
			Halt: true,
		}

	default:
		panic("invalid opcode")
	}
}