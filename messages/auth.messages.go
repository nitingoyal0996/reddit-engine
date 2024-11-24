package messages

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId 		uint64 	`json:"user_id"`
	Username 	string 	`json:"username"`
}
