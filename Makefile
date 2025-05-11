# Retrieve the environment variables from the .env file
ifneq (,$(wildcard .env))
    include .env
endif

# Base Exec name 
EXEC = penguinhomelink

# Environment setup command 
ENV_BUILD = env 
ENV_BUILD += GOOS=$(TARGET_ENV)
ENV_BUILD += GOARCH=$(TARGET_ARCH)

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
BIN_DIR = ./bin

# Output name
EXEC_NAME = $(EXEC)_$(TARGET_ENV)_$(TARGET_ARCH)
# Output file
EXEC_PATH = $(BIN_DIR)/$(EXEC_NAME)

# === Build rule ===
all: $(EXEC_PATH)

$(EXEC_PATH) :
	@echo === "Building $(EXEC_PATH)" ===
	@$(MKDIR_P) $(BIN_DIR)
	$(ENV_BUILD) $(GO) $(GOFLAGS) -o $(EXEC_PATH) $(SRC)
	@echo "=== Build complete ===\n"

# Clean rule
bin-clean:
	@echo === "Cleaning up" ===
	rm -f $(EXEC_PATH)
	@echo "=== Clean complete ===\n"

rebuild: bin-clean all

phony: all bin-clean rebuild

# ===== Packaging rule =====

PACKAGE_OUT_DIR = ./dist
BUILD_DIR = ./build

# Deb
DEB_PACKAGE_NAME = penguinhomelink
PACKAGE_TEMPLATE_DIR = ./packaging/debian
PACKAGE_BUILD_DIR = $(BUILD_DIR)/penguinhomelink_$(VERSION)_$(TARGET_ENV)_$(TARGET_ARCH)

DEB_BIN = /usr/bin/$(EXEC)
DEB_CONTROL = /DEBIAN/control
DEB_CONF = /etc/penguinhomelink/config.yaml
DEB_SERVICE = /lib/systemd/system/penguinhomelink.service

ENV_FILE = .env

deb: $(EXEC_PATH)
	@echo === "Building Debian package" ===
# Build the structure
	@$(MKDIR_P) $(dir $(PACKAGE_BUILD_DIR)/$(DEB_CONTROL))
	@$(MKDIR_P) $(dir $(PACKAGE_BUILD_DIR)/$(DEB_SERVICE))
	@$(MKDIR_P) $(dir $(PACKAGE_BUILD_DIR)/$(DEB_BIN))
	@$(MKDIR_P) $(dir $(PACKAGE_BUILD_DIR)/$(DEB_CONF))
# Copy the files 
	cp $(EXEC_PATH) $(PACKAGE_BUILD_DIR)/$(DEB_BIN)
	cp $(PACKAGE_TEMPLATE_DIR)/$(DEB_CONF) $(PACKAGE_BUILD_DIR)/$(DEB_CONF)
	source ./$(ENV_FILE) && \
		export DEB_BIN=$(DEB_BIN) && \
		export DEB_CONF=$(DEB_CONF) && \
		$(INCLUDE_ENV) envsubst < $(PACKAGE_TEMPLATE_DIR)/$(DEB_CONTROL).in > $(PACKAGE_BUILD_DIR)/$(DEB_CONTROL)
		$(INCLUDE_ENV) envsubst < $(PACKAGE_TEMPLATE_DIR)/$(DEB_SERVICE).in > $(PACKAGE_BUILD_DIR)/$(DEB_SERVICE)

	$(MKDIR_P) $(PACKAGE_OUT_DIR)
	dpkg-deb --build $(PACKAGE_BUILD_DIR) $(PACKAGE_OUT_DIR)/$(DEB_PACKAGE_NAME)_$(VERSION)_$(TARGET_ENV)_$(TARGET_ARCH).deb
	@echo "=== Debian package created at $(PACKAGE_OUT_DIR)/$(DEB_PACKAGE_NAME)_$(VERSION)_$(TARGET_ENV)_$(TARGET_ARCH).deb ===\n"


package-clean:
	@echo === "Cleaning up package build directory" ===
	rm -rf $(PACKAGE_BUILD_DIR)
	@echo "=== Package build directory cleaned ===\n"

phony: deb package-clean


full-clean: 
	@echo === "Cleaning up all build artifacts" ===
	rm -rf $(BIN_DIR)
	rm -rf $(BUILD_DIR)
	rm -rf $(PACKAGE_OUT_DIR)
	@echo "=== All build artifacts cleaned ===\n"
phony: full-clean