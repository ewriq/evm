# Binary Loader

The Binary Loader reads an EVM `.bin` file and prepares it for execution.

```text
.bin
 ↓
Load
 ↓
Read Header
 ↓
Parse / Validate
 ↓
Extract Code
 ↓
Memory
 ↓
CPU
```

The loader is implemented under:

```text
internal/loader/
```

and is responsible for handling the EVM binary before execution.

---
