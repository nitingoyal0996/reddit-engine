package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nitingoyal0996/reddit-clone/messages"
	"github.com/nitingoyal0996/reddit-clone/models"
	"github.com/nitingoyal0996/reddit-clone/repositories"
)

type AuthService struct {
	userRepo 	repositories.UserRepository
	jwtSecret	[]byte
}

func NewAuthService(userRepo repositories.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtSecret: []byte(jwtSecret),
	}
}

func (s *AuthService) RegisterNewUser(username, email, password string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Email: email,
		Password: password,
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) GenerateToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &messages.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "Reddit",
		},
		UserId: user.ID,
		Username: user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *AuthService) ValidateToken(tokenString string) (*messages.Claims, error) {
	claims := &messages.Claims{}
	token, err := jwt.ParseWithClaims(
						tokenString, 
						claims, 
						func(token *jwt.Token) (interface{}, error) {
							return s.jwtSecret, nil
						})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*messages.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepo.CheckPassword(username, password)

	if err != nil {
		return "", err
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}