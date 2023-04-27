package routes

import (
	"e-commerce-api/database"
	customerHandler "e-commerce-api/feature/customer/delivery/http"
	customerRepository "e-commerce-api/feature/customer/repository"
	customerUsecase "e-commerce-api/feature/customer/usecase"
	productRepository "e-commerce-api/feature/product/repository"

	"e-commerce-api/middleware"

	"github.com/gin-gonic/gin"
)

func CustomerRoute(v1 *gin.RouterGroup) {

	productUserRepository := productRepository.NewProductRepository(database.DB)

	customerUserRepository := customerRepository.NewCustomerRepository(database.DB)
	customerUserUsecase := customerUsecase.NewCustomerUsecase(customerUserRepository, productUserRepository)
	customerUserHandler := customerHandler.NewCustomerHandler(customerUserUsecase)

	// Carts route
	authUserRoute := v1.Group("/carts", middleware.JWTAuthenCustomer())
	authUserRoute.POST("/", customerUserHandler.UpdateCarts)
	authUserRoute.DELETE("/:productId", customerUserHandler.DeleteProductFromCarts)

	// Customer route
	customerRoute := v1.Group("/customer", middleware.JWTAuthenCustomer())
	customerRoute.POST("/address", customerUserHandler.CreateAddress)
}
