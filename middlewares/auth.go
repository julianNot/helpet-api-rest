package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func loadSecretKey() string {
	err := godotenv.Load()
	if err != nil {
    log.Fatal("Error loading .env file")
  }
	SecretKey := os.Getenv("MY_SECRET_KEY")
	return SecretKey
}

func AuthMiddlewarePartner(next http.Handler) http.Handler {
	secretKey := loadSecretKey()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if r.URL.Path == "/auth/partners" {
			next.ServeHTTP(w, r)
			return
		}
		if tokenString == "" {
			http.Error(w, "Authorization token required", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
