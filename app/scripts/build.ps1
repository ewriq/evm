$ErrorActionPreference = "Stop"

$PROJECT = Split-Path -Parent $PSScriptRoot
$BUILD = Join-Path $PROJECT "build"

Write-Host ""
Write-Host "===================================="
Write-Host " Wemos C Build"
Write-Host "===================================="
Write-Host ""

New-Item `
    -ItemType Directory `
    -Path $BUILD `
    -Force | Out-Null

Write-Host "[1/4] main.c"

xtensa-lx106-elf-gcc `
    -mlongcalls `
    -ffunction-sections `
    -fdata-sections `
    -Os `
    -I"$PROJECT\include" `
    -c "$PROJECT\src\main.c" `
    -o "$BUILD\main.o"

Write-Host "[2/4] gpio.c"

xtensa-lx106-elf-gcc `
    -mlongcalls `
    -ffunction-sections `
    -fdata-sections `
    -Os `
    -I"$PROJECT\include" `
    -c "$PROJECT\src\gpio.c" `
    -o "$BUILD\gpio.o"

Write-Host "[3/4] startup.S"

xtensa-lx106-elf-gcc `
    -mlongcalls `
    -I"$PROJECT\include" `
    -c "$PROJECT\src\startup.S" `
    -o "$BUILD\startup.o"

Write-Host "[4/4] ELF"

xtensa-lx106-elf-gcc `
    -nostdlib `
    "$BUILD\main.o" `
    "$BUILD\gpio.o" `
    "$BUILD\startup.o" `
    -o "$BUILD\firmware.elf"

Write-Host ""
Write-Host "===================================="
Write-Host " BUILD TAMAMLANDI"
Write-Host "===================================="
Write-Host ""

Write-Host "Firmware:"
Write-Host "$BUILD\firmware.elf"
