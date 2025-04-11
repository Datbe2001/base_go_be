//go:build wireinject

package wire

import (
	"base_go_be/internal/controller"
	"base_go_be/internal/repo"
	"base_go_be/internal/service"
	"github.com/google/wire"
)

func InitProductRouterHandler() (*controller.ProductController, error) {
	wire.Build(
		repo.NewProductRepository,
		service.NewProductService,
		controller.NewProductController,
	)
	return new(controller.ProductController), nil
}
