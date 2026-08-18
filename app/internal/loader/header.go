package loader

import (
	"encoding/binary"
	"fmt"
)

const (
	Magic   = "EVM1"
	Version = byte(1)

	HeaderSize = 14
)

type Header struct {
	Magic      [4]byte
	Version    byte
	Flags      byte
	EntryPoint uint32
	CodeSize   uint32
}

func EncodeHeader(h Header) []byte {
	data := make([]byte, HeaderSize)

	copy(data[0:4], h.Magic[:])

	data[4] = h.Version
	data[5] = h.Flags

	binary.LittleEndian.PutUint32(data[6:10], h.EntryPoint)
	binary.LittleEndian.PutUint32(data[10:14], h.CodeSize)

	return data
}

func DecodeHeader(data []byte) Header {
	if len(data) < HeaderSize {
		panic("invalid header size")
	}

	var magic [4]byte
	copy(magic[:], data[0:4])

	return Header{
		Magic:      magic,
		Version:    data[4],
		Flags:      data[5],
		EntryPoint: binary.LittleEndian.Uint32(data[6:10]),
		CodeSize:   binary.LittleEndian.Uint32(data[10:14]),
	}
}

func ValidateHeader(h Header) error {
	if string(h.Magic[:]) != Magic {
		return fmt.Errorf("invalid magic")
	}

	if h.Version != Version {
		return fmt.Errorf(
			"unsupported version: %d",
			h.Version,
		)
	}

	if h.CodeSize == 0 {
		return fmt.Errorf("code size is zero")
	}

	if h.CodeSize%4 != 0 {
		return fmt.Errorf(
			"invalid code size: %d",
			h.CodeSize,
		)
	}

	return nil
}