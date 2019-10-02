package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mikedutuandu/micro_app/gateway_teacher/services"
	micro_teacher_pb "github.com/mikedutuandu/micro_app/micro_teacher/protos"
	"github.com/spf13/viper"
	"net/http"
	"os"
)



func Init() {


	e := echo.New()

	//1. Group non auth====================================
	nonAuth := e.Group("")
	// Middleware
	nonAuth.Use(middleware.Logger())
	nonAuth.Use(middleware.Recover())
	nonAuth.Use(unrestricted)
	//CORS
	nonAuth.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Routes
	nonAuth.GET("/test", test)
	nonAuth.POST("/login", login)
	nonAuth.POST("/signup", signup)


	//2. Group auth=========================================
	auth := e.Group("/auth")
	// Middleware
	auth.Use(middleware.Logger())
	auth.Use(middleware.Recover())
	auth.Use(restricted)
	//CORS
	auth.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Routes
	auth.GET("/test", testAuth)

	// Start server
	var port = os.Getenv("PORT")
	if port == ""{
		port = viper.GetString("grpc.portDefault")
	}
	e.Logger.Fatal(e.Start(":"+port))
}

func  unrestricted(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("ok men")
		return next(c)
	}
}
func  restricted(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")
		tokenStr := authorization[7:]
		println(tokenStr)
		token := micro_teacher_pb.Token{Token:tokenStr,Valid:false}
		_, err := services.MicroCLI.MicroTeacherAuthClient.ValidateToken(context.Background(),&micro_teacher_pb.ValidateTokenRequest{Token:&token})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "must login")
		}
		return next(c)
	}
}