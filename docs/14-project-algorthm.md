# `.vm → .bin → VM` Chain

This is the complete execution pipeline of the current EVM.

```text
        .vm
         ↓
    Assembler
         ↓
        .bin
         ↓
      Loader
         ↓
      Memory
         ↓
        CPU
         ↓
 Fetch → Decode → Execute
         ↓
      Program
```

Example:

```text
LOAD R0 42
PRINT R0
HALT
```

becomes:

```text
.vm
 ↓
Assembler
 ↓
EVM1 .bin
 ↓
Loader
 ↓
CPU Memory
 ↓
Fetch
 ↓
Decode
 ↓
Execute
 ↓
42
```
