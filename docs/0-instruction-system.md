# Instruction System

Each instruction is **8 bytes long** and consists of an opcode, three 16-bit arguments, and a flags byte.
While running, the CPU processes an instruction through the **fetch → decode → execute** stages.

```text
[ Opcode ][   Arg1   ][   Arg2   ][   Arg3   ][ Flags ]
  1 byte    2 bytes     2 bytes     2 bytes     1 byte
```

Total:

```text
8 bytes = 64 bits
```

Example:

```text
LOAD R0 42
```

Another example:

```text
LOAD  → R0 = 42
PRINT → 42
HALT  → CPU stops
```

The 16-bit arguments allow the instruction system to represent values, registers, memory addresses, and jump targets up to `65535`.
