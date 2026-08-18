package cpu

import (
	mem "evm/internal/memory"
	"evm/internal/register"
)

type CPU struct {
	Registers       register.RegisterFile
	PC              uint32
	SP              uint32
	FP              uint32

	Memory          *mem.Memory
	MemoryInterface *mem.MemoryInterface

	ProgramStart uint32
	ProgramSize  uint32

	StackStart uint32
	StackEnd   uint32

	Running bool
	Yielded  bool
	Debug   bool

	SyscallHandler func(id byte) error
}

const (
	CodeStart = 0
	CodeEnd   = 512

	DataStart = 512
	DataEnd   = 640

	HeapStart = 640
	HeapEnd   = 768

	FreeStart = 768
	FreeEnd   = 896

	StackStartAddress = 896
	StackEndAddress   = 1024
)

type ControlSignal struct {
	ALU      bool
	Register bool
	Memory   bool
	Stack    bool
	PC       bool
	FP       bool
	Halt     bool
}


const (
	ALU_ADD ALUOperation = iota
	ALU_SUB
	ALU_MUL
	ALU_DIV
)