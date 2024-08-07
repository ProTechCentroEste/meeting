package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"time"

	pb "github.com/ProTechCentroEste/meeting/proto"
	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMeetingServiceServer
	db *sql.DB
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Implement authentication logic here
	return &pb.LoginResponse{UserId: "12345", Success: true}, nil
}

func (s *server) Chat(req *pb.ChatRequest, stream pb.MeetingService_ChatServer) error {
	// Implement chat streaming logic here
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		// Process received message and send a response
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

func (s *server) SendMessage(ctx context.Context, msg *pb.ChatMessage) (*pb.ChatMessage, error) {
	// Implement send message logic here
	return &pb.ChatMessage{
		User:      msg.User,
		Message:   msg.Message,
		Timestamp: time.Now().Unix(),
	}, nil
}

func (s *server) FindNearbyUsers(ctx context.Context, req *pb.FindNearbyUsersRequest) (*pb.FindNearbyUsersResponse, error) {
	users, err := findNearbyUsers(s.db, req.Latitude, req.Longitude, req.Radius)
	if err != nil {
		return nil, err
	}

	var userResponses []*pb.User
	for _, user := range users {
		userResponses = append(userResponses, &pb.User{
			Id:    int32(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return &pb.FindNearbyUsersResponse{Users: userResponses}, nil
}

func findNearbyUsers(db *sql.DB, userLat float64, userLon float64, radius float64) ([]User, error) {
	query := `
    SELECT id, name, email
    FROM users
    WHERE ST_DWithin(
        location,
        ST_MakePoint($1, $2)::geography,
        $3
    );
    `
	rows, err := db.Query(query, userLon, userLat, radius)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	connStr := "postgresql://username:password@localhost:5432/mydb"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMeetingServiceServer(s, &server{db: db})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
