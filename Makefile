# Default target: Display help message
default: help

# Help message
.PHONY: help
help:
	@echo "Usage:"
	@echo "	make build DIR=<build_directory>	Run test scripts (Specify the directory)."
	@echo "	make test DIR=<test_directory>		Run test scripts (Specify the directory)."
	@echo "	make submit DIR=<submit_directory>	Run test scripts (Specify the directory)."

# build target
.PHONY: build
build:
	@if [ -z "$(DIR)" ]; then \
		echo "Error: Please specify the directory (e.g., make build DIR=<build_directory>)"; \
		exit 1; \
	fi
	@echo "Build directory: $(DIR)"
	@cd $(DIR) && atcoder-tools compile

# Test target
.PHONY: test
test: build
	@if [ -z "$(DIR)" ]; then \
		echo "Error: Please specify the directory (e.g., make test DIR=<test_directory>)"; \
		exit 1; \
	fi
	@echo "Test directory: $(DIR)"
	@cd $(DIR) && atcoder-tools test

# Submit target
.PHONY: submit
submit: build
	@if [ -z "$(DIR)" ]; then \
		echo "Error: Please specify the directory (e.g., make submit DIR=<submit_directory>)"; \
		exit 1; \
	fi
	@echo "Test directory: $(DIR)"
	@cd $(DIR) && atcoder-tools submit -u
