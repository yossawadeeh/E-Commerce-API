package http

import (
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var err error

type ShopHandler struct {
	shopUsecase domains.ShopUsecase
}

func NewShopHandler(usecase domains.ShopUsecase) *ShopHandler {
	return &ShopHandler{
		shopUsecase: usecase,
	}
}

func (t *ShopHandler) GetEmployeeProfile(c *gin.Context) {
	empIdStr := c.Param("empId")
	empId, _ := strconv.Atoi(empIdStr)
	emp, err := t.shopUsecase.GetEmployeeProfile(uint(empId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.ErrorMessage(err.Error(), http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: &emp,
	}))
}

func (t *ShopHandler) GetAllRoles(c *gin.Context) {
	res, err := t.shopUsecase.GetAllRolesData()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.ErrorMessage(err.Error(), http.StatusNotFound))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Items: res,
	}))
}

func (t *ShopHandler) GetDailyReports(c *gin.Context) {
	shop_id := c.MustGet("shop_id").(float64)

	var req response.DailyReportsRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	req.Date, err = time.Parse("2006-01-02", req.DateText)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	var res *response.DailyReportsResponse
	req.ShopId = uint(shop_id)
	res, err := t.shopUsecase.GetDailyReports(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.ErrorMessage(err.Error(), http.StatusNotFound))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Items: res,
	}))
}
