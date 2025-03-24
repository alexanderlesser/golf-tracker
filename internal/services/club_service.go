package services

import (
	"fmt"

	"github.com/alexanderlesser/golf-tracker/internal/models"
	"github.com/alexanderlesser/golf-tracker/internal/repositories"
)

type ClubService interface {
	FetchAllClubs() ([]models.Club, error)
	UpdateClubValues(id int, n string, l string) error
	RemoveClub(id int) error
	CreateClub(name string, loft string) (models.Club, error)
}

type clubService struct {
	repo repositories.ClubRepository
}

func NewClubService(repo repositories.ClubRepository) ClubService {
	return &clubService{repo: repo}
}

func (s *clubService) FetchAllClubs() ([]models.Club, error) {
	return s.repo.GetAll()
}

func (s *clubService) UpdateClubValues(id int, n string, l string) error {
	club, err := s.repo.GetByID(id)
	if err != nil {
		fmt.Println("Error fetching club by ID:", err)
		return err
	}

	club.Name = n
	club.Loft = l

	// * Save updated name
	if err := s.repo.Update(club); err != nil {
		fmt.Println("Error in repository while updating club:", err)
		return err
	}

	return nil
}

func (s *clubService) RemoveClub(id int) error {
	err := s.repo.Remove(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *clubService) CreateClub(name string, loft string) (models.Club, error) {

	newClub := models.Club{
		ID:   0,
		Name: name,
		Loft: loft,
	}

	club, err := s.repo.Create(newClub)
	if err != nil {
		return models.Club{}, err
	}

	return club, nil
}
