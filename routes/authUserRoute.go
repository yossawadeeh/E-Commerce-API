package routes

import (
	"e-commerce-api/database"
	authUserHandler "e-commerce-api/feature/authUser/delivery/http"
	authUserRepository "e-commerce-api/feature/authUser/repository"
	authUserUsecase "e-commerce-api/feature/authUser/usecase"

	customerRepository "e-commerce-api/feature/customer/repository"
	shopRepository "e-commerce-api/feature/shop/repository"

	"github.com/gin-gonic/gin"
)

func AuthUserRoute(v1 *gin.RouterGroup) {

	authUserRoute := v1.Group("/auth")

	newShopRepository := shopRepository.NewShopRepository(database.DB)
	newCustomerRepository := customerRepository.NewCustomerRepository(database.DB)

	newAuthUserRepository := authUserRepository.NewAuthUserRepository(database.DB)
	newAuthUserUsecase := authUserUsecase.NewAuthUserUsecase(newAuthUserRepository, newShopRepository, newCustomerRepository)
	newAuthUserHandler := authUserHandler.NewAuthUserHandler(newAuthUserUsecase)

	authUserRoute.POST("/employee/register", newAuthUserHandler.RegisterEmployee)
	authUserRoute.POST("/employee/login", newAuthUserHandler.LoginEmployee)
	authUserRoute.POST("/customer/register", newAuthUserHandler.RegisterCustomer)
	authUserRoute.POST("/customer/login", newAuthUserHandler.LoginCustomer)
}
