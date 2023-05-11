package response

import "time"

type LoginEmployeeRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginCustomerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterEmployeeRequest struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	RoleId      uint   `json:"role_id"`
	ShopOwnerId uint   `json:"shop_id"`
}

type RegisterCustomerRequest struct {
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	Firstname    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	Password     string    `json:"password"`
	Phone        string    `json:"phone"`
	Age          *int64    `json:"age"`
	Birthday     time.Time `json:"-"`
	BirthdayText string    `json:"birthday_text"`
}
