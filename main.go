package main

import (
	loadEnv "article_app/helper"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	r := mux.NewRouter()
	port := loadEnv.GetEnv("APP_PORT", "8000")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Author":  "Aliamat Parinduri",
			"Version": "0.0.1",
		})
	})

	log.Println("Server running on http://localhost:" + port)
	http.ListenAndServe(":"+port, r)
}
