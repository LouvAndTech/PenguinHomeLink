#include "Sender.h"

#include "../../../libs/MQTT-C/include/mqtt.h"

int sockfd;

void Sender_new(char* addr, char* port){
    sockfd = open_nb_socket(addr, port);

    if (sockfd == -1) {
        perror("Failed to open socket: ");
        exit_example(EXIT_FAILURE, sockfd, NULL);
    }
}

void Sender_

/* setup a client */
struct mqtt_client client;
uint8_t sendbuf[2048]; /* sendbuf should be large enough to hold multiple whole mqtt messages */
uint8_t recvbuf[1024]; /* recvbuf should be large enough any whole mqtt message expected to be received */
mqtt_init(&client, sockfd, sendbuf, sizeof(sendbuf), recvbuf, sizeof(recvbuf), publish_callback);
/* Create an anonymous session */
const char* client_id = NULL;
/* Ensure we have a clean session */
uint8_t connect_flags = MQTT_CONNECT_CLEAN_SESSION;
/* Send connection request to the broker. */
mqtt_connect(&client, client_id, NULL, NULL, 0, NULL, NULL, connect_flags, 400);