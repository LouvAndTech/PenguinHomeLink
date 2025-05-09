# PenguinHomeLink

PenguinHomeLink is a tool to log your servers' and computers' information into Home Assistant, making it easier to monitor your devices.
If you are anything like me and like to keep an eye on your servers staus, this tool is for you.

## Features

- Integrates on all distributions.
- Seamless integration with Home Assistant.
- Automatic discovery of devices.
- Extensive configuration options.

### Supported Platforms
- Linux (amd64 - arm64) (Various distributions) 
- MacOS (amd64 - arm64)
- Windows* (amd64 - arm64)

**Windows should works too, but may not be as easy to setup*

## How does it work?

PanguininHomeLink uses the Home Assistant MQTT integration to send data to your Home Assistant instance. It collects information about your devices by running command and gathering the output before sending it to Home Assistant, where you can use it in your dashboards.

> *For example, your can use `cat /sys/class/thermal/thermal_zone0/temp | awk '{print $1/1000}'` to gather the temperture of the CPU on debian based distros, note the use of `awk` for formating.*

## Installation

*I want to improve this process in the future when I get the time.*

### Build from source

**Requirements:**
- Golang 1.18 or later

### Steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/LouvAndTech/PenguinHomeLink.git
    ```
2. Navigate to the project directory:
    ```bash
    cd PenguinHomeLink
    ```
3. Create a `.env` file in the root directory following the `.env-template` file to select the environement and architecture you want to build for.
4. Build the project:
    ```bash
    make build
    ```
5. The binary will be created in the `dist` directory. 
6. For now you need to decide how to run the binary yourself. I can advise tower `cron jobs` with the `@reboot` option, althought it is not optimal.
    
    1. Open the crontab file:
        ```bash
        crontab -e
        ```
    2. Add the following line to run the binary at startup:
        ```bash
        @reboot /path/to/PenguinHomeLink/dist/PenguinHomeLink
        ```
    3. Save and exit the crontab file.
 

## Configuration

You can follow the example in `config-termplate.yaml` to create your own configuration file named `config.yaml` placed at the same level as your binary.
Below is a detailed explanation of the configuration file elements:

*Note that the `sensor` category is a list.*

| **Section**     | **Key**               | **Description**                                                                    | **Example**                                                           |
| --------------- | --------------------- | ---------------------------------------------------------------------------------- | --------------------------------------------------------------------- |
| **software**    | `refresh_period_s`    | The interval in seconds at which the data is refreshed and sent to Home Assistant. | `30`                                                                  |
| **device**      | `name`                | The name of the device being monitored.                                            | `"MyLinuxDevice"`                                                     |
|                 | `manufacturer`        | The manufacturer of the device.                                                    | `"DeviceManufacturer"`                                                |
|                 | `model`               | The model of the device.                                                           | `"DeviceModel"`                                                       |
|                 | `serial_number`       | The serial number of the device.                                                   | `"xxxxxxxxxxxxxx..."`                                                 |
| **mqtt_server** | `ip`                  | The IP address of the MQTT server.                                                 | `"192.168.1.50"`                                                      |
|                 | `port`                | The port number of the MQTT server.                                                | `"1883"`                                                              |
|                 | `username`            | The username for authenticating with the MQTT server.                              | `"myuser"`                                                            |
|                 | `password`            | The password for authenticating with the MQTT server.                              | `"mypassword"`                                                        |
| **sensors**[*]  | `name`                | The name of the sensor.                                                            | `"CPU Temperature"`                                                   |
|                 | `command`             | The command to execute for retrieving the sensor's data.                           | `"cat /sys/class/thermal/thermal_zone0/temp | awk '{print $1/1000}'"` |
|                 | `device_class`        | The type of sensor data (e.g., temperature, power, etc.).                          | `"temperature"`                                                       |
|                 | `state_class`         | The state class of the sensor (e.g., measurement).                                 | `"measurement"`                                                       |
|                 | `unit_of_measurement` | The unit in which the sensor data is measured.                                     | `"°C"`                                                                |
|                 | `icon` *(optional)*   | (Optional) The icon to represent the sensor in Home Assistant.                     | `"mdi:cpu-64-bit"`                                                    |

*Note that the example are tested for a proxmox instance.*

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the GPL-2.0 license.
