package register

func ValidRegister(register byte) bool {
	return register < 16
}