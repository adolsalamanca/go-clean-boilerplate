package _interface

import (
	"github.com/adolsalamanca/go-clean-boilerplate/internal/domain"
	"github.com/adolsalamanca/go-clean-boilerplate/internal/infrastructure"
	"github.com/adolsalamanca/go-clean-boilerplate/pkg/config"
	log "github.com/adolsalamanca/go-clean-boilerplate/pkg/logger"
)

type Service struct {
	repo   domain.ItemRepository
	logger Logger
	// metrics
	// tracing
}

func NewService(config config.Provider, logger Logger, collector MetricsCollector) (*Service, error) {
	repo, err := infrastructure.NewPsqlRepository(config)
	if err != nil {
		return nil, err
	}
	return &Service{
		repo:   repo,
		logger: logger,
	}, nil
}

func (s Service) GetItems() ([]domain.Item, error) {
	i, err := s.repo.FindAllItems()
	if err != nil {
		s.logger.Error("could not get items",
			log.NewFieldString("error", err.Error()),
		)

		return nil, err
	}

	return i, nil
}

func (s Service) CreateItem(i domain.Item) error {
	err := s.repo.StoreItem(i)
	if err != nil {
		s.logger.Error("could not create items",
			log.NewFieldString("error", err.Error()),
		)
		return err
	}
	return nil
}
