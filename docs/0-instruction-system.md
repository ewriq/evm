
# Instruction System

Each instruction is 4 bytes long and consists of an opcode and 3 arguments on the Go side. Each instruction is 4 bytes long.
While running, the CPU processes an instruction through the fetch -> decode -> execute stages.

```
[ Opcode ][ Arg1 ][ Arg2 ][ Arg3 ]
  1 byte   1 byte   1 byte   1 byte
```

Example:

```
LOAD R0 10
```

Another example:
```
LOAD  → R0 = 42
PRINT → 42
HALT  → CPU stops
```

