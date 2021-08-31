package domain

type ItemRepository interface {
	FindAllItems() ([]Item, error)
	StoreItem(Item) error
}
