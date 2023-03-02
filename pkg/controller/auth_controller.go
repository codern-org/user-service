package controller

import (
	"context"

	"github.com/codern-org/user-service/pkg/pb"
	"github.com/codern-org/user-service/pkg/port"
)

type AuthController struct {
	pb.UnimplementedAuthServiceServer

	authService    port.AuthService
	googleService  port.GoogleService
	sessionService port.SessionService
	userService    port.UserService
}

func NewAuthController(
	authService port.AuthService,
	googleService port.GoogleService,
	sessionService port.SessionService,
	userService port.UserService,
) *AuthController {
	return &AuthController{
		authService:    authService,
		googleService:  googleService,
		sessionService: sessionService,
		userService:    userService,
	}
}

func (controller *AuthController) GetGoogleOAuthUrl(
	ctx context.Context, in *pb.GoogleOAuthUrlRequest,
) (*pb.GoogleOAuthUrlResponse, error) {
	return &pb.GoogleOAuthUrlResponse{
		Url: controller.googleService.GetOAuthUrl(),
	}, nil
}
