package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pranotobudi/go-akselerasi-batch3-project3/middleware"
)

type AuthService interface {
	CreateAccessToken(role string, ID uint) (string, error)
}

type authService struct {
}

func NewAuthService() *authService {
	return &authService{}
}

func (s *authService) CreateAccessToken(role string, ID uint) (string, error) {
	claims := &middleware.JwtCustomClaims{
		Role: role,
		ID:   ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedKey, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return signedKey, err
	}

	return signedKey, nil
}
