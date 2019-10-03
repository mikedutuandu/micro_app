package tests

import (
	"strings"
	"testing"

	auth "github.com/mikedutuandu/micro_app/micro_learner/protos"
	services "github.com/mikedutuandu/micro_app/micro_learner/services"
)

var (
	user = &auth.User{
		Id:    "abc123",
		Email: "ewan.valentine89@gmail.com",
	}
)

type MockRepo struct{}


func (repo *MockRepo) CreateUser(user *auth.User) error {

	return nil
}

func (repo *MockRepo) GetUserByEmail(email string) (*auth.User, error) {
	var user *auth.User
	return user, nil
}

func newInstance() services.Authable {
	return &services.TokenService{}
}

func TestCanCreateToken(t *testing.T) {
	srv := newInstance()
	token, err := srv.Encode(user)
	if err != nil {
		t.Fail()
	}

	if token == "" {
		t.Fail()
	}

	if len(strings.Split(token, ".")) != 3 {
		t.Fail()
	}
}

func TestCanDecodeToken(t *testing.T) {
	srv := newInstance()
	token, err := srv.Encode(user)
	if err != nil {
		t.Fail()
	}
	claims, err := srv.Decode(token)
	if err != nil {
		t.Fail()
	}
	if claims.User == nil {
		t.Fail()
	}
	if claims.User.Email != "ewan.valentine89@gmail.com" {
		t.Fail()
	}
}

func TestThrowsErrorIfTokenInvalid(t *testing.T) {
	srv := newInstance()
	_, err := srv.Decode("nope.nope.nope")
	if err == nil {
		t.Fail()
	}
}
