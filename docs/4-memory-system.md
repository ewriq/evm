# Memory System

The memory system provides storage for the EVM CPU. It allows the CPU to read and write program data during execution.

The memory is accessed through the `Memory` and `MemoryInterface` structures.

```text
CPU
 ↓
MemoryInterface
 ↓
Memory
```

The memory is divided into different regions:

```text
Code
Data
Heap
Free
Stack
```

Example:

```text
LOAD R0 42
STORE R0 512
```

This stores the value of `R0` into memory address `512`.

---
