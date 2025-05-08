// Libraries
#include <stdio.h>
#include <stdlib.h>

// includes
#include "./includes/main.h"
#include "./includes/conf_types.h"

// packages
#include "./packages/ConfParser/ConfParser.h"

int main(void){
    // Initialize configuration parser
    char *config_file = "config.cfg"; // Path to your configuration file
    ConfParser_initialize(config_file);

    // Load device information
    DeviceInfo device;
    ConfParser_loadDevice(&device);
    printf("Device Name: %s\n", device.name);

    // Load server information
    ServerInfo server;
    ConfParser_loadServer(&server);
    printf("Server IP: %s\n", server.ip);
    printf("Server Port: %s\n", server.port);
    printf("Server Username: %s\n", server.username);
    printf("Server Password: %s\n", server.password);

    // Load sensor information
    SensorInfo sensors[10]; // Assuming a maximum of 10 sensors
    int sensor_count = 0;
    ConfParser_loadSensor(sensors, &sensor_count);
    printf("Number of Sensors: %d\n", sensor_count);
    for (int i = 0; i < sensor_count; i++) {
        printf("Sensor %d:\n", i + 1);
        printf("  Name: %s\n", sensors[i].name);
        printf("  Command: %s\n", sensors[i].command);
        printf("  Device Class: %s\n", sensors[i].device_class);
        printf("  State Class: %s\n", sensors[i].state_class);
        printf("  Unit of Measurement: %s\n", sensors[i].unit_of_measurement);
    }

    // Clean up
    ConfParser_destroy();

    return 0;
}