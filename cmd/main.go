package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/codern-org/user-service/pkg/controller"
	"github.com/codern-org/user-service/pkg/infrastructure"
	"github.com/codern-org/user-service/pkg/pb"
	"github.com/codern-org/user-service/pkg/repository"
	"github.com/codern-org/user-service/pkg/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Load .env
	if error := godotenv.Load(); error != nil {
		log.Fatalln("error loading .env file")
	}

	// Database
	db, error := infrastructure.OpenMySqlConnection(os.Getenv("MYSQL_URI"))
	if error != nil {
		log.Fatalf("cannot open database connection %v", error)
	} else {
		log.Println("database connection established")
	}

	// Repository
	sessionRepository := repository.NewSessionRepository(db)
	userRepository := repository.NewUserRepository(db)

	// Service
	googleService := service.NewGoogleService(
		os.Getenv("GOOGLE_CLIENT_ID"),
		os.Getenv("GOOGLE_CLIENT_SECRET"),
		os.Getenv("GOOGLE_REDIRECT_URI"),
	)
	sessionService := service.NewSessionService(os.Getenv("SESSION_SECRET"), sessionRepository)
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(googleService, sessionService, userService)

	// GRPC server
	port := os.Getenv("PORT")
	listener, error := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if error != nil {
		log.Fatalf("failed to listen to port %s: %v", port, error)
	}

	grpcServer := grpc.NewServer()

	// Register RPC services
	authController := controller.NewAuthController(
		authService, googleService, sessionService, userService,
	)
	pb.RegisterAuthServiceServer(grpcServer, authController)

	log.Printf("listening to port %s\n", port)
	if error := grpcServer.Serve(listener); error != nil {
		log.Fatalf("failed to server grpc server: %v", error)
	}
}
