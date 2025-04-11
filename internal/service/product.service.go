package service

import (
	"base_go_be/internal/dto"
	"base_go_be/internal/model"
	"base_go_be/internal/repo"
)

type IProductService interface {
	GetProductByID(id uint) (*dto.ProductDetailDto, error)
	GetListProduct() ([]dto.ProductResponseDto, error)
	CreateProduct(name, description string, userID uint) (uint, error)
}

type ProductService struct {
	productRepo repo.IProductRepository
}

func NewProductService(productRepo repo.IProductRepository) IProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (ps *ProductService) GetProductByID(id uint) (*dto.ProductDetailDto, error) {
	product, err := ps.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.ProductDetailDto{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}

func (ps *ProductService) GetListProduct() ([]dto.ProductResponseDto, error) {
	products, err := ps.productRepo.FindAll()
	if err != nil {
		return nil, err
	}
	var productDto []dto.ProductResponseDto
	for _, product := range products {
		productDto = append(productDto, dto.ProductResponseDto{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			UserID:      product.UserID,
			User: dto.UserResponseDto{
				Id:       product.User.ID,
				Email:    product.User.Email,
				Username: product.User.Username,
				Role:     product.User.Role,
			},
		})
	}

	return productDto, nil
}

func (ps *ProductService) CreateProduct(name, description string, userID uint) (uint, error) {

	product := &model.Product{
		Name:        name,
		Description: description,
		UserID:      userID,
	}

	createdProduct, err := ps.productRepo.Create(product)
	if err != nil {
		return 0, err
	}

	return createdProduct.ID, nil
}
