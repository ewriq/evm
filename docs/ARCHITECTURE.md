# EVM Architecture Summary

Bu belge, kapsamlı teknik araştırma raporunun uygulama özeti ve repository eşlemesidir.

## Kararlar (ADR Özet)

| Karar | Seçim |
|-------|-------|
| VM modeli | Register-based (8 GPR + stack frame) |
| Bellek | Sabit harita, 32 KB sanal adres uzayı |
| Heap | Arena allocator, GC yok |
| Storage | FAT32, SD→RAM tam yükleme |
| Sandbox | Validator + bounds check + instruction budget |
| Dil | C (VM core platform-agnostic) |

## Bellek Haritası

```
0x0000  CODE   (8 KB)
0x2000  CONST  (4 KB)
0x3000  DATA   (4 KB)
0x4000  HEAP   (12 KB)
0x7000  STACK  (4 KB)
0x8000  limit
```

## Bytecode Header (28 byte)

| Offset | Alan | Boyut |
|--------|------|-------|
| 0x00 | MAGIC "EVM1" | 4 |
| 0x04 | VERSION | 2 |
| 0x06 | FLAGS | 2 |
| 0x08 | CODE_SIZE | 4 |
| 0x0C | DATA_SIZE | 4 |
| 0x10 | CONST_SIZE | 4 |
| 0x14 | ENTRY_POINT | 4 |
| 0x18 | CHECKSUM (CRC32) | 4 |

## ISA (Opcode Tablosu)

Tam liste: `firmware/vm/include/evm_opcodes.h`

## Syscall ABI

- Argümanlar: R0–R3
- Dönüş: R0
- Hata: FLAGS.C set, R0 = hata kodu

## Roadmap

| Faz | Durum |
|-----|-------|
| 1–4 VM core + bytecode | ✅ Host'ta implemente |
| 5 SD loader | ✅ Boot katmanı (host FS) |
| 7 Validator | ✅ |
| 8 Crash dump | ✅ Temel |
| 6 Debugger UI | 🔲 MVP sonrası |
| 9 Assembler | ✅ vm-asm |
| 0 ESP32 donanım | 🔲 platform_esp32.c |

## SD Dizin Yapısı (Hedef)

```
/programs/   → .vm dosyaları
/config/     → config.toml
/crash/      → crash dump
/logs/       → runtime log
```
