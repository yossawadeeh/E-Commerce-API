package http

import (
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/models"
	"e-commerce-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var err error

type CustomerHandler struct {
	customerUsecase domains.CustomerUsecase
}

func NewCustomerHandler(usecase domains.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{
		customerUsecase: usecase,
	}
}

func (t *CustomerHandler) UpdateCarts(c *gin.Context) {
	var req models.Cart
	customerId := c.MustGet("id").(float64)

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	req.CustomerId = uint(customerId)
	var res *response.AddToCartResponse
	if res, err = t.customerUsecase.UpdateCarts(req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: res,
	}))
}

func (t *CustomerHandler) DeleteProductFromCarts(c *gin.Context) {
	customerId := c.MustGet("id").(float64)
	productIdStr := c.Param("productId")
	productId, _ := strconv.Atoi(productIdStr)

	if err = t.customerUsecase.DeleteProductFromCarts(uint(customerId), uint(productId)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	var status int = 200
	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		StatusCode: &status,
	}))
}

func (t *CustomerHandler) CreateAddress(c *gin.Context) {
	customerId := c.MustGet("id").(float64)
	var req models.Address
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	req.CustomerId = uint(customerId)
	if err := t.customerUsecase.CreateAddress(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: &req,
	}))
}
