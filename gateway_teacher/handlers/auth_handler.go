package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/mikedutuandu/micro_app/gateway_teacher/services"

	micro_teacher_pb "github.com/mikedutuandu/micro_app/micro_teacher/protos"
	"net/http"
	//"context"
	//"log"
	//"fmt"
	//services "github.com/mikedutuandu/micro_app/gateway_teacher/services"
)

func test(c echo.Context) error {

	return c.JSON(http.StatusOK,"ok men")
}

func login(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	email := m["email"].(string)
	password := m["password"].(string)
	res, err := services.MicroCLI.MicroTeacherAuthClient.Auth(context.Background(),&micro_teacher_pb.AuthRequest{Email: email,Password:password})
	if err != nil {
		return c.JSON(http.StatusUnauthorized,err)
	}

	return c.JSON(http.StatusOK,res)
}
func signup(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	email := m["email"].(string)
	password := m["password"].(string)
	user := micro_teacher_pb.User{Email:email,Password:password}
	res, err := services.MicroCLI.MicroTeacherAuthClient.CreateUser(context.Background(),&micro_teacher_pb.CreateUserRequest{User:&user})
	if err != nil {
		return c.JSON(http.StatusBadRequest,err)
	}

	return c.JSON(http.StatusOK,res)

}

func testAuth(c echo.Context) error {


	return c.JSON(http.StatusOK,"logined")

}