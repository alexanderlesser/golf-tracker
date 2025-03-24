package repositories

import (
	"database/sql"
	"fmt"

	"github.com/alexanderlesser/golf-tracker/internal/models"
)

type ClubRepository interface {
	GetAll() ([]models.Club, error)
	GetByID(id int) (models.Club, error)
	Update(c models.Club) error
	Remove(id int) error
	Create(c models.Club) (models.Club, error)
}

type clubRepo struct {
	db *sql.DB
}

func NewClubRepo(db *sql.DB) ClubRepository {
	return &clubRepo{db: db}
}

func (r *clubRepo) GetAll() ([]models.Club, error) {
	rows, err := r.db.Query("SELECT * FROM clubs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clubs []models.Club
	for rows.Next() {
		var club models.Club
		if err := rows.Scan(&club.ID, &club.Name, &club.Loft); err != nil {
			return nil, err
		}
		clubs = append(clubs, club)
	}
	return clubs, nil
}

func (r *clubRepo) GetByID(id int) (models.Club, error) {
	var club models.Club
	query := "SELECT id, name, loft FROM clubs WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&club.ID, &club.Name, &club.Loft)
	if err != nil {
		return club, err
	}
	return club, nil
}

func (r *clubRepo) Update(c models.Club) error {
	query := "UPDATE clubs SET name = ?, loft = ? WHERE id = ?"
	_, err := r.db.Exec(query, c.Name, c.Loft, c.ID)
	if err != nil {
		return fmt.Errorf("failed to update club: %w", err)
	}
	return nil
}

func (r *clubRepo) Remove(id int) error {
	query := "DELETE FROM clubs WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to remove club: %w", err)
	}
	return nil
}

func (r *clubRepo) Create(c models.Club) (models.Club, error) {
	var createdClub models.Club

	// SQL query to insert a new row into the clubs table and retrieve the generated ID
	query := "INSERT INTO clubs (name, loft) VALUES (?, ?)"

	// Execute the query, passing in the club's name and loft
	result, err := r.db.Exec(query, c.Name, c.Loft)
	if err != nil {
		return models.Club{}, fmt.Errorf("failed to create club: %w", err)
	}

	// Get the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return models.Club{}, fmt.Errorf("failed to get last inserted ID: %w", err)
	}

	// Retrieve the newly created club from the database
	query = "SELECT id, name, loft FROM clubs WHERE id = ?"
	err = r.db.QueryRow(query, lastID).Scan(&createdClub.ID, &createdClub.Name, &createdClub.Loft)
	if err != nil {
		return models.Club{}, fmt.Errorf("failed to retrieve created club: %w", err)
	}

	return createdClub, nil
}
