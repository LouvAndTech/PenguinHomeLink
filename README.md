# PenguinHomeLink

PenguinHomeLink is a tool to log your servers' and computers' information into Home Assistant, making it easier to monitor your devices. If you are anything like me and like to keep an eye on your servers' status, this tool is for you.

## Features

- Compatible with virtually all operating systems and distributions.
- Seamless integration with Home Assistant, including automatic device discovery.
- Ultra-lightweight compared to alternatives.
- Extensive and flexible configuration options.

### Supported Platforms
- Linux (amd64 - arm64) (Various distributions) 
- macOS (amd64 - arm64)
- Windows* (amd64 - arm64)

*Windows should work too, but may not be as easy to set up.*

## How does it work?

PenguinHomeLink leverages the Home Assistant MQTT integration to transmit data to your Home Assistant instance. What makes it particularly interesting and lightweight is its ability to gather information about your devices by executing custom commands that you define.

This design makes PenguinHomeLink highly flexible and adaptable, allowing it to work seamlessly on virtually any system where commands can be used to retrieve device information.

> *For example, you can use `cat /sys/class/thermal/thermal_zone0/temp | awk '{print $1/1000}'` to gather the temperature of the CPU on Debian-based distros. Note the use of `awk` for formatting.*

## Installation

For now, I support packaged installation only for Debian-based OS.
For other Linux distributions, macOS, and Windows, you can use the binary or build from source.

**Requirements:**
- Home Assistant with MQTT integration enabled.
- A HA user with access to the MQTT server.

### Package installation (Debian-based OS)

**Requirements:**
- Debian-based OS (Debian, Ubuntu, etc.)

### Steps:
1. Download the latest package from the [releases page](https://github.com/LouvAndTech/PenguinHomeLink/releases).
2. Install the package using `dpkg` or `apt`:
    ```bash
    sudo dpkg -i penguin-home-link_0.1.0_amd64.deb
    ```
    or
    ```bash
    sudo apt install ./penguin-home-link_0.1.0_amd64.deb
    ```
3. Copy the example configuration to `/etc/penguin-home-link/config.yaml`:
    ```bash
    sudo cp /usr/share/doc/penguin-home-link/examples/config-template.yaml /etc/penguin-home-link/config.yaml
    ```
4. Edit the configuration file to set your MQTT server and sensors following the [configuration section below](#configuration).

5. Start the service:
    ```bash
    sudo systemctl start penguinhomelink
    ```

### Other platforms

For this method, you can use the precompiled binaries for your platform or build from source. If you want to use the precompiled binaries, you can download them from the [releases page]() and move to the section [Using the binary](#using-the-binary).

#### Build from source

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
3. Create a `.env` file in the root directory following the `.env-template` file to select the environment and architecture you want to build for.
4. Build the project:
    ```bash
    make build
    ```
5. The binary will be created in the `dist` directory. 

#### Using the binary
For now, you need to decide how to run the binary yourself. I suggest two options, but you can use any method you prefer:
- You could use a `cron job` with the `@reboot` option, although it is not optimal and not reliable. Refer to [Using a cron job](#using-a-cron-job).
- You can use a `systemd` service.

Once you decide, you can configure the app following the section [Configuration](#configuration).

##### Using a cron job
1. Open the crontab file:
    ```bash
    crontab -e
    ```
2. Add the following line to run the binary at startup:
    ```bash
    @reboot /path/to/PenguinHomeLink/binary /path/to/config.yaml
    ```
3. Save and exit the crontab file.

##### Using a systemd service
1. Create a new 'penguinhomelink.service' file based on the example below:
    ```conf
    [Unit]
    Description=PenguinHomeLink - Self ran
    After=network.target

    [Service]
    ExecStart=/path/to/PenguinHomeLink/binary /path/to/config.yaml
    Restart=on-failure
    User=root

    [Install]
    WantedBy=multi-user.target
    ```

## Configuration

You can follow the example in `config-template.yaml` to create your own configuration file named `config.yaml` placed at the same level as your binary.
Below is a detailed explanation of the configuration file elements. You will also find examples of the configuration file for different hosts in the `configs` directory.

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
|                 | `command`             | The command to execute for retrieving the sensor's data.                           | `"cat /sys/class/thermal/thermal_zone0/temp \| awk '{print $1/1000}'"` |
|                 | `device_class`        | The type of sensor data (e.g., temperature, power, etc.).                          | `"temperature"`                                                       |
|                 | `state_class`         | The state class of the sensor (e.g., measurement).                                 | `"measurement"`                                                       |
|                 | `unit_of_measurement` | The unit in which the sensor data is measured.                                     | `"°C"`                                                                |
|                 | `icon` *(optional)*   | (Optional) The icon to represent the sensor in Home Assistant.                     | `"mdi:cpu-64-bit"`                                                    |

*Note that the examples are tested for a Proxmox instance.*

## Improvements

This is a list of potential improvements for the project. Feel free to suggest more in the issues section:

- Add an automatic GitHub action to build the project for all platforms.
- Improve the installation process using something like a `systemd` service.
- Allow a deeper configuration of the sensors based on Home Assistant's capabilities, and allow more of the elements to be optional.
- Add more configuration examples.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the GPL-2.0 license.
