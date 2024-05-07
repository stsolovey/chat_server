package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func generateToken(username string, mySigningKey []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", fmt.Errorf("generateToken failed: %w", err)
	}

	return tokenString, nil
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Panic("Error loading .env file")
	}

	var mySigningKey = []byte(os.Getenv("JWT_SECRET"))
	if len(mySigningKey) == 0 {
		log.WithError(err).Panic("JWT_SECRET is not set")
	}

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)

			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)

			return
		}

		username := r.FormValue("username")
		if username == "" {
			http.Error(w, "Username is required", http.StatusBadRequest)

			return
		}

		tokenString, err := generateToken(username, mySigningKey)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)

			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "session_token",
			Value: tokenString,
			Path:  "/",
		})

		w.Write([]byte("Logged in successfully with token: " + tokenString))
	})

	log.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
