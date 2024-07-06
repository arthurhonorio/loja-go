package usecase

import "go-api/model"

type productUseCase struct {
	//Repository
}

func NewProductUseCase() productUseCase {
	return productUseCase{}
}

func (pu *productUsecase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}
