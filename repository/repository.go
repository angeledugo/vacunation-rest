package repository

import (
	"context"

	"github.com/angeledugo/vacunation-rest/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertDrug(ctx context.Context, drug *models.Drug) error
	GetDrugById(ctx context.Context, id string) (*models.Drug, error)
	UpdateDrug(ctx context.Context, drug *models.Drug) error
	DeleteDrug(ctx context.Context, id string) error
	ListDrug(ctx context.Context, page uint64) ([]*models.Drug, error)
	InsertVaccination(ctx context.Context, drug *models.Vaccination) error
	GetVaccinationById(ctx context.Context, id string) (*models.Vaccination, error)
	UpdateVaccination(ctx context.Context, drug *models.Vaccination) error
	DeleteVaccination(ctx context.Context, id string) error
	ListVaccination(ctx context.Context, page uint64) ([]*models.Vaccination, error)
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func InsertDrug(ctx context.Context, drug *models.Drug) error {
	return implementation.InsertDrug(ctx, drug)
}

func GetDrugById(ctx context.Context, id string) (*models.Drug, error) {
	return implementation.GetDrugById(ctx, id)
}

func UpdateDrug(ctx context.Context, drug *models.Drug) error {
	return implementation.UpdateDrug(ctx, drug)
}

func DeleteDrug(ctx context.Context, id string) error {
	return implementation.DeleteDrug(ctx, id)
}

func ListDrug(ctx context.Context, page uint64) ([]*models.Drug, error) {
	return implementation.ListDrug(ctx, page)
}

func InsertVaccination(ctx context.Context, vaccination *models.Vaccination) error {
	return implementation.InsertVaccination(ctx, vaccination)
}

func GetVaccinationById(ctx context.Context, id string) (*models.Vaccination, error) {
	return implementation.GetVaccinationById(ctx, id)
}

func UpdateVaccination(ctx context.Context, vaccination *models.Vaccination) error {
	return implementation.UpdateVaccination(ctx, vaccination)
}

func DeleteVaccination(ctx context.Context, id string) error {
	return implementation.DeleteVaccination(ctx, id)
}

func ListVaccination(ctx context.Context, page uint64) ([]*models.Vaccination, error) {
	return implementation.ListVaccination(ctx, page)
}
