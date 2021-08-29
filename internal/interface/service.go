package _interface

import (
	"fmt"

	"github.com/adolsalamanca/go-rest-boilerplate/internal/domain/entities"
	"github.com/adolsalamanca/go-rest-boilerplate/internal/domain/repository"
	"github.com/adolsalamanca/go-rest-boilerplate/internal/infrastructure/config"
	"github.com/adolsalamanca/go-rest-boilerplate/internal/infrastructure/persistence"
)

type Service struct {
	repo repository.ItemRepository
}

func (s Service) GetItems() ([]*entities.Item, error) {
	_, err := s.repo.GetItems()
	if err != nil {
		fmt.Printf("error getting items, %v", err)
		return nil, err
	}

	return nil, nil
}

func (s Service) CreateItem(i entities.Item) error {
	err := s.repo.CreateItem(i)
	if err != nil {
		fmt.Printf("error creating items, %v", err)
		return err
	}
	return nil
}

func NewService(config config.Provider) *Service {
	return &Service{
		repo: persistence.NewPsqlRepository(config),
	}
}
