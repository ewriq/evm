# Opcode System

Each instruction has an opcode that tells the CPU which operation to execute. Opcodes are defined as a `byte` and each operation has a unique numeric value.

```text
LOAD   → 0
ADD    → 1
PRINT  → 2
HALT   → 3
...
```

The CPU reads the opcode during the decode stage and determines which operation should be executed.

Example:

```text
LOAD R0 10
```

The assembler converts `LOAD` into its corresponding opcode, and the CPU decodes it:

```text
Opcode → LOAD
       → R0 = 10
```

The `Valid()` function is used to check whether an opcode is supported by the EVM.
