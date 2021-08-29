package repository

import "github.com/adolsalamanca/go-rest-boilerplate/internal/domain/entities"

type ItemRepository interface {
	GetItems() ([]*entities.Item, error)
	CreateItem(entities.Item) error
}
