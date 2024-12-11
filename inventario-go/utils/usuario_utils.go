package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go" //token da validacion
)

var jwtSecret = []byte("your_secret_key")

type Claims struct {
	UsuarioID uint   `json:"user_id"`
	Rol       string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(userID uint, role string) (string, error) {
	claims := Claims{
		UsuarioID: userID,
		Rol:       role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(token string) (*Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := parsedToken.Claims.(*Claims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, err
}
