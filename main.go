package main

import (
	"log"
	"net/http"
)

func main() {
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

		http.SetCookie(w, &http.Cookie{
			Name:  "session_token",
			Value: username,
			Path:  "/",
		})

		w.Write([]byte("Logged in successfully"))
	})

	log.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
