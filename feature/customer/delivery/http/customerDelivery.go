package http

import (
	"e-commerce-api/constant"
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/models"
	"e-commerce-api/utils"
	"fmt"
	"io/ioutil"
	"net/http"

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
	// productIdStr := c.Param("productId")
	// productId, _ := strconv.Atoi(productIdStr)
	var productIds response.ProductIdRequest

	if err := c.Bind(&productIds); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	if err = t.customerUsecase.DeleteProductFromCarts(uint(customerId), productIds); err != nil {
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

	if req.AddressDetail == "" || req.Districts == "" || req.Province == "" || req.Country == "" || req.PostalCode == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(constant.InvalidField, http.StatusBadRequest))
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

func (t *CustomerHandler) UploadImagesProfile(c *gin.Context) {
	customerId := c.MustGet("id").(float64)

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	c.Request.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := c.Request.FormFile("myFile")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	if err := t.customerUsecase.UploadImagesProfile(uint(customerId), fileBytes); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: fileBytes,
	}))
}

func (t *CustomerHandler) GetImageProfileBytes(c *gin.Context) {
	customerId := c.MustGet("id").(float64)

	var imgBytes []byte
	if imgBytes, err = t.customerUsecase.GetImageProfileBytes(uint(customerId)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorMessage(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Item: imgBytes,
	}))
}
