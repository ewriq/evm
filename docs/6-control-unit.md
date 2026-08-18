
# Control Unit

The Control Unit determines which CPU components are required for an instruction.

The EVM uses `DecodeControl()` to generate control signals.

```text
Opcode
  ↓
DecodeControl()
  ↓
ControlSignal
```

Control signals can activate components such as:

```text
ALU
Register
Memory
Stack
PC
FP
Halt
```

Example:

```text
ADD
 ↓
ALU = true
Register = true
```

---