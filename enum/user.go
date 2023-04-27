package enum

type UserType struct {
	EMPLOYEE uint
	CUSTOMER uint
}

var UserEnum = UserType{
	EMPLOYEE: 1,
	CUSTOMER: 2,
}
