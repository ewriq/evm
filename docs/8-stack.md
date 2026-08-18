# Stack

The stack is a memory region used for temporary data and function execution.

The CPU uses:

```text
SP → Stack Pointer
FP → Frame Pointer
```

The main stack instructions are:

```text
PUSH
POP
```

Example:

```text
LOAD R0 42
PUSH R0
POP R1
```

Result:

```text
R1 = 42
```

---
