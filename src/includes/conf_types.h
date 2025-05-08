#ifndef CONF_TYPES_H
#define CONF_TYPES_H

#include "main.h"

typedef struct {
    char name[128];
} DeviceInfo;

typedef struct {
    char ip[64];
    char port[16];
    char username[64];
    char password[64];
} ServerInfo;

typedef struct {
    DeviceInfo parent[MAX_SENSORS];
    char name[128];
    char command[256];
    char device_class[64];
    char state_class[64];
    char unit_of_measurement[32];
} SensorInfo;

#endif // CONF_TYPES_H