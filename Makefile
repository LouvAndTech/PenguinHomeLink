# Retrieve the environment variables from the .env file
ifneq (,$(wildcard .env))
	include .env
	export
endif

# Base Exec name 
EXEC = PenguinHomeLink

# Environment setup command 
ENV_CMD = env 
ENV_CMD += GOOS=$(ENV)
ENV_CMD += GOARCH=$(ARCH)

# Create directories
MKDIR_P = mkdir -p

# Build command
GO = go
GO += build
# Build flags
GOFLAGS = 

# Source folder
SRC = ./src

# Output folder
OUT_DIR = ./dist

# Output name
OUTNAME = $(EXEC)_$(ENV)_$(ARCH)
# Output file
OUTFILE = $(OUT_DIR)/$(OUTNAME)

# Build rule 
all: $(OUTFILE)

$(OUTFILE) :
	@echo === "Building $(OUTFILE)" ===
	@$(MKDIR_P) $(OUT_DIR)
	$(ENV_CMD) $(GO) $(GOFLAGS) -o $(OUTFILE) $(SRC)
	@echo "=== Build complete ===\n"

# Clean rule
clean:
	@echo === "Cleaning up" ===
	rm -f $(OUTFILE)
	@echo "=== Clean complete ===\n"

rebuild: clean all

phony: all clean rebuild