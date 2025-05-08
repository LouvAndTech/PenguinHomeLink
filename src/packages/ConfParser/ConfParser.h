#ifndef CONFPARSER_H
#define CONFPARSER_H

#include "../../includes/conf_types.h"

void ConfParser_initialize(char *config_file);
void ConfParser_destroy(void);

void ConfParser_loadDevice(DeviceInfo *device);
void ConfParser_loadServer(ServerInfo *server);
void ConfParser_loadSensor(SensorInfo sensors[], int *sensor_count);

#endif // CONFPARSER_H