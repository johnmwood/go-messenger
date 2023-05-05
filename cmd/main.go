package main

import (
	"context"
	"fmt"
	"log"

	"github.com/johnmwood/go-messenger/proto"

	"google.golang.org/grpc"
)

const (
	PORT = 8080
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := proto.NewMetricsClient(conn)

	// Call the MyMethod method on the server
	response, err := client.ReportMetrics(context.Background(), &proto.MetricRequest{Message: "world"})
	if err != nil {
		log.Fatalf("Failed to call ReportMetrics: %v", err)
	}

	// Print the response from the server
	fmt.Println(response.Message)
}
