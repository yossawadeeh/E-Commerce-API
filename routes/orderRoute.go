package routes

// import (
// 	"e-commerce-api/database"
// 	orderHandler "e-commerce-api/feature/order/delivery/http"
// 	orderRepository "e-commerce-api/feature/order/repository"
// 	orderUsecase "e-commerce-api/feature/order/usecase"
// 	"e-commerce-api/middleware"

// 	"github.com/gin-gonic/gin"
// )

// func OrderRoute(v1 *gin.Context) {
// 	orderRepository := orderRepository.NewOrderRepository(database.DB)
// 	orderUsecase := orderUsecase.NewCustomerUsecase(orderRepository)
// 	orderHandler := orderHandler.NewCustomerHandler(orderUsecase)

// 	orderRoute := v1.Group("/orders", middleware.JWTAuthenCustomer())
// 	//orderRoute.POST("/", orderHandler.UpdateCarts)

// 	orderShopRoute := v1.Group("shop/orders", middleware.JWTAuthenEmployee())
// 	//orderShopRoute.POST("/", orderHandler.UpdateCarts)

// }
