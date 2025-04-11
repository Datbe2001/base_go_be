package controller

import (
	"base_go_be/internal/dto"
	"base_go_be/internal/service"
	"base_go_be/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

// GetProductByID @Summary Get product by ID
// @Description Get product details by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dto.ProductDetailDto
// @Router /product/detail/{id} [get]
func (pc *ProductController) GetProductByID(c *gin.Context) {
	idParam := c.Param("id")
	idUint64, err := strconv.ParseUint(idParam, 10, 0)
	id := uint(idUint64)
	if err != nil {
		response.DataDetailResponse(c, 422, response.ErrCodeInvalidParams, nil)
		return
	}

	product, err := pc.productService.GetProductByID(id)
	if err != nil {
		response.ErrorResponse(c, 404, err.Error())
		return
	}
	response.SuccessResponse(c, product)
}

// GetListProduct @Summary Get list of products
// @Description Get all products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} dto.ProductResponseDto
// @Router /product/list [get]
func (pc *ProductController) GetListProduct(c *gin.Context) {
	products, err := pc.productService.GetListProduct()
	if err != nil {
		response.ErrorResponse(c, 500, err.Error())
		return
	}
	response.SuccessResponse(c, products)
}

// CreateProduct @Summary Create a new product
// @Description Create a new product with name, description, and user ID
// @Tags Product
// @Accept json
// @Produce json
// @Param product body dto.ProductRequestDto true "Product Request"
// @Success 200 {object} map[string]uint
// @Router /product/create [post]
func (pc *ProductController) CreateProduct(c *gin.Context) {
	productRequest := dto.ProductRequestDto{}

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		response.ErrorResponse(c, 400, "Invalid request payload")
		return
	}

	productID, err := pc.productService.CreateProduct(productRequest.Name, productRequest.Description, productRequest.UserID)
	if err != nil {
		response.ErrorResponse(c, 405, err.Error())
		return
	}

	response.SuccessResponse(c, gin.H{"product_id": productID})
}
