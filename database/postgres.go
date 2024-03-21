package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/angeledugo/vacunation-rest/models"
	_ "github.com/lib/pq" // Import the Postgres
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil

}

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, name, email, password) values ($1,$2, $3, $4)", user.Id, user.Name, user.Email, user.Password)

	return err
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, email FROM users where id = $1", id) // TODO: implement query with parameter

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Name, &user.Email); err == nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil

}

func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, email, password FROM users where email = $1", email) // TODO: implement query with parameter

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil

}

func (repo *PostgresRepository) InsertDrug(ctx context.Context, drug *models.Drug) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO drugs (id, name, approved, min_dose, max_dose, available_at) values ($1,$2, $3, $4, $5, $6)", drug.Id, drug.Name, drug.Approved, drug.Min_dose, drug.Max_dose, drug.Available_at)
	return err
}

func (repo *PostgresRepository) GetDrugById(ctx context.Context, id string) (*models.Drug, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, approved, min_dose, max_dose, available_at FROM drugs where id = $1", id) // TODO: implement query with parameter

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var drug = models.Drug{}
	for rows.Next() {
		if err = rows.Scan(&drug.Id, &drug.Name, &drug.Approved, &drug.Min_dose, &drug.Max_dose, &drug.Available_at); err == nil {
			return &drug, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &drug, nil

}

func (repo *PostgresRepository) UpdateDrug(ctx context.Context, drug *models.Drug) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE drugs SET name = $1, approved = $2, min_dose = $3, max_dose = $4, available_at = $5 where id = $6", drug.Name, drug.Approved, drug.Min_dose, drug.Max_dose, drug.Available_at, drug.Id)

	return err
}

func (repo *PostgresRepository) ListDrug(ctx context.Context, page uint64) ([]*models.Drug, error) {

	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, approved, min_dose, max_dose, available_at FROM drugs LIMIT $1 OFFSET $2", 2, page*2)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var drugs []*models.Drug

	for rows.Next() {
		var drug = models.Drug{}
		if err = rows.Scan(&drug.Id, &drug.Name, &drug.Approved, &drug.Min_dose, &drug.Max_dose, &drug.Available_at); err == nil {

			drugs = append(drugs, &drug)
		}
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return drugs, nil
}

func (repo *PostgresRepository) DeleteDrug(ctx context.Context, id string) error {

	_, err := repo.db.ExecContext(ctx, "DELETE FROM drugs WHERE id = $1", id)

	return err
}

func (repo *PostgresRepository) InsertVaccination(ctx context.Context, vaccination *models.Vaccination) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO vaccinations (id, name, drug_id, dose, date) values ($1,$2, $3, $4, $5)", vaccination.Id, vaccination.Name, vaccination.Drug_id, vaccination.Dose, vaccination.Date)

	return err
}

func (repo *PostgresRepository) GetVaccinationById(ctx context.Context, id string) (*models.Vaccination, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, drug_id, dose, date FROM vaccinations where id = $1", id) // TODO: implement query with parameter

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var vaccination = models.Vaccination{}
	for rows.Next() {
		if err = rows.Scan(&vaccination.Id, &vaccination.Name, &vaccination.Drug_id, &vaccination.Dose, &vaccination.Date); err == nil {
			return &vaccination, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &vaccination, nil
}

func (repo *PostgresRepository) UpdateVaccination(ctx context.Context, vaccination *models.Vaccination) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE vaccinations SET name = $1, drug_id = $2, dose = $3, date = $4 where id = $5", vaccination.Name, vaccination.Drug_id, vaccination.Dose, vaccination.Date, vaccination.Id)

	return err
}

func (repo *PostgresRepository) DeleteVaccination(ctx context.Context, id string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM vaccinations WHERE id = $1", id)

	return err
}

func (repo *PostgresRepository) ListVaccination(ctx context.Context, page uint64) ([]*models.Vaccination, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, drug_id, dose, date FROM vaccinations LIMIT $1 OFFSET $2", 2, page*2)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var vaccinations []*models.Vaccination

	for rows.Next() {
		var vaccination = models.Vaccination{}
		if err = rows.Scan(&vaccination.Id, &vaccination.Name, &vaccination.Drug_id, &vaccination.Dose, &vaccination.Date); err == nil {

			vaccinations = append(vaccinations, &vaccination)
		}
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return vaccinations, nil
}
