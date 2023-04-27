package usecase

import (
	"e-commerce-api/constant"
	"e-commerce-api/domains"
	"e-commerce-api/domains/response"
	"e-commerce-api/models"
	"e-commerce-api/utils"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var err error

type authUserUsecase struct {
	authUserRepo domains.AuthUserRepository
	shopRepo     domains.ShopRepository
	customerRepo domains.CustomerRepository
}

func NewAuthUserUsecase(authRepo domains.AuthUserRepository, shopRepo domains.ShopRepository, customerRepo domains.CustomerRepository) domains.AuthUserUsecase {
	return &authUserUsecase{
		authUserRepo: authRepo,
		shopRepo:     shopRepo,
		customerRepo: customerRepo,
	}
}

func (t *authUserUsecase) CreateEmployee(req *models.Employee) (res *response.EmployeeProfileResponse, err error) {

	if req.Username == "" || req.Email == "" || req.Password == "" || req.ShopOwnerId <= 0 {
		return nil, errors.New(constant.InvalidField)
	}

	emailIsExist, err := t.authUserRepo.CheckEmailEmployeeIsExist(req.Email)
	if emailIsExist == true {
		return nil, errors.New(constant.EmailDuplicate)
	}

	usernameIsExist, err := t.authUserRepo.CheckUsernameEmployeeIsExist(req.Username)
	if usernameIsExist == true {
		return nil, errors.New(constant.UsernameDuplicate)
	}

	_, err = t.shopRepo.GetShopById(req.ShopOwnerId)
	if err != nil {
		return nil, errors.New(constant.ShopIdNotFound)
	}

	// hash password
	encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	req.Password = string(encryptPassword)
	if res, err = t.authUserRepo.CreateEmployee(req); err != nil {
		return nil, err
	}
	return res, nil
}

func (t authUserUsecase) EmployeeLogin(loginData response.LoginEmployeeRequest) (empToken *string, err error) {

	if loginData.Email == "" || loginData.Password == "" {
		return nil, errors.New(constant.InvalidField)
	}
	var empData *models.Employee
	if empData, err = t.shopRepo.GetEmployeeByEmail(loginData.Email); err != nil {
		return nil, errors.New(constant.LoginFailed)
	}

	err = bcrypt.CompareHashAndPassword([]byte(empData.Password), []byte(loginData.Password))
	if err != nil {
		return nil, errors.New(constant.LoginFailed)
	}

	var hmacSampleSecret []byte
	hmacSampleSecret = []byte(utils.ViperGetString("jwt.secretKeyEmployee"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      empData.ID,
		"role_id": empData.RoleId,
		"shop_id": empData.ShopOwnerId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	return &tokenString, err
}

func (t *authUserUsecase) CreateCustomer(req models.Customer) (res *response.CustomerProfileResponse, err error) {

	if req.Username == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New(constant.InvalidField)
	}

	emailIsExist, err := t.authUserRepo.CheckEmailCustomerIsExist(req.Email)
	if emailIsExist == true {
		return nil, errors.New(constant.EmailDuplicate)
	}

	usernameIsExist, err := t.authUserRepo.CheckUsernameCustomerIsExist(req.Username)
	if usernameIsExist == true {
		return nil, errors.New(constant.UsernameDuplicate)
	}

	// hash password
	encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	req.Password = string(encryptPassword)
	if res, err = t.authUserRepo.CreateCustomer(req); err != nil {
		return nil, err
	}
	return res, nil
}

func (t *authUserUsecase) CustomerLogin(loginData response.LoginCustomerRequest) (cusToken *string, err error) {

	if loginData.Email == "" || loginData.Password == "" {
		return nil, errors.New(constant.InvalidField)
	}
	var cusData *models.Customer
	if cusData, err = t.customerRepo.GetCustomerByEmail(loginData.Email); err != nil {
		return nil, errors.New(constant.LoginFailed)
	}

	fmt.Println("ggg ", cusData)

	err = bcrypt.CompareHashAndPassword([]byte(cusData.Password), []byte(loginData.Password))
	if err != nil {
		return nil, errors.New(constant.LoginFailed)
	}

	fmt.Println("ggsssssg ", cusData.ID)

	var hmacSampleSecret []byte
	hmacSampleSecret = []byte(utils.ViperGetString("jwt.secretKeyCustomer"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  cusData.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	return &tokenString, err
}
