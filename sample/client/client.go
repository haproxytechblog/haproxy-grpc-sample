package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	// The Protobuf generated file
	creator "app/codenamecreator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	address := os.Getenv("SERVER_ADDRESS")
	crt := os.Getenv("TLS_CERT")

	creds, err := credentials.NewClientTLSFromFile(crt, "")
	if err != nil {
		log.Fatalf("Failed to load TLS certificate")
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Did not connect, %v", err)
	}
	defer conn.Close()

	client := creator.NewCodenameCreatorClient(conn)
	ctx := context.Background()

	// Optional: add some metadata
	ctx = metadata.AppendToOutgoingContext(ctx, "mysecretpassphrase", "abc123")

	fmt.Println("Generating codenames...")

	stream, err := client.KeepGettingCodenames(ctx)

	if err != nil {
		log.Fatalf("Could not get stream, %v", err)
	}

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("%v", err)
			}

			log.Printf("Received: %s\n", in.Name)
		}
	}()

	category := "Science"
	for {
		if category == "Science" {
			category = "Animals"
		} else {
			category = "Science"
		}

		err := stream.Send(&creator.NameRequest{Category: category})
		if err != nil {
			log.Fatalf("%v", err)
		}
		time.Sleep(10 * time.Second)
	}
}
