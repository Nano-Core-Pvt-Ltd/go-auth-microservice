package grpcserver

import (
	"context"
	"fmt"
	"net"

	"github.com/create-go-app/fiber-go-template/platform/database"
	authpb "github.com/create-go-app/fiber-go-template/proto/auth"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ProfileServer exposes the user profile data needed by other services.
type ProfileServer struct {
	authpb.UnimplementedAuthServiceServer
}

// NewProfileServer returns a ready-to-register profile gRPC server.
func NewProfileServer() *ProfileServer {
	return &ProfileServer{}
}

// GetMyProfile loads the basic profile for the requested user.
func (s *ProfileServer) GetMyProfile(ctx context.Context, req *authpb.GetMyProfileRequest) (*authpb.GetMyProfileResponse, error) {
	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id: %v", err)
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to open database connection: %v", err)
	}

	user, err := db.GetUserByID(userID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	return &authpb.GetMyProfileResponse{
		Id:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// StartGRPCServer starts the Auth profile gRPC server on the given address.
func StartGRPCServer(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	server := grpc.NewServer()
	authpb.RegisterAuthServiceServer(server, NewProfileServer())

	fmt.Printf("Auth gRPC server listening on %s\n", addr)
	if err := server.Serve(lis); err != nil {
		return fmt.Errorf("gRPC server failed: %w", err)
	}
	return nil
}
