package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/angeledugo/vacunation-rest/repository"
	"github.com/gorilla/mux"

	"github.com/angeledugo/vacunation-rest/models"
	"github.com/angeledugo/vacunation-rest/server"
	"github.com/segmentio/ksuid"
)

type UpsertDrugRequest struct {
	Name         string `json:"name" validate:"required,min=3,max=20"`
	Approved     bool   `json:"approved"`
	Min_dose     int64  `json:"min_dose"`
	Max_dose     int64  `json:"max_dose"`
	Available_at string `json:"available_at"`
}

type DrugResponse struct {
	Id           string    `json:"id"`
	Name         string    `json:"name" validate:"required,min=3,max=20"`
	Approved     bool      `json:"approved"`
	Min_dose     int64     `json:"min_dose"`
	Max_dose     int64     `json:"max_dose"`
	Available_at time.Time `json:"available_at"`
}

type DrugUpdateResponse struct {
	Message string `json:"message"`
}

func InsertDrugHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var drugRequest = UpsertDrugRequest{}

		if err := json.NewDecoder(r.Body).Decode(&drugRequest); err != nil {

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		availableAt, _ := time.Parse("2006-01-02", drugRequest.Available_at)
		drug := models.Drug{Id: id.String(), Name: drugRequest.Name, Approved: drugRequest.Approved, Min_dose: drugRequest.Min_dose, Max_dose: drugRequest.Max_dose, Available_at: availableAt}

		err = repository.InsertDrug(r.Context(), &drug)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(DrugResponse{
			Id:           drug.Id,
			Name:         drug.Name,
			Approved:     drug.Approved,
			Min_dose:     drug.Min_dose,
			Max_dose:     drug.Max_dose,
			Available_at: drug.Available_at,
		})

	}
}

func GetDrugByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		drug, err := repository.GetDrugById(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(drug)
	}
}

func UpdateDrugHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var drugRequest = UpsertDrugRequest{}

		if err := json.NewDecoder(r.Body).Decode(&drugRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		availableAt, _ := time.Parse("2006-01-02", drugRequest.Available_at)

		drug := models.Drug{Id: params["id"], Name: drugRequest.Name, Approved: drugRequest.Approved, Min_dose: drugRequest.Min_dose, Max_dose: drugRequest.Max_dose, Available_at: availableAt}

		err := repository.UpdateDrug(r.Context(), &drug)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(DrugUpdateResponse{
			Message: "Drug Update",
		})
	}
}

func DeleteDrugHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		err := repository.DeleteDrug(r.Context(), params["id"])

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(DrugUpdateResponse{
			Message: "Drug deleted",
		})
	}
}

func ListDrugHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		pageStr := r.URL.Query().Get("page")
		var page = uint64(0)
		if pageStr != "" {
			page, err = strconv.ParseUint(pageStr, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		}
		drugs, err := repository.ListDrug(r.Context(), page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(drugs)

	}

}
