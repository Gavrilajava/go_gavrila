package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "go-gavrila/task-14/messages"
)

func main() {
	conn, err := grpc.Dial("localhost:8005",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewHubClient(conn)

	ctx := context.Background()

	client.Send(context.Background(), &pb.Message{Text: "Ground control to Major Tom..."})
	client.Send(context.Background(), &pb.Message{Text: "Ground control to Major Tom:"})
	client.Send(context.Background(), &pb.Message{Text: "Lock your Soyuz hatch and put your helmet on!"})

	err = printAllBooksOnserver(ctx, client)
	if err != nil {
		log.Fatal(err)
	}
}

func printAllBooksOnserver(ctx context.Context, client pb.HubClient) error {
	fmt.Println("Printing all the messages from server:")
	stream, err := client.Messages(context.Background(), &pb.Empty{})
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			message, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			fmt.Printf("Message: %v\n", message)
		}
	}
}
