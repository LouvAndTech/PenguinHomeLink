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
DEB_TEMPLATE_CONF = /etc/penguinhomelink/config-template.yaml
DEB_CONF = /etc/penguinhomelink/config.yaml
DEB_SERVICE = /lib/systemd/system/penguinhomelink.service

ENV_FILE = .env

deb: $(PACKAGE_OUT_DIR)/$(DEB_PACKAGE_NAME)_$(VERSION)_$(TARGET_ENV)_$(TARGET_ARCH).deb

$(PACKAGE_OUT_DIR)/$(DEB_PACKAGE_NAME)_$(VERSION)_$(TARGET_ENV)_$(TARGET_ARCH).deb: $(PACKAGE_BUILD_DIR) $(EXEC_PATH)
	@echo === "Building Debian package" ===
	@$(MKDIR_P) $(PACKAGE_OUT_DIR)
	dpkg-deb --build $(PACKAGE_BUILD_DIR) $@
	@echo "=== Debian package created at $@ ===\n"

$(PACKAGE_BUILD_DIR): $(PACKAGE_BUILD_DIR)/$(DEB_CONTROL) $(PACKAGE_BUILD_DIR)/$(DEB_SERVICE) $(PACKAGE_BUILD_DIR)/$(DEB_BIN) $(PACKAGE_BUILD_DIR)/$(DEB_TEMPLATE_CONF)

$(PACKAGE_BUILD_DIR)/$(DEB_CONTROL): $(PACKAGE_TEMPLATE_DIR)/$(DEB_CONTROL).in
	@echo === "Creating control file" ===
	@$(MKDIR_P) $(dir $@)
	@bash -c '\
		source ./$(ENV_FILE); \
		envsubst < $< > $@; \
	'
	@echo "=== Control file created at $@ ===\n"

$(PACKAGE_BUILD_DIR)/$(DEB_SERVICE): $(PACKAGE_TEMPLATE_DIR)/$(DEB_SERVICE).in
	@echo === "Creating service file" ===
	@$(MKDIR_P) $(dir $@)
	@bash -c '\
		source ./$(ENV_FILE); \
		export DEB_BIN=$(DEB_BIN); \
		export DEB_CONF=$(DEB_CONF); \
		envsubst < $< > $@; \
	'
	@echo "=== Service file created at $@ ===\n"

$(PACKAGE_BUILD_DIR)/$(DEB_BIN): $(EXEC_PATH)
	@echo === "Copying binary to package directory" ===
	@$(MKDIR_P) $(dir $@)
	cp $< $@
	@echo "=== Binary copied to $@ ===\n"

$(PACKAGE_BUILD_DIR)/$(DEB_TEMPLATE_CONF): $(PACKAGE_TEMPLATE_DIR)/$(DEB_TEMPLATE_CONF)
	@echo === "Copying config file to package directory" ===
	@$(MKDIR_P) $(dir $@)
	cp $< $@
	@echo "=== Config file copied to $@ ===\n"

package-clean:
	@echo === "Cleaning up package build directory" ===
	rm -rf $(PACKAGE_BUILD_DIR)
	rm -rf $(PACKAGE_OUT_DIR)
	@echo "=== Package build directory cleaned ===\n"

phony: deb package-clean


full-clean: 
	@echo === "Cleaning up all build artifacts" ===
	rm -rf $(BIN_DIR)
	rm -rf $(BUILD_DIR)
	rm -rf $(PACKAGE_OUT_DIR)
	@echo "=== All build artifacts cleaned ===\n"
phony: full-clean