package main

import (
	"evm/internal/cpu"
	"evm/internal/loader"
	mem "evm/internal/memory"
)

func main() {
	data, err := loader.LoadBinary("a.bin")
	if err != nil {
		panic(err)
	}

	header := loader.DecodeHeader(data)

	if err := loader.ValidateHeader(header); err != nil {
		panic(err)
	}

	if int(header.CodeSize)+loader.HeaderSize > len(data) {
		panic("binary is smaller than declared code size")
	}

	codeStart := loader.HeaderSize
	codeEnd := codeStart + int(header.CodeSize)

	code := data[codeStart:codeEnd]

	memory := mem.NewMemory(2048)

	c := cpu.CPU{
		Memory:          &memory,
		MemoryInterface: mem.NewMemoryInterface(&memory),
		Debug:           false,
	}

	c.Load(code)
	c.PC = c.ProgramStart + header.EntryPoint*8

	c.Run()
}