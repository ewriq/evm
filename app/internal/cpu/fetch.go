package cpu

import (
	crypto"evm/internal/crypto"
	oc "evm/internal/opcode"
)

func (c *CPU) Fetch() oc.Instruction {
	if c.PC+4 > c.ProgramStart+c.ProgramSize {
		panic("program counter out of bounds")
	}

	data := c.MemoryInterface.Memory.Read(
		int(c.PC),
		4,
	)

	return crypto.Decode(data)
}