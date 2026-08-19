package cpu

import "fmt"

const InstructionSize = 8

func (c *CPU) Run() {
	c.Running = true

	for c.Running {

		if c.Debug {
			fmt.Printf(
				"PC=%d SP=%d FP=%d\n",
				c.PC,
				c.SP,
				c.FP,
			)
		}

		oldPC := c.PC

		instruction := c.Fetch()

		signal := DecodeControl(
			instruction.Opcode,
		)

		c.Execute(
			instruction,
			signal,
		)

		if c.Running && c.PC == oldPC {
			c.PC += InstructionSize
		}
	}
}