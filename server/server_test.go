package main

import (
	"context"
	"testing"
	"time"

	pb "github.com/ProTechCentroEste/meeting/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type mockServer struct {
	pb.UnimplementedMeetingServiceServer
}

func (s *mockServer) Chat(req *pb.ChatRequest, stream pb.MeetingService_ChatServer) error {
	// Mock implementation for the Chat method
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		response := &pb.ChatResponse{
			UserId:    msg.UserId,
			Message:   msg.Message,
			Timestamp: time.Now().Unix(),
		}

		if err := stream.Send(response); err != nil {
			return err
		}
	}
}

func (s *mockServer) SendMessage(ctx context.Context, msg *pb.ChatMessage) (*pb.ChatMessage, error) {
	// Mock implementation for SendMessage method
	return &pb.ChatMessage{
		User:      msg.User,
		Message:   msg.Message,
		Timestamp: time.Now().Unix(),
	}, nil
}

func TestLogin(t *testing.T) {
	// Initialize the server and other dependencies
	s := &mockServer{}

	// Create a mock request with appropriate field names
	req := &pb.LoginRequest{
		AuthProvider: "Google",
		Token:        "sample_token",
	}

	// Call the Login method
	res, err := s.Login(context.Background(), req)

	// Assert no error
	assert.NoError(t, err)

	// Assert expected response values
	assert.Equal(t, "12345", res.UserId)
	assert.True(t, res.Success)
}

func TestChat(t *testing.T) {
	// Initialize the server and other dependencies
	s := &mockServer{}

	// Set up the gRPC client connection
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMeetingServiceClient(conn)

	// Create a chat stream
	stream, err := client.Chat(context.Background())
	if err != nil {
		t.Fatalf("could not create chat stream: %v", err)
	}

	// Send a chat message
	err = stream.Send(&pb.ChatRequest{
		UserId:  "user1",
		Message: "Hello",
	})
	if err != nil {
		t.Fatalf("could not send message: %v", err)
	}

	// Receive a chat message
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("could not receive message: %v", err)
	}

	// Assert response values
	assert.Equal(t, "user1", resp.UserId)
	assert.Equal(t, "Hello", resp.Message)
}

func TestSendMessage(t *testing.T) {
	// Initialize the server and other dependencies
	s := &mockServer{}

	// Create a mock request
	req := &pb.ChatMessage{
		User:      "user1",
		Message:   "Hello",
		Timestamp: time.Now().Unix(),
	}

	// Call the SendMessage method
	res, err := s.SendMessage(context.Background(), req)

	// Assert no error
	assert.NoError(t, err)

	// Assert response values
	assert.Equal(t, "user1", res.User)
	assert.Equal(t, "Hello", res.Message)
}
