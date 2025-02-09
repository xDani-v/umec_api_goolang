package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("DataEncrypt"))

func GenerarToken(usuarioID uint64, correo string, duracion time.Duration) (string, error) {
	expirationTime := time.Now().Add(duracion)

	claims := jwt.MapClaims{
		"id":    usuarioID,
		"email": correo,
		"exp":   expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
