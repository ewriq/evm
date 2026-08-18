# Function / Call Stack

The call stack allows the EVM to execute functions using `CALL` and `RET`.

When `CALL` is executed, the CPU stores the return address and current frame information on the stack.

```text
CALL
 ↓
Save return address
 ↓
Create stack frame
 ↓
Jump to function
```

`RET` restores the previous frame and returns to the saved address.

```text
RET
 ↓
Restore FP
 ↓
Restore return address
 ↓
Continue execution
```

---
