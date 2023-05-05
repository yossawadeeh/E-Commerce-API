package delivery

import (
	"e-commerce-api/constant"
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/models"
	"e-commerce-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var err error

type OrderHandler struct {
	orderUsecase domains.OrderUsecase
}

func NewCustomerHandler(usecase domains.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		orderUsecase: usecase,
	}
}

func (t *OrderHandler) CreateOrder(c *gin.Context) {
	cutomerId := c.MustGet("id").(float64)

	var req *response.OrderRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	if req.Order.AddressId == 0 || req.Order.ShipperId == 0 {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(constant.InvalidField, http.StatusBadRequest))
		return
	}

	req.Order.CustomerId = uint(cutomerId)
	if err := t.orderUsecase.CreateOrder(req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: req,
	}))
}

func (t *OrderHandler) GetOrderCustomerById(c *gin.Context) {
	cutomerId := c.MustGet("id").(float64)
	orderIdStr := c.Param("orderId")
	orderId, _ := strconv.Atoi(orderIdStr)

	var order *models.Order
	if order, err = t.orderUsecase.GetOrderCustomerById(uint(orderId), uint(cutomerId)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: order,
	}))
}

func (t *OrderHandler) GetOrderCustomerByIdResponse(c *gin.Context) {
	cutomerId := c.MustGet("id").(float64)
	orderIdStr := c.Param("orderId")
	orderId, _ := strconv.Atoi(orderIdStr)

	var orderResponse *response.OrderResponse
	if orderResponse, err = t.orderUsecase.GetOrderCustomerByIdResponse(uint(orderId), uint(cutomerId)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: orderResponse,
	}))
}

func (t *OrderHandler) GetOrderById(c *gin.Context) {
	orderIdStr := c.Param("orderId")
	orderId, _ := strconv.Atoi(orderIdStr)

	var order *models.Order
	if order, err = t.orderUsecase.GetOrderById(uint(orderId)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: order,
	}))
}

// GetOrderByIdResponse godoc
// @Summary Get order by id response
// @Description Get order and order details by id
// @Security 	bearer
// @securityDefinitions.apikey bearer
// @in 			header
// @name Authorization
// @Tags    [Shop] Order
// @Produce json
// @Param        orderId   path      uint  true  "Order ID"
// @Success 200 {object} utils.SuccessMessagePrototype
// @Failure 400 {object} utils.ErrorMessagePrototype
// @Router /v1/shop/orders/{orderId} [get]
func (t *OrderHandler) GetOrderByIdResponse(c *gin.Context) {
	orderIdStr := c.Param("orderId")
	orderId, _ := strconv.Atoi(orderIdStr)

	var orderResponse *response.OrderResponse
	if orderResponse, err = t.orderUsecase.GetOrderByIdResponse(uint(orderId)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: orderResponse,
	}))
}

func (t *OrderHandler) GetAllCustomerOrders(c *gin.Context) {
	cutomerId := c.MustGet("id").(float64)

	var orders *[]models.Order
	if orders, err = t.orderUsecase.GetAllCustomerOrders(uint(cutomerId)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: orders,
	}))
}

func (t *OrderHandler) CreatePayment(c *gin.Context) {
	cutomerId := c.MustGet("id").(float64)

	var req *models.Payment
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	if req.OrderId == 0 || req.PaymentTypeId == 0 || req.PaymentAmount == 0 {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(constant.InvalidField, http.StatusBadRequest))
		return
	}

	if err := t.orderUsecase.CreatePayment(req, uint(cutomerId)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: req,
	}))
}

func (t *OrderHandler) UpdateOrder(c *gin.Context) {
	var req response.UpdateOrderRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	if req.OrderId == 0 || req.ShipperId == 0 || req.OrderStatusId == 0 || req.AddressId == 0 {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(constant.InvalidField, http.StatusBadRequest))
		return
	}

	var order *models.Order
	if order, err = t.orderUsecase.UpdateOrder(req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: order,
	}))
}
