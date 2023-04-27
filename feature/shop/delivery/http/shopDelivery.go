package http

import (
	"e-commerce-api/domains"
	"e-commerce-api/utils"
	"net/http"
	"strconv"

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
