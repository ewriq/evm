package cpu

import (
	"fmt"

	oc "evm/internal/opcode"
)

func (c *CPU) Execute(instruction oc.Instruction, signal ControlSignal) {
	if signal.Halt {
		c.Running = false
		return
	}

	switch instruction.Opcode {

	case oc.LOAD:
		c.Registers.Write(
			byte(instruction.Arg1),
			uint32(instruction.Arg2),
		)

	case oc.ADD:
		a := c.Registers.Read(byte(instruction.Arg2))
		b := c.Registers.Read(byte(instruction.Arg3))

		c.Registers.Write(
			byte(instruction.Arg1),
			ALU(ALU_ADD, a, b),
		)

	case oc.SUB:
		a := c.Registers.Read(byte(instruction.Arg2))
		b := c.Registers.Read(byte(instruction.Arg3))

		c.Registers.Write(
			byte(instruction.Arg1),
			ALU(ALU_SUB, a, b),
		)

	case oc.MUL:
		a := c.Registers.Read(byte(instruction.Arg2))
		b := c.Registers.Read(byte(instruction.Arg3))

		c.Registers.Write(
			byte(instruction.Arg1),
			ALU(ALU_MUL, a, b),
		)

	case oc.DIV:
		a := c.Registers.Read(byte(instruction.Arg2))
		b := c.Registers.Read(byte(instruction.Arg3))

		c.Registers.Write(
			byte(instruction.Arg1),
			ALU(ALU_DIV, a, b),
		)

	case oc.PRINT:
		fmt.Println(
			c.Registers.Read(byte(instruction.Arg1)),
		)

	case oc.JMP:
		c.PC = c.ProgramStart + uint32(instruction.Arg1)*8

	case oc.JZ:
		if c.Registers.Read(byte(instruction.Arg1)) == 0 {
			c.PC = c.ProgramStart + uint32(instruction.Arg2)*8
		}

	case oc.STORE:
		address := uint32(instruction.Arg2)

		c.MemoryInterface.Write(
			address,
			byte(c.Registers.Read(byte(instruction.Arg1))),
		)

	case oc.LOADM:
		address := uint32(instruction.Arg2)

		value := c.MemoryInterface.Read(address)

		c.Registers.Write(
			byte(instruction.Arg1),
			uint32(value),
		)

	case oc.PUSH:
		if c.SP <= c.StackStart {
			panic("stack overflow")
		}

		c.SP--

		c.MemoryInterface.Write(
			c.SP,
			byte(c.Registers.Read(byte(instruction.Arg1))),
		)

	case oc.POP:
		if c.SP >= c.StackEnd {
			panic("stack underflow")
		}

		value := c.MemoryInterface.Read(c.SP)

		c.Registers.Write(
			byte(instruction.Arg1),
			uint32(value),
		)

		c.SP++

	case oc.CALL:
		if c.SP < c.StackStart+8 {
			panic("stack overflow")
		}

		c.SP -= 4

		c.MemoryInterface.Write(c.SP, byte(c.FP))
		c.MemoryInterface.Write(c.SP+1, byte(c.FP>>8))
		c.MemoryInterface.Write(c.SP+2, byte(c.FP>>16))
		c.MemoryInterface.Write(c.SP+3, byte(c.FP>>24))
		returnAddress := c.PC + 8

		c.SP -= 4

		c.MemoryInterface.Write(c.SP, byte(returnAddress))
		c.MemoryInterface.Write(c.SP+1, byte(returnAddress>>8))
		c.MemoryInterface.Write(c.SP+2, byte(returnAddress>>16))
		c.MemoryInterface.Write(c.SP+3, byte(returnAddress>>24))
		c.FP = c.SP
		c.PC = c.ProgramStart + uint32(instruction.Arg1)*8

	case oc.RET:
		if c.FP < c.StackStart ||
			c.FP+8 > c.StackEnd {
			panic("invalid stack frame")
		}

		returnAddress :=
			uint32(c.MemoryInterface.Read(c.FP)) |
				uint32(c.MemoryInterface.Read(c.FP+1))<<8 |
				uint32(c.MemoryInterface.Read(c.FP+2))<<16 |
				uint32(c.MemoryInterface.Read(c.FP+3))<<24

		oldFPAddress := c.FP + 4

		oldFP :=
			uint32(c.MemoryInterface.Read(oldFPAddress)) |
				uint32(c.MemoryInterface.Read(oldFPAddress+1))<<8 |
				uint32(c.MemoryInterface.Read(oldFPAddress+2))<<16 |
				uint32(c.MemoryInterface.Read(oldFPAddress+3))<<24

		c.SP = c.FP + 8
		c.FP = oldFP
		c.PC = returnAddress
	}
}