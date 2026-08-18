package cpu

func (c *CPU) Load(data []byte) {
	if len(data) > CodeEnd-CodeStart {
		panic("program too large")
	}

	c.ProgramStart = CodeStart
	c.ProgramSize = uint32(len(data))

	c.Memory.Write(
		CodeStart,
		data,
	)

	c.Reset()
}