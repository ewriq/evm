# CPU Fetch / Decode / Execute Cycle

The CPU executes instructions through three main stages:

```text
Fetch
  ↓
Decode
  ↓
Execute
```

**Fetch** reads the instruction from memory using the Program Counter (`PC`).

**Decode** converts the instruction bytes into an `Instruction` and determines the operation.

**Execute** performs the operation using the CPU components.

Example:

```text
PC
 ↓
Fetch instruction
 ↓
Decode LOAD
 ↓
Execute
 ↓
R0 = 42
```

---
