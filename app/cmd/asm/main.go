package main

import (

	crypto "evm/internal/crypto"
	"evm/internal/loader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	text := "a"

	lines, err := loader.File(text + ".vm")
	if err != nil {
		panic(err)
	}

	labels := make(map[string]byte)

	instructionIndex := byte(0)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasSuffix(line, ":") {
			label := strings.TrimSuffix(line, ":")

			if label == "" {
				panic("empty label")
			}

			if _, exists := labels[label]; exists {
				panic("duplicate label: " + label)
			}

			labels[label] = instructionIndex
			continue
		}

		instructionIndex++
	}

	fmt.Println("Labels:", labels)

	var code []byte

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasSuffix(line, ":") {
			continue
		}

		parts := strings.Fields(line)

		if parts[0] == "CALL" && len(parts) == 2 {
			target, exists := labels[parts[1]]
			if !exists {
				panic("unknown label: " + parts[1])
			}

			line = "CALL " + strconv.Itoa(int(target))
		}

		if parts[0] == "JMP" && len(parts) == 2 {
			target, exists := labels[parts[1]]
			if !exists {
				panic("unknown label: " + parts[1])
			}

			line = "JMP " + strconv.Itoa(int(target))
		}

		if parts[0] == "JZ" && len(parts) == 3 {
			target, exists := labels[parts[2]]
			if !exists {
				panic("unknown label: " + parts[2])
			}

			line = "JZ " + parts[1] + " " + strconv.Itoa(int(target))
		}

		instruction, err := loader.Parser(line)
		if err != nil {
			panic(err)
		}

		data := crypto.Encode(instruction)

		code = append(code, data...)
	}

	header := loader.Header{
		Magic:      [4]byte{'E', 'V', 'M', '1'},
		Version:    loader.Version,
		Flags:      0,
		EntryPoint: 0,
		CodeSize:   uint32(len(code)),
	}

	headerData := loader.EncodeHeader(header)

	binary := append(headerData, code...)

	err = loader.SaveBinary(text+".bin", binary)
	if err != nil {
		panic(err)
	}

	fmt.Println("Code size:", len(code))
	fmt.Println("Header size:", len(headerData))
	fmt.Println("Binary size:", len(binary))
}