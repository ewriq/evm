package register

func NewRegisterFile() *RegisterFile {
	return &RegisterFile{}
}

func (r *RegisterFile) Read(register byte) uint32 {
	if !ValidRegister(register) {
		panic("invalid register")
	}

	return r.Registers[register]
}

func (r *RegisterFile) Write(register byte, value uint32) {
	if !ValidRegister(register) {
		panic("invalid register")
	}

	r.Registers[register] = value
}

func (r *RegisterFile) Reset() {
	r.Registers = [16]uint32{}
}