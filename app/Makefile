PROJECT := wemos-c
BUILD := build

all:
	powershell -ExecutionPolicy Bypass -File scripts\build.ps1

clean:
	powershell -Command "Remove-Item -Recurse -Force build\*"

test:
	powershell -ExecutionPolicy Bypass -File scripts\test.ps1

flash:
	powershell -ExecutionPolicy Bypass -File scripts\flash.ps1

help:
	@echo "make        - Build"
	@echo "make clean  - Clean"
	@echo "make test   - Test Wemos"
	@echo "make flash  - Create ESP8266 image"
