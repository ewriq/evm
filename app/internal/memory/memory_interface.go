package cpu

type MemoryInterface struct {
	Memory *Memory
}

func NewMemoryInterface(memory *Memory) *MemoryInterface {
	return &MemoryInterface{
		Memory: memory,
	}
}

func (m *MemoryInterface) Read(address uint32) byte {
	if address >= uint32(len(m.Memory.Data)) {
		panic("memory access violation")
	}

	return m.Memory.Data[address]
}

func (m *MemoryInterface) Write(address uint32, value byte) {
	if address >= uint32(len(m.Memory.Data)) {
		panic("memory access violation")
	}

	m.Memory.Data[address] = value
}