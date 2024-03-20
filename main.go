package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/angeledugo/vacunation-rest/handlers"
	"github.com/angeledugo/vacunation-rest/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRECT := os.Getenv("JWT_SECRECT")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRECT,
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)

}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.SignUpHandler(s)).Methods(http.MethodGet)
}
