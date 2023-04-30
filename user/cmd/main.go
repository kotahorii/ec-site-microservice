package main

import (
	"log"
	"net"

	"github.com/microservice-ec-site/user/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialize your infrastructure, for example, the UserRepository implementation.
	userRepository := repositories.NewDynamoDBUserRepository()

	// Initialize the UserUsecase with the UserRepository.
	userUsecase := usecases.NewUserUsecase(userRepository)

	// Create the gRPC server instance
	s := grpc.NewServer()

	// Create the UserService server instance with the UserUsecase
	userService := grpc.NewUserServiceServer(userUsecase)

	// Register the UserService server with the gRPC server
	proto.RegisterUserServiceServer(s, userService)

	log.Println("gRPC server started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
