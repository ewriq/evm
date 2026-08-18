# Assembler

The assembler converts EVM assembly source code (`.vm`) into EVM binary code (`.bin`).

```text
.vm
 ↓
Assembler
 ↓
.bin
```

Example:

```text
LOAD R0 42
PRINT R0
HALT
```

The assembler converts each instruction into its binary representation.

---
