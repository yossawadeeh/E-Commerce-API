package routes

import (
	"e-commerce-api/database"
	orderHandler "e-commerce-api/feature/order/delivery/http"
	orderRepository "e-commerce-api/feature/order/repository"
	orderUsecase "e-commerce-api/feature/order/usecase"
	productRepository "e-commerce-api/feature/product/repository"

	"e-commerce-api/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoute(v1 *gin.RouterGroup) {
	productRepository := productRepository.NewProductRepository(database.DB)

	orderRepository := orderRepository.NewOrderRepository(database.DB)
	orderUsecase := orderUsecase.NewCustomerUsecase(orderRepository, productRepository)
	orderHandler := orderHandler.NewCustomerHandler(orderUsecase)

	orderRoute := v1.Group("/orders", middleware.JWTAuthenCustomer())
	orderRoute.POST("/", orderHandler.CreateOrder)
	//orderRoute.GET("/:orderId", orderHandler.GetOrderCustomerById)
	orderRoute.GET("/:orderId", orderHandler.GetOrderCustomerByIdResponse)
	orderRoute.GET("/", orderHandler.GetAllCustomerOrders)

	paymentRoute := v1.Group("/payments", middleware.JWTAuthenCustomer())
	paymentRoute.POST("/", orderHandler.CreatePayment)

	orderShopRoute := v1.Group("shop/orders", middleware.JWTAuthenEmployee())
	//orderShopRoute.GET("/:orderId", orderHandler.GetOrderById)
	orderShopRoute.GET("/:orderId", orderHandler.GetOrderByIdResponse)
	orderShopRoute.PUT("/", orderHandler.UpdateOrder)
}
