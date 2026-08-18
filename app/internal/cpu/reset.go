package cpu

func (c *CPU) Reset() {
	c.Registers.Reset()

	c.PC = c.ProgramStart

	c.StackStart = StackStartAddress
	c.StackEnd = StackEndAddress

	c.SP = c.StackEnd
	c.FP = c.SP

	c.Running = true
}