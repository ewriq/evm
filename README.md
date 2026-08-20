# EVM
An experimental CPU simulation written in Go. It was designed from scratch primarily to understand how an instruction is processed within a CPU.
### Flow

```text
.vm -> Assembler -> .bin -> Loader -> CPU (Fetch/Decode/Execute)
```

### Architectural Components

| Module | Scope |
| :--- | :--- |
| **Logic** | ALU, Control Unit, Opcode System |
| **Storage** | Memory, Stack, Call Stack, Registers |
| **Flow** | Instruction System, Program Flow, Labels |
| **Tools** | Assembler, Binary Loader, EVM Format |

### Project Layout

```text
evm/
├── app/
│   ├── cmd/
│   │   ├── asm/ 
│   │   └── vm/ 
│   └── internal/
│       ├── cpu/
│       ├── crypto/
│       ├── loader/
│       ├── memory/
│       ├── opcode/
│       └── register/
└── docs/ (details)
```

### Note
For detailed technical matters, you can consult the `docs/` folder. Design decisions and instruction sets are documented there.