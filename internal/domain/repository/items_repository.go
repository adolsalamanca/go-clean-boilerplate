package repository

import "github.com/adolsalamanca/go-clean-boilerplate/internal/domain/entities"

type ItemRepository interface {
	FindAllItems() ([]entities.Item, error)
	StoreItem(entities.Item) error
}
