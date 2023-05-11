package http

import (
	"e-commerce-api/constant"
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
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

// RegisterEmployee godoc
// @Summary Register Employee
// @Tags    Employee Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.SuccessMessagePrototype
// @Failure 404 {object} utils.ErrorMessagePrototype
// @Param Body body response.RegisterEmployeeRequest true "username, email, password, firstname, lastname, phone, shop_id, role_id"
// @Router /v1/auth/employee/register [post]
func (t *AuthUserHandler) RegisterEmployee(c *gin.Context) {
	employeeReq := response.RegisterEmployeeRequest{}
	if err := c.Bind(&employeeReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	if employeeReq.Username == "" || employeeReq.Email == "" || employeeReq.Password == "" || employeeReq.ShopOwnerId <= 0 {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(constant.InvalidField, http.StatusBadRequest))
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

// LoginEmployee godoc
// @Summary Employee login
// @Tags    Employee Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.SuccessMessagePrototype
// @Failure 404 {object} utils.ErrorMessagePrototype
// @Param Body body response.LoginEmployeeRequest  true "email, password"
// @Router /v1/auth/employee/login [post]
func (t *AuthUserHandler) LoginEmployee(c *gin.Context) {
	loginData := response.LoginEmployeeRequest{}
	if err := c.Bind(&loginData); err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorMessage(constant.LoginFailed, http.StatusUnauthorized))
		return
	}

	if loginData.Email == "" || loginData.Password == "" {
		c.JSON(http.StatusUnauthorized, utils.ErrorMessage(constant.InvalidField, http.StatusUnauthorized))
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

// RegisterCustomer godoc
// @Summary Customer register
// @Tags    Customer Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.SuccessMessagePrototype
// @Failure 404 {object} utils.ErrorMessagePrototype
// @Param Body body response.RegisterCustomerRequest true "username, email, password, firstname, lastname, phone, birthday_text"
// @Router /v1/auth/customer/register [post]
func (t *AuthUserHandler) RegisterCustomer(c *gin.Context) {
	customerReq := response.RegisterCustomerRequest{}
	if err := c.Bind(&customerReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	if customerReq.Username == "" || customerReq.Email == "" || customerReq.Password == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(constant.InvalidField, http.StatusBadRequest))
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

// LoginCustomer godoc
// @Summary Customer login
// @Tags    Customer Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.SuccessMessagePrototype
// @Failure 404 {object} utils.ErrorMessagePrototype
// @Param Body body response.LoginCustomerRequest  true "email, password"
// @Router /v1/auth/customer/login [post]
func (t *AuthUserHandler) LoginCustomer(c *gin.Context) {
	loginData := response.LoginCustomerRequest{}
	if err := c.Bind(&loginData); err != nil {
		c.JSON(http.StatusUnauthorized, utils.ErrorMessage(constant.LoginFailed, http.StatusUnauthorized))
		return
	}

	if loginData.Email == "" || loginData.Password == "" {
		c.JSON(http.StatusUnauthorized, utils.ErrorMessage(constant.InvalidField, http.StatusUnauthorized))
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
