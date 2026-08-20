# EVM Binary Format

The EVM binary format is the executable representation of an EVM program.

The binary consists of a **header** followed by the encoded program code.

```text
┌─────────────────┐
│      Header     │
├─────────────────┤
│       Code      │
└─────────────────┘
```

The header contains information required by the binary loader, while the code contains the encoded EVM instructions.

Each instruction is **8 bytes (64 bits)** long.

The current implementation uses a **16-byte header**.

```text
┌──────────────────────────────┐
│ Header       │ 16 bytes      │
├──────────────────────────────┤
│ Instruction  │ 8 bytes       │
│ Instruction  │ 8 bytes       │
│ Instruction  │ 8 bytes       │
│ ...                          │
└──────────────────────────────┘
```

The binary execution chain is:

```text
.vm → Assembler → .bin → Binary Loader → CPU
```

The loader reads the header, extracts the program code, and provides the encoded instructions to the CPU.
