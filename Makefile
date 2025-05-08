# Executable name
EXEC_NAME = PenguinHomeLink

# Folder structure
SRC_DIR = ./src
BUILD_DIR = ./build
BIN_DIR = ./bin

# Executable path
EXEC = $(BIN_DIR)/$(EXEC_NAME)

# Commands
RM = rm -f
CC = gcc
MKDIR_P = mkdir -p
MAKE_DEP = make

# Compilation flags
CCFLAGS = -Wall -MMD -MP
CCFLAGS += -O0
CCFLAGS += -g
CCFLAGS += -I/usr/local/include
CCFLAGS += -Ilibs/MQTT-C/include 

# Linker flags
LDFLAGS = -lm
LDFLAGS += -L/usr/local/lib
LDFLAGS += -lconfig
LDFLAGS += libs/MQTT-C/build/libmqttc.a

# Error handeling 
CCFLAGS += -Werror=uninitialized

# Source files
SRC = $(wildcard $(shell find $(SRC_DIR) -type f -regex ".*\.c"))

# Object files
OBJ = $(patsubst $(SRC_DIR)/%, $(BUILD_DIR)/%, $(SRC:.c=.o))
DEP = $(OBJ:.o=.d)

# Compilation rules
all: $(EXEC)

debug: CCFLAGS += -g
debug: rebuild

rebuild: distclean all

# Link the object files to create the executable
$(EXEC): $(OBJ)
	@ echo "\n==== Linker ========="
	@ $(MKDIR_P) $(dir $@)
	$(CC) $^ -o $@ $(LDFLAGS)
	@ echo "==== Linker done ====\n"

# Dependances
-include $(DEP)

# Compile the source files to object files
$(BUILD_DIR)/%.o: $(SRC_DIR)/%.c
	@ echo "\n==== Build Object $< ========="
	@ $(MKDIR_P) $(dir $@)
	$(CC) $(CCFLAGS) -c $< -o $@
	@ echo "==== Build Object done $< ====\n"

run: 
	@ echo "\n==== Run $(EXEC_NAME) ========="
	$(EXEC)
	@ echo "==== Run $(EXEC_NAME) done ====\n"

# Clean the project
clean:
	@ echo "\n==== Cleaning build files ===="
	$(RM) -r $(BUILD_DIR)
	@ echo "==== Build files cleaned =====\n"

distclean: clean
	@ echo "\n==== Cleaning binary files ===="
	$(RM) -r $(BIN_DIR)
	@ echo "==== Binary files cleaned =====\n"

build-dep:
	@ echo "\n==== Building dependencies in $(MAKE_DEP) ===="
	$(MAKE_DEP) -C libs
	@ echo "==== Dependencies built in $(MAKE_DEP) ====\n"

# Phony targets
.PHONY: all clean run distclean rebuild debug build-dep