.PHONY: all build test clean examples run-add run-fib

BUILD_DIR := build

all: build

build:
	cmake -B $(BUILD_DIR) -DCMAKE_BUILD_TYPE=Debug
	cmake --build $(BUILD_DIR)

test: build
	ctest --test-dir $(BUILD_DIR) --output-on-failure
	$(BUILD_DIR)/vm-check $(BUILD_DIR)/examples/add.vm
	$(BUILD_DIR)/vm-run $(BUILD_DIR)/examples/add.vm

examples: build

run-add: build
	$(BUILD_DIR)/vm-run $(BUILD_DIR)/examples/add.vm

run-fib: build
	$(BUILD_DIR)/vm-run $(BUILD_DIR)/examples/fib.vm

clean:
	rm -rf $(BUILD_DIR)

help:
	@echo "make        - Configure and build"
	@echo "make test   - Run unit tests and smoke tests"
	@echo "make run-add - Run add.vm on host emulator"
	@echo "make clean  - Remove build directory"
