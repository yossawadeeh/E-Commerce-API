package delivery

import (
	"e-commerce-api/domains"
	"e-commerce-api/models"
	"e-commerce-api/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase domains.ProductUsecase
}

var err error

func NewProductHandler(usecase domains.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: usecase,
	}
}

func (t *ProductHandler) GetAllProducts(c *gin.Context) {
	shopId := c.MustGet("shop_id").(float64)

	var products *[]models.Product
	products, err = t.productUsecase.GetAllProducts(uint(shopId))

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Items: products,
	}))
}

func (t *ProductHandler) GetProductById(c *gin.Context) {
	productIdParam := c.Param("productId")
	productId, _ := strconv.Atoi(productIdParam)

	var product *models.Product
	product, err = t.productUsecase.GetProductById(uint(productId))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorMessage(err.Error(), http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: product,
	}))
}

func (t *ProductHandler) CreateProduct(c *gin.Context) {
	shopId := c.MustGet("shop_id").(float64)
	var product models.Product
	if err := c.Bind(&product); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	product.ShopOwnerId = uint(shopId)
	err = t.productUsecase.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: product,
	}))
}

func (t *ProductHandler) UpdateProduct(c *gin.Context) {
	shopId := c.MustGet("shop_id").(float64)
	var product models.Product
	if err := c.Bind(&product); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	product.ShopOwnerId = uint(shopId)
	err = t.productUsecase.UpdateProduct(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: product,
	}))
}

func (t *ProductHandler) DeleteProduct(c *gin.Context) {
	shopId := c.MustGet("shop_id").(float64)
	productIdParam := c.Param("productId")
	productId, _ := strconv.Atoi(productIdParam)

	fmt.Println("productIdParam: ", productIdParam)
	fmt.Println("productId: ", productId)

	var id uint
	id, err := t.productUsecase.DeleteProduct(uint(shopId), uint(productId))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Id: id,
	}))
}
