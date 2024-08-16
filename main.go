package main

import (
	pb "craxiom.com/ns-messaging-go-grpc-server/messaging"
	"craxiom.com/ns-messaging-go-grpc-server/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":2621")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register the services
	pb.RegisterConnectionHandshakeServer(grpcServer, &server.ConnectionHandshakeServer{})
	pb.RegisterWirelessSurveyServer(grpcServer, &server.WirelessSurveyServer{})
	pb.RegisterDeviceStatusServer(grpcServer, &server.DeviceStatusServer{})

	log.Println("gRPC server is running on port 2621...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
