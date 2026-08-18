# Program Flow

Program flow determines which instruction the CPU executes next. The main component responsible for this is the Program Counter (`PC`).

Normally:

```text
PC → Fetch → Execute → PC + 4
```

Jump instructions can change the normal flow.

```text
JMP
 ↓
PC = target
```

Conditional jumps use a register value:

```text
JZ R0 target
```

If `R0` is zero, the CPU jumps to the target address.

---
