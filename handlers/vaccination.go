package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/angeledugo/vacunation-rest/models"
	"github.com/angeledugo/vacunation-rest/repository"
	"github.com/angeledugo/vacunation-rest/server"
	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

type UpsertVaccinationRequest struct {
	Id      string `json:"id"`
	Name    string `json:"name" validate:"required,min=3,max=20"`
	Drug_id string `json:"drug_id"`
	Dose    int64  `json:"dose"`
	Date    string `json:"date"`
}

type VaccinationResponse struct {
	Id      string    `json:"id"`
	Name    string    `json:"name" validate:"required,min=3,max=20"`
	Drug_id string    `json:"drug_id"`
	Dose    int64     `json:"dose"`
	Date    time.Time `json:"date"`
}

type VaccinationUpdateResponse struct {
	Message string `json:"message"`
}

func InsertVaccinationHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vaccinationRequest = UpsertVaccinationRequest{}

		if err := json.NewDecoder(r.Body).Decode(&vaccinationRequest); err != nil {

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		date, _ := time.Parse("2006-01-02", vaccinationRequest.Date)
		vaccination := models.Vaccination{Id: id.String(), Name: vaccinationRequest.Name, Drug_id: vaccinationRequest.Drug_id, Dose: vaccinationRequest.Dose, Date: date}

		err = repository.InsertVaccination(r.Context(), &vaccination)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(VaccinationResponse{
			Id:      vaccination.Id,
			Name:    vaccination.Name,
			Drug_id: vaccination.Drug_id,
			Dose:    vaccination.Dose,
			Date:    vaccination.Date,
		})

	}
}

func GetVaccinationByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		vaccination, err := repository.GetVaccinationById(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(vaccination)
	}
}

func UpdateVaccinationHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var vaccinationRequest = UpsertVaccinationRequest{}

		if err := json.NewDecoder(r.Body).Decode(&vaccinationRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		date, _ := time.Parse("2006-01-02", vaccinationRequest.Date)

		vaccination := models.Vaccination{Id: params["id"], Name: vaccinationRequest.Name, Drug_id: vaccinationRequest.Drug_id, Dose: vaccinationRequest.Dose, Date: date}

		err := repository.UpdateVaccination(r.Context(), &vaccination)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(VaccinationUpdateResponse{
			Message: "Vaccination Update",
		})
	}
}

func DeleteVaccinationHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		err := repository.DeleteVaccination(r.Context(), params["id"])

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(VaccinationUpdateResponse{
			Message: "Vaccination deleted",
		})
	}
}

func ListVaccinationHandler(s server.Server) http.HandlerFunc {
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
		vaccinations, err := repository.ListVaccination(r.Context(), page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(vaccinations)
	}
}
