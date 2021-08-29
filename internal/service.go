package internal

type ItemRepository interface {
	CreateItem(Item) error
}

type Service struct {
	repo ItemRepository
}

func (s Service) CreateItem(i Item) error {
	return s.repo.CreateItem(i)
}

func NewService() *Service {
	return &Service{}
}
