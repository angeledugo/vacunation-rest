package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/angeledugo/vacunation-rest/handlers"
	"github.com/angeledugo/vacunation-rest/middleware"
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
	// Path Api
	api := r.PathPrefix("/api/v1").Subrouter()
	//  Health Check Route Route
	api.Use(middleware.CheckAuthMiddleware(s))

	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/signin", handlers.LoginHandler(s)).Methods(http.MethodPost)

	// Drugs routes
	api.HandleFunc("/drugs", handlers.InsertDrugHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/drugs/{id}", handlers.GetDrugByIdHandler(s)).Methods(http.MethodGet)
	api.HandleFunc("/drugs/{id}", handlers.UpdateDrugHandler(s)).Methods(http.MethodPut)
	api.HandleFunc("/drugs/{id}", handlers.DeleteDrugHandler(s)).Methods(http.MethodDelete)
	api.HandleFunc("/drugs", handlers.ListDrugHandler(s)).Methods(http.MethodGet)

	// Vaccination routes
	api.HandleFunc("/vaccinations", handlers.InsertVaccinationHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/vaccinations/{id}", handlers.GetVaccinationByIdHandler(s)).Methods(http.MethodGet)
	api.HandleFunc("/vaccinations/{id}", handlers.UpdateVaccinationHandler(s)).Methods(http.MethodPut)
	api.HandleFunc("/vaccinations/{id}", handlers.DeleteVaccinationHandler(s)).Methods(http.MethodDelete)
	api.HandleFunc("/vaccinations", handlers.ListVaccinationHandler(s)).Methods(http.MethodGet)
}
