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

// GetProductByID godoc
// @Summary Get product by ID
// @Description Get product details by ID
// @Tags product
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} response.Response{data=dto.ProductDetailDto}
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 404 {object} response.Response "Product not found"
// @Failure 422 {object} response.Response "Invalid product ID"
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

// GetListProduct godoc
// @Summary Get list of products
// @Description Get all products
// @Tags product
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Failure 401 {object} response.Response "Unauthorized"
// @Success 200 {object} response.Response{data=[]dto.ProductResponseDto}
// @Failure 500 {object} response.Response "Internal server error"
// @Router /product/list [get]
func (pc *ProductController) GetListProduct(c *gin.Context) {
	products, err := pc.productService.GetListProduct()
	if err != nil {
		response.ErrorResponse(c, 500, err.Error())
		return
	}
	response.SuccessResponse(c, products)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with name, description, and user ID
// @Tags product
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product body dto.ProductRequestDto true "Product Request"
// @Success 200 {object} response.Response{data=map[string]uint}
// @Failure 400 {object} response.Response "Invalid request payload"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 405 {object} response.Response "Method not allowed"
// @Router /product/create [post]
func (pc *ProductController) CreateProduct(c *gin.Context) {
	productRequest := dto.ProductRequestDto{}

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		response.ErrorResponse(c, 400, "Invalid request payload")
		return
	}

	// Get user ID from JWT token context instead of request
	userID, exists := c.Get("userID")
	if !exists {
		response.ErrorResponse(c, 401, "Unauthorized")
		return
	}

	productID, err := pc.productService.CreateProduct(productRequest.Name, productRequest.Description, userID.(uint))
	if err != nil {
		response.ErrorResponse(c, 405, err.Error())
		return
	}

	response.SuccessResponse(c, gin.H{"product_id": productID})
}
