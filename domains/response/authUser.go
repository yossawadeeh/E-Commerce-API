package response

type LoginEmployeeRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginCustomerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
