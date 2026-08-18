package cpu

type Memory struct {
	Data []byte
}

func NewMemory(size int) Memory {
	return Memory{
		Data: make([]byte, size),
	}
}

func (m *Memory) Write(address int, data []byte) {
	if address < 0 || address+len(data) > len(m.Data) {
		panic("memory write out of bounds")
	}

	copy(m.Data[address:], data)
}
func (m *Memory) Read(address int, size int) []byte {
	if address < 0 || size < 0 || address+size > len(m.Data) {
		panic("memory read out of bounds")
	}

	return m.Data[address : address+size]
}