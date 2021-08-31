package _interface

import (
	"github.com/adolsalamanca/go-clean-boilerplate/internal/domain/entities"
	"github.com/adolsalamanca/go-clean-boilerplate/internal/domain/repository"
	"github.com/adolsalamanca/go-clean-boilerplate/internal/infrastructure/persistence"
	"github.com/adolsalamanca/go-clean-boilerplate/pkg/config"
	log "github.com/adolsalamanca/go-clean-boilerplate/pkg/logger"
)

type Service struct {
	repo   repository.ItemRepository
	logger Logger
	// metrics
	// tracing
}

func NewService(config config.Provider, logger Logger, collector MetricsCollector) (*Service, error) {
	repo, err := persistence.NewPsqlRepository(config)
	if err != nil {
		return nil, err
	}
	return &Service{
		repo:   repo,
		logger: logger,
	}, nil
}

func (s Service) GetItems() ([]entities.Item, error) {
	i, err := s.repo.FindAllItems()
	if err != nil {
		s.logger.Error("could not get items",
			log.NewFieldString("error", err.Error()),
		)

		return nil, err
	}

	return i, nil
}

func (s Service) CreateItem(i entities.Item) error {
	err := s.repo.StoreItem(i)
	if err != nil {
		s.logger.Error("could not create items",
			log.NewFieldString("error", err.Error()),
		)
		return err
	}
	return nil
}
