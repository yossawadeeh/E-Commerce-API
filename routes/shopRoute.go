package routes

import (
	"e-commerce-api/database"
	shopHandler "e-commerce-api/feature/shop/delivery/http"
	shopRepository "e-commerce-api/feature/shop/repository"
	shopUsecase "e-commerce-api/feature/shop/usecase"
	"e-commerce-api/middleware"

	productHandler "e-commerce-api/feature/product/delivery/http"
	productRepository "e-commerce-api/feature/product/repository"
	productUsecase "e-commerce-api/feature/product/usecase"

	"github.com/gin-gonic/gin"
)

func ShopRoute(v1 *gin.RouterGroup) {

	newShopRepository := shopRepository.NewShopRepository(database.DB)
	newShopUsecase := shopUsecase.NewShopUsecase(newShopRepository)
	newShopHandler := shopHandler.NewShopHandler(newShopUsecase)

	newProductRepository := productRepository.NewProductRepository(database.DB)
	newProductUsecase := productUsecase.NewProductUsecase(newProductRepository)
	newProductHandler := productHandler.NewProductHandler(newProductUsecase)

	shopRoute := v1.Group("/shop", middleware.JWTAuthenEmployee())
	shopRoute.GET("/roles", newShopHandler.GetAllRoles)
	shopRoute.GET("/employee/:empId", newShopHandler.GetEmployeeProfile)
	shopRoute.POST("/product", newProductHandler.CreateProduct)
	shopRoute.PUT("/product", newProductHandler.UpdateProduct)
	shopRoute.DELETE("/product/:productId", newProductHandler.DeleteProduct)

	reportsRoute := v1.Group("/reports", middleware.JWTAuthenEmployee())
	reportsRoute.GET("/daily", newShopHandler.GetDailyReports)

	newShopRoute := v1.Group("/newShop")
	newShopRoute.POST("/", newShopHandler.CreateShop)
	newShopRoute.PUT("/", newShopHandler.UpdateShop)
	newShopRoute.DELETE("/:shopId", newShopHandler.DeleteShop)
}
