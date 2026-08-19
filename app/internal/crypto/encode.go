package crypto

import (
	"encoding/binary"

	oc "evm/internal/opcode"
)
func Encode(instruction oc.Instruction) []byte {
	data := make([]byte, InstructionSize)

	data[0] = byte(instruction.Opcode)

	binary.LittleEndian.PutUint16(
		data[1:3],
		instruction.Arg1,
	)

	binary.LittleEndian.PutUint16(
		data[3:5],
		instruction.Arg2,
	)

	binary.LittleEndian.PutUint16(
		data[5:7],
		instruction.Arg3,
	)

	data[7] = instruction.Flags

	return data
}