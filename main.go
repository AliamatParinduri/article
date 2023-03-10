package main

import (
	"article_app/helper"
	auth "article_app/modules/auth/delivery/http"
	post "article_app/modules/post/delivery/http"
	tag "article_app/modules/tag/delivery/http"
	user "article_app/modules/user/delivery/http"
	"article_app/repository"
	"encoding/json"
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	r := mux.NewRouter()
	port := helper.GetEnv("APP_PORT", "8000")

	// CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Author":  "Aliamat Parinduri",
			"Version": "0.0.1",
		})
	})

	auth.AuthRouter(r)
	user.UserRouter(r)
	tag.TagRouter(r)
	post.PostRouter(r)

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		InitCommands()
	} else {
		log.Println("Server running on http://localhost:" + port)
		log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(r)))
	}
}

func InitCommands() {
	var serverPG = repository.ServerPG{}
	serverPG.InitCommands()
}
