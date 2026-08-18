package loader

import (
	"fmt"
	"os"
)

func SaveBinary(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func LoadBinary(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("binary is empty")
	}

	return data, nil
}