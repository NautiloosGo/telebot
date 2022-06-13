package product

import "log"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Sku {
	return catalog.Products
}

func (s *Service) Get(idx int) (*Sku, bool) {
	if idx < 0 || idx >= len(catalog.Products) {
		log.Println("there are no element in Products with number ", idx)
		return nil, false
	}
	return &catalog.Products[idx], true
}
