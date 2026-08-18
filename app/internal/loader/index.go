package loader

import (
	"evm/internal/register"

	"fmt"
	"strconv"
	"strings"
)

func RegisterParser(value string) (byte, error) {
	if !strings.HasPrefix(value, "R") {
		return 0, fmt.Errorf("invalid register: %s", value)
	}

	number, err := strconv.Atoi(value[1:])
	if err != nil {
		return 0, fmt.Errorf("invalid register: %s", value)
	}

	if number < 0 || !register.ValidRegister(byte(number)) {
		return 0, fmt.Errorf("unknown register: %s", value)
	}

	return byte(number), nil
}