# Network Survey Messaging Go gRPC Server

This simple go application represents an example of a gRPC server that can be used to receive messages from the
[Network Survey Android app](https://github.com/christianrowlands/android-network-survey). The functionality of this
server does not exist other than as a test server for testing the Network Survey Android app, or as an example of how
to leverage the .proto files from the
[Network Survey Messaging Library](https://github.com/christianrowlands/network-survey-messaging) to create a gRPC
server written in Go.

## Prerequisites

1. [Install Go](https://go.dev/doc/install)

2. Download Dependencies
   Before working with this application, you need to download all the dependencies. To accomplish this use the command:

```shell
go get craxiom.com/ns-messaging-go-grpc-server
```

## Running the gRPC server

The application can be compiled and run using the following command

```shell
go run .
```

## Network Survey Connection

Use the "Server Connection" in the
[Network Survey Android App](https://play.google.com/store/apps/details?id=com.craxiom.networksurvey) to connect to
the gRPC server. The server address should be the IP address of the machine running the gRPC server, and the port should
be `2621`.

## Updating the generated go code from the gRPC server

I generated the go code in the `messaging` directory by running the commands found in the
[NS Messaging Repo](https://github.com/christianrowlands/network-survey-messaging#build-and-development-instructions)

From there, I copied the output into the `messaging` directory of this project.

## Contact

* **Christian Rowlands** - [Craxiom](https://github.com/christianrowlands)
