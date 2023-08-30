package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "go-gavrila/task-14/messages"
)

var counter int64

type Messages struct {
	Data []*pb.Message
	pb.UnimplementedHubServer
}

func (b *Messages) Messages(_ *pb.Empty, stream pb.Hub_MessagesServer) error {
	for i := range b.Data {
		stream.Send(b.Data[i])
	}
	return nil
}

func (b *Messages) Send(_ context.Context, m *pb.Message) (*pb.Empty, error) {
	counter++
	m.Id = counter
	m.SentAt = timestamppb.Now()
	b.Data = append(b.Data, m)
	return new(pb.Empty), nil
}

func main() {
	srv := Messages{}

	lis, err := net.Listen("tcp", ":8005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHubServer(grpcServer, &srv)
	grpcServer.Serve(lis)
}
