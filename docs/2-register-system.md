# Register System

The register system provides small, fast storage locations inside the EVM CPU. Registers are used to store values temporarily while instructions are being executed.

The CPU accesses registers through the `RegisterFile` structure.

Example:

```text
LOAD R0 10
```

This stores the value `10` in register `R0`.

```text
R0 = 10
R1 = 20
```

Registers are used by arithmetic and other instructions:

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

The register system provides functions for reading and writing register values, while `ValidRegister()` checks whether a register identifier is valid.
