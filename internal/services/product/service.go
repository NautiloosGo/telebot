package product

import "log"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	if idx < 0 || idx >= len(allProducts) {
		log.Println("there are no element in Products with number", idx)
		return &Product{Title: "there are no element in Products with number"}, nil
	}
	return &allProducts[idx], nil
}
