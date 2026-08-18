# Label System

Labels provide names for instruction addresses, allowing programs to use symbolic jump targets instead of manually calculating addresses.

Example:

```text
start:
    LOAD R0 42
    JMP start
```

The assembler keeps track of labels and resolves them to their corresponding instruction addresses.

```text
start
 ↓
Instruction address
 ↓
JMP target
```

---
