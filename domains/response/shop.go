package response

type EmployeeProfileResponse struct {
	ID          uint   `json:"employee_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Phone       string `json:"phone"`
	ShopOwnerId uint   `json:"user_id"`
	RoleId      uint   `json:"role_id"`
}
