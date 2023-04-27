package http

import (
	"e-commerce-api/constant"
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/models"
	"e-commerce-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var err error

type AuthUserHandler struct {
	authUserUsecase domains.AuthUserUsecase
}

func NewAuthUserHandler(usecase domains.AuthUserUsecase) *AuthUserHandler {
	return &AuthUserHandler{
		authUserUsecase: usecase,
	}
}

func (t *AuthUserHandler) RegisterEmployee(c *gin.Context) {
	employeeReq := models.Employee{}
	if err := c.Bind(&employeeReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	var res *response.EmployeeProfileResponse
	if res, err = t.authUserUsecase.CreateEmployee(&employeeReq); err != nil {
		switch err.Error() {
		case constant.InvalidField:
			c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
			return
		case constant.EmailDuplicate:
			c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
			return
		case constant.UsernameDuplicate:
			c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
			return
		case constant.ShopIdNotFound:
			c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
			return
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorMessage(err.Error(), http.StatusInternalServerError))
			return
		}

	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: res,
	}))
}

func (t *AuthUserHandler) LoginEmployee(c *gin.Context) {
	loginData := response.LoginEmployeeRequest{}
	if err := c.Bind(&loginData); err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorMessage(constant.LoginFailed, http.StatusUnauthorized))
		return
	}

	token, err := t.authUserUsecase.EmployeeLogin(loginData)
	if err != nil {
		switch err.Error() {
		case constant.LoginFailed:
			c.JSON(http.StatusUnauthorized, utils.ErrorMessage(err.Error(), http.StatusUnauthorized))
			return
		case constant.InvalidField:
			c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
			return
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorMessage(err.Error(), http.StatusInternalServerError))
			return
		}
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: token,
	}))
}

func (t *AuthUserHandler) RegisterCustomer(c *gin.Context) {
	customerReq := models.Customer{}
	if err := c.Bind(&customerReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}
	customerReq.Birthday, _ = time.Parse("2006-01-02", customerReq.BirthdayText)
	var res *response.CustomerProfileResponse
	if res, err = t.authUserUsecase.CreateCustomer(customerReq); err != nil {
		switch err.Error() {
		case constant.InvalidField:
			c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
			return
		case constant.EmailDuplicate:
			c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
			return
		case constant.UsernameDuplicate:
			c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
			return
		case constant.ShopIdNotFound:
			c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
			return
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorMessage(err.Error(), http.StatusInternalServerError))
			return
		}

	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: res,
	}))
}

func (t *AuthUserHandler) LoginCustomer(c *gin.Context) {
	loginData := response.LoginCustomerRequest{}
	if err := c.Bind(&loginData); err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorMessage(constant.LoginFailed, http.StatusUnauthorized))
		return
	}

	token, err := t.authUserUsecase.CustomerLogin(loginData)
	if err != nil {
		switch err.Error() {
		case constant.LoginFailed:
			c.JSON(http.StatusUnauthorized, utils.ErrorMessage(err.Error(), http.StatusUnauthorized))
			return
		case constant.InvalidField:
			c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
			return
		default:
			c.JSON(http.StatusInternalServerError, utils.ErrorMessage(err.Error(), http.StatusInternalServerError))
			return
		}
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: token,
	}))
}
