param(
    [Parameter(Mandatory=$false)]
    [string]$Port = "COM3"
)

$ErrorActionPreference = "Stop"

$PROJECT = Split-Path -Parent $PSScriptRoot
$BUILD = Join-Path $PROJECT "build"

$ELF = Join-Path $BUILD "firmware.elf"

if (!(Test-Path $ELF)) {

    Write-Host ""
    Write-Host "Firmware bulunamadi."
    Write-Host ""
    Write-Host "Once:"
    Write-Host ".\scripts\build.ps1"
    Write-Host ""

    exit 1
}

Write-Host ""
Write-Host "===================================="
Write-Host " ESP8266 IMAGE"
Write-Host "===================================="
Write-Host ""

esptool `
    --chip esp8266 `
    elf2image `
    $ELF

Write-Host ""
Write-Host "Olusan dosyalar:"
Get-ChildItem "$BUILD\*.bin"

Write-Host ""
Write-Host "Port: $Port"
Write-Host ""

Write-Host "IMAGE olusturuldu."
Write-Host "Flash adresleri SDK/image yapisina gore"
Write-Host "belirlenecektir."
