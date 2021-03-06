package handlers

import (

	"context"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	repos "github.com/mikedutuandu/micro_app/micro_learner/repositories"
	services "github.com/mikedutuandu/micro_app/micro_learner/services"
	//publishers "github.com/mikedutuandu/micro_app/micro_teacher/publishers"


	"fmt"
	microlearnerpb "github.com/mikedutuandu/micro_app/micro_learner/protos"

)

type AuthService struct {

}



func (srv *AuthService) Auth(ctx context.Context, req *microlearnerpb.AuthRequest) (*microlearnerpb.AuthResponse, error) {
	log.Println("Logging in with:", req.Email, req.Password)
	user, err := repos.AuthRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Not found: %v", err),
		)
	}
	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Wrong password: %v", err),
		)
	}
	token, err := services.TokenSv.Encode(user)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return &microlearnerpb.AuthResponse{
		Token: &microlearnerpb.Token{
			Valid:   true,
			Token: token,
		},
	}, nil
}

func (srv *AuthService) CreateUser(ctx context.Context, req *microlearnerpb.CreateUserRequest) (*microlearnerpb.CreateUserResponse, error) {

	log.Println("Creating user: ", req.User)
	user, err := repos.AuthRepo.GetUserByEmail(req.User.Email)
	if user != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Email is existed"),
		)
	}
	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	req.User.Password = string(hashedPass)
	_,err = repos.AuthRepo.CreateUser(req.User)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	token, err := services.TokenSv.Encode(req.User)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return &microlearnerpb.CreateUserResponse{
		Token: &microlearnerpb.Token{
			Valid:   true,
			Token: token,
		},
	}, nil
}

func (srv *AuthService) ValidateToken(ctx context.Context, req *microlearnerpb.ValidateTokenRequest) (*microlearnerpb.ValidateTokenResponse,error) {

	//Decode token
	claims, err := services.TokenSv.Decode(req.Token.Token)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	if claims.User.Email == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid token"),
		)
	}

	return &microlearnerpb.ValidateTokenResponse{
		Token: req.Token,
	}, nil
}
