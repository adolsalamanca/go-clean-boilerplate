package _interface

import (
	"fmt"

	"github.com/adolsalamanca/go-clean-boilerplate/internal/application"
	"github.com/adolsalamanca/go-clean-boilerplate/internal/domain/entities"
	"github.com/adolsalamanca/go-clean-boilerplate/internal/domain/repository"
	"github.com/adolsalamanca/go-clean-boilerplate/internal/infrastructure/config"
	"github.com/adolsalamanca/go-clean-boilerplate/internal/infrastructure/persistence"
)

type Service struct {
	repo   repository.ItemRepository
	logger application.Logger
	// collector
	// tracing
}

func NewService(config config.Provider, logger *StandardLogger) *Service {
	return &Service{
		repo:   persistence.NewPsqlRepository(config),
		logger: logger,
	}
}

func (s Service) GetItems() ([]entities.Item, error) {
	i, err := s.repo.FindAllItems()
	if err != nil {
		fmt.Printf("error getting items, %v", err)
		return nil, err
	}

	return i, nil
}

func (s Service) CreateItem(i entities.Item) error {
	err := s.repo.StoreItem(i)
	if err != nil {
		fmt.Printf("error creating items, %v", err)
		return err
	}
	return nil
}
