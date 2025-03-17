package auth

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthJWT struct {
	log *log.Logger
}

func New(log *log.Logger) *AuthJWT {
	return &AuthJWT{log: log}
}

func (i *AuthJWT) GenerateJWT(userID string) (string, error) {
	exp, err := strconv.Atoi(os.Getenv("EXPIRES_AT"))
	secret := os.Getenv("JWT_SECRET")
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * time.Duration(exp)).Unix(),
		"iat":     time.Now().Unix(),
	}

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTk, err := tk.SignedString([]byte(secret))
	return signedTk, err
}

func (i *AuthJWT) VerifyJWT(tk string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenUnverifiable
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}