package messages

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId 		uint 	`json:"user_id"`
	Username 	string 	`json:"username"`
}

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

type RegisterResponse struct {
    ID   		uint
    Username 	string
    Error    	string
}

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Token   string
	Error   string
}


type TokenValidationRequest struct {
	Token 	string
}

type TokenValidationResponse struct {
	Valid  	bool
	Claims 	*Claims
	Error 	string
}
