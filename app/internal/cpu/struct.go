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

    DataStart = 0
    DataEnd   = 1024

    CodeStart = 1024
    CodeEnd   = 2048
	HeapStart = 1280
	HeapEnd   = 1536

	FreeStart = 1536
	FreeEnd   = 1792

	StackStartAddress = 1792
	StackEndAddress   = 2048
	  MemorySize = 4096


    StackStart = 2048
    StackEnd   = 4096
   
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

