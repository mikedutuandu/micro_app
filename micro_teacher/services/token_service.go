package services

import (
	"github.com/spf13/viper"
	"time"

	"github.com/dgrijalva/jwt-go"
	auth "github.com/mikedutuandu/micro_app/micro_teacher/protos"
)

var (

	// Define a secure key string used
	// as a salt when hashing our tokens.
	// Please make your own way more secure than this,
	// use a randomly generated md5 hash or something.
	secretKey = viper.GetString("jwt.secretKey")
	key = []byte(secretKey)
)

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *auth.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *auth.User) (string, error)
}


type TokenService struct {
}


// Decode a token string into a token object
func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil{
		return nil, err
	}

	// Validate the token and return the custom claims
	claims, ok := token.Claims.(*CustomClaims)
	if ok == false {
		return nil,err
	}
	return claims,nil

}

// Encode a claim into a JWT
func (srv *TokenService) Encode(user *auth.User) (string, error) {

	tokenExpired := viper.GetInt("jwt.tokenExpired")
	expireToken := time.Now().Add(time.Hour * time.Duration(tokenExpired)).Unix()

	// Create the Claims
	issuer := viper.GetString("jwt.issuer")
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer: issuer,
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	tk,err := token.SignedString(key)
	return tk,err
}
