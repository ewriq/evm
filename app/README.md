# Wemos D1 Mini - Pure C

## Target

ESP8266EX
Xtensa LX106

## Structure

src/
    main.c
    gpio.c
    startup.S

include/
    gpio.h
    esp8266.h

linker/
    eagle.app.v6.ld

scripts/
    build.ps1
    test.ps1
    flash.ps1

build/
    generated files

Makefile
README.md

## Test Wemos

USB ile Wemos'u bilgisayara bagla.

COM portunu bul:

    COM3

Test:

    .\scripts\test.ps1 -Port COM3

## Build

    .\scripts\build.ps1

veya:

    make

## Image

    .\scripts\flash.ps1 -Port COM3

## Make

    make
    make test
    make clean
    make flash

## Pipeline

C
 â†“
Xtensa GCC
 â†“
Object
 â†“
Linker
 â†“
ELF
 â†“
ESP8266 image
 â†“
esptool
 â†“
USB
 â†“
Wemos D1 Mini
