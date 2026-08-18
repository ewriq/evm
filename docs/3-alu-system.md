# ALU

The ALU (Arithmetic Logic Unit) is the part of the EVM CPU responsible for performing arithmetic operations on register values.

Currently, the ALU supports:

```text
ADD
SUB
MUL
DIV
```

The CPU reads the required values from registers, sends them to the ALU, and stores the result in a destination register.

Example:

```text
LOAD R0 10
LOAD R1 20
ADD R2 R0 R1
```

Result:

```text
R2 = R0 + R1
R2 = 30
```

The basic flow is:

```text
Registers
    ↓
   ALU
    ↓
Result Register
```
