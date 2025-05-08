// header 
#include "ConfParser.h"

// Libraries
#include <stdio.h>
#include <libconfig.h>

config_t cfg;

void ConfParser_initialize(char *config_file) {
    config_init(&cfg);

    if (!config_read_file(&cfg, config_file)) {
        fprintf(stderr, "Error reading config file %s:%d - %s\n",
                config_error_file(&cfg),
                config_error_line(&cfg),
                config_error_text(&cfg));
        config_destroy(&cfg);
        return;
    }
}

void ConfParser_destroy(void) {
    config_destroy(&cfg);
}

void ConfParser_loadDevice(DeviceInfo *device) {
    const char *device_name;
    if (config_lookup_string(&cfg, "device.name", &device_name)) {
        snprintf(device->name, sizeof(device->name), "%s", device_name);
    }
}

void ConfParser_loadServer(ServerInfo *server) {
    const char *ip, *username, *password , *port;

    if (config_lookup_string(&cfg, "mqtt_server.ip", &ip)) {
        snprintf(server->ip, sizeof(server->ip), "%s", ip);
    }
    if (config_lookup_string(&cfg, "mqtt_server.port", &port)) {
        snprintf(server->port, sizeof(server->port), "%s", port);
    }
    if (config_lookup_string(&cfg, "mqtt_server.username", &username)) {
        snprintf(server->username, sizeof(server->username), "%s", username);
    }
    if (config_lookup_string(&cfg, "mqtt_server.password", &password)) {
        snprintf(server->password, sizeof(server->password), "%s", password);
    }
}

void ConfParser_loadSensor(SensorInfo sensors[], int *sensor_count) {
    config_setting_t *setting;
    setting = config_lookup(&cfg, "sensors");
    if (setting != NULL) {
        int count = config_setting_length(setting);
        *sensor_count = count;

        for (int i = 0; i < count; ++i) {
            config_setting_t *sensor = config_setting_get_elem(setting, i);

            const char *name, *command, *device_class, *state_class, *unit;

            if (config_setting_lookup_string(sensor, "name", &name))
                snprintf(sensors[i].name, sizeof(sensors[i].name), "%s", name);
            if (config_setting_lookup_string(sensor, "command", &command))
                snprintf(sensors[i].command, sizeof(sensors[i].command), "%s", command);
            if (config_setting_lookup_string(sensor, "device_class", &device_class))
                snprintf(sensors[i].device_class, sizeof(sensors[i].device_class), "%s", device_class);
            if (config_setting_lookup_string(sensor, "state_class", &state_class))
                snprintf(sensors[i].state_class, sizeof(sensors[i].state_class), "%s", state_class);
            if (config_setting_lookup_string(sensor, "unit_of_measurement", &unit))
                snprintf(sensors[i].unit_of_measurement, sizeof(sensors[i].unit_of_measurement), "%s", unit);
        }
    }
}