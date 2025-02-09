package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xDani-v/umec_api_goolang/utils"
)

var jwtKey = []byte(os.Getenv("DataEncrypt"))

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response := utils.ResponseMsg{
				Msg:    "No autorizado: falta el encabezado de autorización",
				Data:   nil,
				Status: http.StatusUnauthorized,
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) != 2 {
			response := utils.ResponseMsg{
				Msg:    "Formato de token inválido",
				Data:   nil,
				Status: http.StatusUnauthorized,
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		tokenString := parts[1]
		claims := &jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			response := utils.ResponseMsg{
				Msg:    "Token inválido",
				Data:   nil,
				Status: http.StatusUnauthorized,
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
