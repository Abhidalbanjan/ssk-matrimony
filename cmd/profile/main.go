package main

import (
	"database/sql"
	"log"
	"net"

	db "github.com/Abhidalbanjan/ssk-matrimony/db/generated"
	pb "github.com/Abhidalbanjan/ssk-matrimony/proto/profile"
	profileService "github.com/Abhidalbanjan/ssk-matrimony/services/profile"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Connect to Postgres (docker container)
	connStr := "postgresql://root:admin@localhost:5432/ssk_matrimony?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	// Initialize SQLC Queries
	queries := db.New(dbConn)

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register ProfileService
	pb.RegisterProfileServiceServer(grpcServer, profileService.NewProfileServer(queries))

	// Enable reflection for grpcurl
	reflection.Register(grpcServer)

	log.Println("🚀 ProfileService gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
