package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/angeledugo/vacunation-rest/server"
)

type UserResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(UserResponse{
			Message: "Welcome to Vacunation test",
			Status:  true,
		})
	}
}
