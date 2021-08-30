package _interface

import (
	"fmt"

	"github.com/adolsalamanca/go-clean-boilerplate/internal/domain/entities"
	"github.com/adolsalamanca/go-clean-boilerplate/internal/domain/repository"
	"github.com/adolsalamanca/go-clean-boilerplate/internal/infrastructure/config"
	"github.com/adolsalamanca/go-clean-boilerplate/internal/infrastructure/persistence"
)

type Facade struct {
	repo repository.ItemRepository
	// collector
	// tracing
	// logger
}

func NewFacade(config config.Provider) *Facade {
	return &Facade{
		repo: persistence.NewPsqlRepository(config),
	}
}

func (s Facade) GetItems() ([]entities.Item, error) {
	i, err := s.repo.FindAllItems()
	if err != nil {
		fmt.Printf("error getting items, %v", err)
		return nil, err
	}

	return i, nil
}

func (s Facade) CreateItem(i entities.Item) error {
	err := s.repo.StoreItem(i)
	if err != nil {
		fmt.Printf("error creating items, %v", err)
		return err
	}
	return nil
}
