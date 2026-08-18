# EVM Binary Format

The EVM binary format is the executable representation of an EVM program.

The current structure contains a header followed by the program code.

```text
┌─────────────────┐
│      Header     │
├─────────────────┤
│       Code      │
└─────────────────┘
```

The header contains information required by the loader, while the code contains the encoded EVM instructions.

The current implementation uses a **14-byte header**.

---
