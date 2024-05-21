.PHONY: all clean install

# Default target
all: clean build install

# Clean up
clean:
	@echo "Cleaning up..."
	@$(MAKE) -s cleanup

# Build binary
build:
	@echo "Building binary..."
	@go build -o out/

# Install binary
install: build
	@echo "Installing binary..."
	@mkdir -p ~/bin && cp out/grep-in-go ~/bin/grep-in-go
	@echo 'function Ggrip() { ~/bin/grep-in-go $1 $2 $3 }' >> ~/.$(SHELL)rc

# Cleanup
cleanup:
	@echo "Performing cleanup..."
	@rm -rf out/
