package loader

import (
	"fmt"
	"strconv"
	"strings"

	oc "evm/internal/opcode"
)

func Parser(line string) (oc.Instruction, error) {
	parts := strings.Fields(line)

	if len(parts) == 0 {
		return oc.Instruction{}, fmt.Errorf("empty instruction")
	}

	switch parts[0] {

	case "LOAD":
		if len(parts) != 3 {
			return oc.Instruction{}, fmt.Errorf("LOAD requires 2 arguments")
		}

		register, err := RegisterParser(parts[1])
		if err != nil {
			return oc.Instruction{}, err
		}

		value, err := strconv.Atoi(parts[2])
		if err != nil {
			return oc.Instruction{}, fmt.Errorf("invalid value: %s", parts[2])
		}

		if value < 0 || value > 65535 {
			return oc.Instruction{}, fmt.Errorf("value out of range: %d", value)
		}

		return oc.Instruction{
			Opcode: oc.LOAD,
			Arg1:   uint16(register),
			Arg2:   uint16(value),
		}, nil

	case "SUB":
		if len(parts) != 4 {
			return oc.Instruction{}, fmt.Errorf("SUB requires 3 arguments")
		}

		arg1, err := RegisterParser(parts[1])
		if err != nil {
			return oc.Instruction{}, err
		}

		arg2, err := RegisterParser(parts[2])
		if err != nil {
			return oc.Instruction{}, err
		}

		arg3, err := RegisterParser(parts[3])
		if err != nil {
			return oc.Instruction{}, err
		}

		return oc.Instruction{
			Opcode: oc.SUB,
			Arg1:   uint16(arg1),
			Arg2:   uint16(arg2),
			Arg3:   uint16(arg3),
		}, nil

	case "ADD":
		if len(parts) != 4 {
			return oc.Instruction{}, fmt.Errorf("ADD requires 3 arguments")
		}

		arg1, err := RegisterParser(parts[1])
		if err != nil {
			return oc.Instruction{}, err
		}

		arg2, err := RegisterParser(parts[2])
		if err != nil {
			return oc.Instruction{}, err
		}

		arg3, err := RegisterParser(parts[3])
		if err != nil {
			return oc.Instruction{}, err
		}

		return oc.Instruction{
			Opcode: oc.ADD,
			Arg1:   uint16(arg1),
			Arg2:   uint16(arg2),
			Arg3:   uint16(arg3),
		}, nil

	case "PRINT":
		if len(parts) != 2 {
			return oc.Instruction{}, fmt.Errorf("PRINT requires 1 argument")
		}

		register, err := RegisterParser(parts[1])
		if err != nil {
			return oc.Instruction{}, err
		}

		return oc.Instruction{
			Opcode: oc.PRINT,
			Arg1:   uint16(register),
		}, nil

	case "HALT":
		if len(parts) != 1 {
			return oc.Instruction{}, fmt.Errorf("HALT requires no arguments")
		}

		return oc.Instruction{
			Opcode: oc.HALT,
		}, nil

	case "MUL":
		if len(parts) != 4 {
			return oc.Instruction{}, fmt.Errorf("MUL requires 3 arguments")
		}

		arg1, err := RegisterParser(parts[1])
		if err != nil {
			return oc.Instruction{}, err
		}

		arg2, err := RegisterParser(parts[2])
		if err != nil {
			return oc.Instruction{}, err
		}

		arg3, err := RegisterParser(parts[3])
		if err != nil {
			return oc.Instruction{}, err
		}

		return oc.Instruction{
			Opcode: oc.MUL,
			Arg1:   uint16(arg1),
			Arg2:   uint16(arg2),
			Arg3:   uint16(arg3),
		}, nil

	case "DIV":
		if len(parts) != 4 {
			return oc.Instruction{}, fmt.Errorf("DIV requires 3 arguments")
		}

		arg1, err := RegisterParser(parts[1])
		if err != nil {
			return oc.Instruction{}, err
		}

		arg2, err := RegisterParser(parts[2])
		if err != nil {
			return oc.Instruction{}, err
		}

		arg3, err := RegisterParser(parts[3])
		if err != nil {
			return oc.Instruction{}, err
		}

		return oc.Instruction{
			Opcode: oc.DIV,
			Arg1:   uint16(arg1),
			Arg2:   uint16(arg2),
			Arg3:   uint16(arg3),
		}, nil

	case "JMP":
		if len(parts) != 2 {
			return oc.Instruction{}, fmt.Errorf("JMP requires 1 argument")
		}

		target, err := strconv.Atoi(parts[1])
		if err != nil {
			return oc.Instruction{}, fmt.Errorf(
				"invalid JMP target: %s",
				parts[1],
			)
		}

		if target < 0 || target > 65535 {
			return oc.Instruction{}, fmt.Errorf(
				"JMP target out of range: %d",
				target,
			)
		}

		return oc.Instruction{
			Opcode: oc.JMP,
			Arg1:   uint16(target),
		}, nil

	case "JZ":
		if len(parts) != 3 {
			return oc.Instruction{}, fmt.Errorf("JZ requires 2 arguments")
		}

		register, err := RegisterParser(parts[1])
		if err != nil {
			return oc.Instruction{}, err
		}

		target, err := strconv.Atoi(parts[2])
		if err != nil {
			return oc.Instruction{}, fmt.Errorf(
				"invalid JZ target: %s",
				parts[2],
			)
		}

		if target < 0 || target > 65535 {
			return oc.Instruction{}, fmt.Errorf(
				"JZ target out of range: %d",
				target,
			)
		}

		return oc.Instruction{
			Opcode: oc.JZ,
			Arg1:   uint16(register),
			Arg2:   uint16(target),
		}, nil

	case "STORE":
		if len(parts) != 3 {
			return oc.Instruction{}, fmt.Errorf("STORE requires 2 arguments")
		}

		register, err := RegisterParser(parts[1])
		if err != nil {
			return oc.Instruction{}, err
		}

		address, err := strconv.Atoi(parts[2])
		if err != nil {
			return oc.Instruction{}, fmt.Errorf(
				"invalid memory address: %s",
				parts[2],
			)
		}

		if address < 0 || address > 65535 {
			return oc.Instruction{}, fmt.Errorf(
				"memory address out of range: %d",
				address,
			)
		}

		return oc.Instruction{
			Opcode: oc.STORE,
			Arg1:   uint16(register),
			Arg2:   uint16(address),
		}, nil

	case "LOADM":
		if len(parts) != 3 {
			return oc.Instruction{}, fmt.Errorf("LOADM requires 2 arguments")
		}

		register, err := RegisterParser(parts[1])
		if err != nil {
			return oc.Instruction{}, err
		}

		address, err := strconv.Atoi(parts[2])
		if err != nil {
			return oc.Instruction{}, fmt.Errorf(
				"invalid memory address: %s",
				parts[2],
			)
		}

		if address < 0 || address > 65535 {
			return oc.Instruction{}, fmt.Errorf(
				"memory address out of range: %d",
				address,
			)
		}

		return oc.Instruction{
			Opcode: oc.LOADM,
			Arg1:   uint16(register),
			Arg2:   uint16(address),
		}, nil

	case "PUSH":
		if len(parts) != 2 {
			return oc.Instruction{}, fmt.Errorf("PUSH requires 1 argument")
		}

		register, err := RegisterParser(parts[1])
		if err != nil {
			return oc.Instruction{}, err
		}

		return oc.Instruction{
			Opcode: oc.PUSH,
			Arg1:   uint16(register),
		}, nil

	case "CALL":
		if len(parts) != 2 {
			return oc.Instruction{}, fmt.Errorf("CALL requires 1 argument")
		}

		target, err := strconv.Atoi(parts[1])
		if err != nil {
			return oc.Instruction{}, fmt.Errorf(
				"invalid CALL target: %s",
				parts[1],
			)
		}

		if target < 0 || target > 65535 {
			return oc.Instruction{}, fmt.Errorf(
				"CALL target out of range: %d",
				target,
			)
		}

		return oc.Instruction{
			Opcode: oc.CALL,
			Arg1:   uint16(target),
		}, nil

	case "RET":
		if len(parts) != 1 {
			return oc.Instruction{}, fmt.Errorf("RET requires no arguments")
		}

		return oc.Instruction{
			Opcode: oc.RET,
		}, nil

	case "POP":
		if len(parts) != 2 {
			return oc.Instruction{}, fmt.Errorf("POP requires 1 argument")
		}

		register, err := RegisterParser(parts[1])
		if err != nil {
			return oc.Instruction{}, err
		}

		return oc.Instruction{
			Opcode: oc.POP,
			Arg1:   uint16(register),
		}, nil

	default:
		return oc.Instruction{}, fmt.Errorf(
			"unknown instruction: %s",
			parts[0],
		)
	}
} 