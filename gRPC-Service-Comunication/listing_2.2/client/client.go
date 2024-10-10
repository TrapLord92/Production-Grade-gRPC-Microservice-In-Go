package main

import (
	"context"
	"log"

	payment "listing_2.2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Create a new gRPC client connection using NewClient
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials())) // Insecure credentials for non-TLS
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	paymentClient := payment.NewPaymentServiceClient(conn)
	ctx := context.Background()
	_, err = paymentClient.Create(ctx, &payment.CreatePaymentRequest{Price: 23})
	if err != nil {
		log.Println("Don't worry, we don't expect to see it is working.")
	}
}
