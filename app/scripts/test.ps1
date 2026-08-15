param(
    [Parameter(Mandatory=$false)]
    [string]$Port = "COM3"
)

$ErrorActionPreference = "Stop"

Write-Host ""
Write-Host "===================================="
Write-Host " Wemos D1 Mini Test"
Write-Host "===================================="
Write-Host ""

Write-Host "Port: $Port"
Write-Host ""

esptool `
    --chip esp8266 `
    --port $Port `
    chip-id

Write-Host ""

esptool `
    --chip esp8266 `
    --port $Port `
    flash-id

Write-Host ""
Write-Host "ESP8266 test tamamlandi."
